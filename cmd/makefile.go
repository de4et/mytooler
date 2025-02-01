package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const makefileJavaString = `
build:
	javac src/*.java -d ./bin

run: build
	java -cp ./bin $(program)	
`

const makefileGoString = `
build:
	@go build -o ./bin/app.exe main.go

run: build
	@./bin/app.exe
`

// makefileCmd represents the makefile command
var makefileCmd = &cobra.Command{
	Use:   "makefile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := os.Getwd()
		if err != nil {
			cmd.PrintErr(err)
		}

		templates := args
		addFlag, _ := cmd.Flags().GetBool("add")
		fmt.Println(addFlag)

		if addFlag {
			err = appendToMakefile(path, templates)
		} else {
			err = createMakefile(path, templates)
		}
		if err != nil {
			cmd.PrintErr(err)
		}

	},
}

func appendToMakefile(path string, templates []string) error {
	file, err := os.OpenFile(path+"\\Makefile", os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	finalStr := ""
	for _, v := range templates {
		str, err := getTemplateStringByName(v)
		if err != nil {
			return err
		}
		finalStr += "\n" + strings.TrimPrefix(str, "\n")
	}

	file.WriteString(finalStr)
	file.Close()
	return nil
}

func createMakefile(path string, templates []string) error {
	file, err := os.Create(path + "\\Makefile")
	if err != nil {
		return err
	}

	finalStr := ""
	for _, v := range templates {
		str, err := getTemplateStringByName(v)
		if err != nil {
			return err
		}
		finalStr += strings.TrimPrefix(str, "\n") + "\n"
	}

	file.WriteString(finalStr)
	file.Close()
	return nil
}

func getTemplateStringByName(name string) (string, error) {
	var str string
	switch name {
	case "go":
		str = makefileGoString
	case "java":
		str = makefileJavaString
	default:
		return str, fmt.Errorf("there is no such makefile template: %s", name)
	}
	return str, nil
}

func init() {
	rootCmd.AddCommand(makefileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makefileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	makefileCmd.Flags().BoolP("add", "a", false, "append to existed file or not")
}
