/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resources (services, routes...)",
	Long:  `Delete resources (services, routes...)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choose resources ")
	},
}

func init() {
	deleteCmd.AddCommand(routesDeleteCmd)
	deleteCmd.AddCommand(servicesDeleteCmd)
	deleteCmd.PersistentFlags().StringVarP(&read, "read", "r", "", "read from file")
}
