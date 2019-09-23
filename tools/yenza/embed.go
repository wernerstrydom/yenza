package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func embed(path, packageName, variableName, outputPath string) error {
	files, err := readFiles(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}

	data := &embedTemplateArgs{
		PackageName:  packageName,
		VariableName: variableName,
		Files:        files,
	}

	tmpl, err := template.
		New("hello").
		Funcs(template.FuncMap{
			"bytes": bytes,
		}).
		Parse(text)
	if err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}

	return nil
}

var text = `
package {{ .PackageName }}

var {{ .VariableName }} = map[string][]byte{
{{- range $key, $value := .Files}}
  "{{ $key }}": []byte {
{{ bytes $value }}
  },
{{ end -}}
}
`

type embedTemplateArgs struct {
	PackageName  string
	VariableName string
	Files        map[string][]byte
}

func bytes(buffer []byte) string {
	var builder strings.Builder
	n := 0
	for index := range buffer {
		if n%10 == 0 {
			builder.WriteString("    ")
		}
		n++
		builder.WriteString(fmt.Sprintf("0x%02x, ", buffer[index]))
		if n%10 == 0 && len(buffer) != n {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}

func readFiles(path string) (map[string][]byte, error) {
	r := make(map[string][]byte)
	err := filepath.Walk(path,
		func(p string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fi.IsDir() {
				return nil
			}

			s := strings.TrimPrefix(p, path)
			s = strings.Trim(s, "/")
			s = strings.Replace(s, "\\", "/", -1)

			bytes, err := ioutil.ReadFile(p)
			r[s] = bytes
			return nil
		})
	return r, err
}
