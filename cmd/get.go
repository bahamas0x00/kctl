/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resources (services, routes...)",
	Long:  `Get resources (services, routes...)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choose resources ")
	},
}

func init() {
	getCmd.AddCommand(routesGetCmd)
	getCmd.AddCommand(servicesGetCmd)
}
