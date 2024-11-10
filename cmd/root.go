/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	apiEndpoint string
	workspace   string
	filePath    string
	serviceName string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kctl",
	Short: "CLI for Kong Gateway Admin api management",
	Long:  `CLI for Kong Gateway Admin api management`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.PersistentFlags().StringVarP(&apiEndpoint, "api", "i", "http://localhost:8001", "Use -i to set the api endpoint")
	rootCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "", "Set the workspace")
}
