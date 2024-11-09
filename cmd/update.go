/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update resources (services, routes...)",
	Long:  `Update resources (services, routes...)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choose resources ")
	},
}

func init() {
	updateCmd.AddCommand(routesUpdateCmd)
	updateCmd.AddCommand(servicesUpdateCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
