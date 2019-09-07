/*
Copyright © 2019 Werner Strydom <hello@wernerstrydom.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var dataType string
var typeName string
var packageName string
var elementTypeName string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a strongly typed data type",
	Long:  `Generates uses a template to generate a strongly typed object`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Data Type", dataType)
		text := templates[dataType]
		tmpl, err := template.New("test").Parse(text)
		if err != nil {
			panic(err)
		}

		type Data struct {
			Package  string
			TypeName string
			T        string
			TRef     string
		}

		if len(typeName) == 0 {
			typeName = elementTypeName + dataType
		}

		data := Data{
			T:        elementTypeName,
			TRef:     "*" + elementTypeName,
			TypeName: typeName,
			Package:  packageName,
		}

		err = tmpl.Execute(os.Stdout, data)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().StringVarP(&dataType, "template", "t", "", "The template")
	generateCmd.MarkFlagRequired("template")
	generateCmd.PersistentFlags().StringVarP(&packageName, "package", "p", "", "The package.")
	generateCmd.MarkFlagRequired("package")
	generateCmd.PersistentFlags().StringVarP(&typeName, "name", "n", "", "The name of the class.")
	generateCmd.PersistentFlags().StringVarP(&elementTypeName, "element", "e", "", "The element.")
	generateCmd.MarkFlagRequired("elementTypeName")
}
