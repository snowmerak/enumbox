package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

const ENUMBOX = "enumbox.yml"

func main() {
	enumbox := kingpin.New("enumbox", "A tool to generate Go code from YAML files.")

	generateCmd := enumbox.Command("generate", "Generate Go code from YAML files.")
	generateNameArg := generateCmd.Arg("name", "The name of the package.").Required().String()

	initCmd := enumbox.Command("init", "Initialize a new YAML file.")
	initNameArg := initCmd.Arg("name", "The name of the package.").Required().String()

	cmd, err := enumbox.Parse(os.Args[1:])
	if err != nil {
		enumbox.FatalUsage(err.Error())
	}

	switch cmd {
	case generateCmd.FullCommand():
		base := filepath.Join(strings.Split(*generateNameArg, "/")...)
		path := filepath.Join(base, ENUMBOX)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			panic("File does not exist.")
		}
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		decoder := yaml.NewDecoder(f)
		box := EnumBox{}
		if err := decoder.Decode(&box); err != nil {
			panic(err)
		}
		f.Close()
		f, err = os.Create(filepath.Join(base, "enumbox.go"))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fmt.Fprintln(f, "package "+filepath.Base(base)+"\n")
		// Variables
		fmt.Fprint(f, "var (\n")
		for _, variable := range box.Variables {
			fmt.Fprintf(f, "\t%s %s = %s\n", variable.Name, variable.Type, variable.Value)
		}
		fmt.Fprint(f, ")\n\n")
		// Functions
		for _, variable := range box.Variables {
			fmt.Fprintf(f, "func %s() %s {\n", strings.ToUpper(variable.Name), variable.Type)
			fmt.Fprintf(f, "\treturn %s\n", variable.Name)
			fmt.Fprint(f, "}\n\n")
		}
		// GetName
		fmt.Fprint(f, "func NameOf(data interface{}) string {\n")
		fmt.Fprint(f, "\tswitch data {\n")
		for _, variable := range box.Variables {
			fmt.Fprintf(f, "\tcase %s:\n", variable.Name)
			fmt.Fprintf(f, "\t\treturn \"%s\"\n", variable.Name)
		}
		fmt.Fprint(f, "\t}\n")
		fmt.Fprint(f, "\treturn \"\"\n")
		fmt.Fprint(f, "}\n\n")
		// GetIndex
		fmt.Fprint(f, "func IndexOf(data interface{}) int {\n")
		fmt.Fprint(f, "\tswitch data {\n")
		for i, variable := range box.Variables {
			fmt.Fprintf(f, "\tcase %s:\n", variable.Name)
			fmt.Fprintf(f, "\t\treturn %d\n", i)
			i++
		}
		fmt.Fprint(f, "\t}\n")
		fmt.Fprint(f, "\treturn -1\n")
		fmt.Fprint(f, "}\n\n")
		// At
		fmt.Fprint(f, "func At(index int) interface{} {\n")
		fmt.Fprint(f, "\tswitch index {\n")
		for i, variable := range box.Variables {
			fmt.Fprintf(f, "\tcase %d:\n", i)
			fmt.Fprintf(f, "\t\treturn %s\n", variable.Name)
		}
		fmt.Fprint(f, "\t}\n")
		fmt.Fprint(f, "\treturn nil\n")
		fmt.Fprint(f, "}\n\n")
		// Euqal
		for _, variable := range box.Variables {
			fmt.Fprintf(f, "func EqualTo%s(b *%s) bool {\n", strings.ToUpper(variable.Name), variable.Type)
			fmt.Fprintf(f, "\treturn %s == *b\n", variable.Name)
			fmt.Fprint(f, "}\n\n")
		}
		// Names
		fmt.Fprint(f, "func Names() []string {\n")
		fmt.Fprint(f, "\treturn []string{\n")
		for _, variable := range box.Variables {
			fmt.Fprintf(f, "\t\t\"%s\",\n", variable.Name)
		}
		fmt.Fprint(f, "\t}\n")
		fmt.Fprint(f, "}\n\n")
	case initCmd.FullCommand():
		initData := EnumBox{
			Version: "0.0.1",
			Variables: []Variable{
				{
					Name:  "name",
					Type:  "string",
					Value: `"merak"`,
				},
				{
					Name:  "age",
					Type:  "int",
					Value: "28",
				},
			},
		}
		path := filepath.Join(strings.Split(*initNameArg, "/")...)
		if err := os.MkdirAll(path, 0755); err != nil {
			panic(err)
		}
		path = filepath.Join(path, ENUMBOX)
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		encoder := yaml.NewEncoder(f)
		if err := encoder.Encode(initData); err != nil {
			panic(err)
		}
	}
}
