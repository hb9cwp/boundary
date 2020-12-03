package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

// generate looks for migration sql in a directory for the given dialect and
// applies the templates below to the contents of the files, building up a
// migrations map for the dialect
func generate(dialect string) {
	baseDir := os.Getenv("GEN_BASEPATH") + "/internal/db/migrations"
	dir, err := os.Open(fmt.Sprintf("%s/%s", baseDir, dialect))
	if err != nil {
		fmt.Printf("error opening dir with dialect %s: %v\n", dialect, err)
		os.Exit(1)
	}
	versions, err := dir.Readdirnames(0)
	if err != nil {
		fmt.Printf("error reading dir names with dialect %s: %v\n", dialect, err)
		os.Exit(1)
	}
	outBuf := bytes.NewBuffer(nil)
	valuesBuf := bytes.NewBuffer(nil)

	sort.Strings(versions)

	isDev := false
	largestVer := 0
	for _, ver := range versions {
		verVal, err := strconv.Atoi(ver)
		if err != nil {
			if ver != "dev" {
				fmt.Printf("error reading major schema version directory %q.  Must be a number or 'dev'\n", ver)
				continue
			}
			verVal = largestVer + 1
		}
		if verVal > largestVer {
			largestVer = verVal
		}

		dir, err := os.Open(fmt.Sprintf("%s/%s/%s", baseDir, dialect, ver))
		if err != nil {
			fmt.Printf("error opening dir with dialect %s: %v\n", dialect, err)
			os.Exit(1)
		}
		names, err := dir.Readdirnames(0)
		if err != nil {
			fmt.Printf("error reading dir names with dialect %s: %v\n", dialect, err)
			os.Exit(1)
		}

		if ver == "ver" && len(names) > 0 {
			isDev = true
		}

		sort.Strings(names)
		for _, name := range names {
			if !strings.HasSuffix(name, ".sql") {
				continue
			}

			contents, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s/%s", baseDir, dialect, ver, name))
			if err != nil {
				fmt.Printf("error opening file %s with dialect %s: %v", name, dialect, err)
				os.Exit(1)
			}

			vName := name
			nameParts := strings.SplitN(name, "_", 2)
			if len(nameParts) != 2 {
				continue
			}

			nameVer, err := strconv.Atoi(nameParts[0])
			if err != nil {
				fmt.Printf("Unable to get file version from %q\n", name)
				continue
			}
			vName = fmt.Sprintf("%02d_%s", (verVal*1000)+nameVer, nameParts[1])

			if err := migrationsValueTemplate.Execute(valuesBuf, struct {
				Name     string
				Contents string
			}{
				Name:     vName,
				Contents: string(contents),
			}); err != nil {
				fmt.Printf("error executing migrations value template for file %s/%s: %s", ver, name, err)
				os.Exit(1)
			}
		}
	}
	if err := migrationsTemplate.Execute(outBuf, struct {
		Type         string
		Values       string
		DevMigration bool
	}{
		Type:         dialect,
		Values:       valuesBuf.String(),
		DevMigration: isDev,
	}); err != nil {
		fmt.Printf("error executing migrations value template for dialect %s: %s", dialect, err)
		os.Exit(1)
	}

	outFile := fmt.Sprintf("%s/%s.gen.go", baseDir, dialect)
	if err := ioutil.WriteFile(outFile, outBuf.Bytes(), 0644); err != nil {
		fmt.Printf("error writing file %q: %v\n", outFile, err)
		os.Exit(1)
	}
}

var migrationsTemplate = template.Must(template.New("").Parse(
	`// Code generated by "make migrations"; DO NOT EDIT.
package migrations
	
import (
	"bytes"
)

var DevMigration = {{ .DevMigration }}

var {{ .Type }}Migrations = map[string]*fakeFile{
	"migrations": {
		name: "migrations",
	},
	{{ .Values }}
}
`))

var migrationsValueTemplate = template.Must(template.New("").Parse(
	`"migrations/{{ .Name }}": {
		name: "{{ .Name }}",
		bytes: []byte(` + "`\n{{ .Contents }}\n`" + `),
	},
`))
