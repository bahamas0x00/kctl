/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var apiEndpoint, workspace, read, output string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kctl",
	Short: "CLI for Kong Gateway Admin api management",
	Long:  `CLI for Kong Gateway Admin api management`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.PersistentFlags().StringVarP(&apiEndpoint, "api", "i", "http://localhost:8001", "Use -i to set the api endpoint")
	rootCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "", "Set the workspace")
	rootCmd.PersistentFlags().StringVarP(&read, "read", "r", "", "Read from file (json)")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output to file ")
}
