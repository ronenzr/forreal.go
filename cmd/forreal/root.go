package forreal

import (
	"fmt"
	"github.com/ronenzr/forreal.go/api/httpclient"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var rootCmd = &cobra.Command{
	Use:   "ForReal",
	Short: "ForReal - CLI For Real - CLI NLP tool for generating commands",
	Long: `ForReal is a command-line interface (CLI) tool that allows you to generate commands based on natural language. 
			With ForReal, 
			you can quickly and easily perform complex tasks without having to memorize complex command-line syntax.`,
	Run: func(cmd *cobra.Command, args []string) {
		execute, _ := cmd.Flags().GetBool("execute")
		phrase, _ := cmd.Flags().GetString("phrase")
		if phrase == "" {
			println("ForReal is missing a phrase argument, please add a phrase with the -p attribute")
		} else {
			command := httpclient.GetCommand(phrase)
			if command != "" {
				if execute {
					fmt.Println("executing: ", command)
					//args := strings.Fields(command)
					//fmt.Println(args)
					cmd := exec.Command("sh", "-c", command)
					out, err := cmd.Output()
					fmt.Print(string(out))
					if err != nil {
						log.Fatal(err)
					}
				} else {
					fmt.Println(command)
				}
			}

		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("execute", "x", false, "author name for copyright attribution")
	rootCmd.PersistentFlags().StringP("phrase", "p", "", "Describe in natural language what command do you want")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
