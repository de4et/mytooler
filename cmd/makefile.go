package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const makefileJavaString = `
build:
	javac *.java -d ./bin

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
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		template := args[0]
		createMakefile(path, template)
	},
}

func createMakefile(path string, template string) {
	file, err := os.Create(path + "\\Makefile")
	if err != nil {
		panic(err)
	}

	var str string
	switch template {
	case "go":
		str = makefileGoString
	case "java":
		str = makefileJavaString
	default:
		panic(fmt.Errorf("there is no such makefile template: %s", template))

	}

	file.WriteString(str)
	file.Close()
}

func init() {
	rootCmd.AddCommand(makefileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makefileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makefileCmd.Flags().StringToStringVar("toggle", "t", false, "Help message for toggle")
}
