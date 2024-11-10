/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resources (services, routes...)",
	Long:  `Delete resources (services, routes...)`,
}

func init() {
	deleteCmd.AddCommand(routesDeleteCmd)
	deleteCmd.AddCommand(servicesDeleteCmd)
	deleteCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Read from file (json)")
}
