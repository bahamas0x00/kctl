/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resources (services, routes)",
	Long:  `Use kctl get services or kctl get routes to fetch resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	getCmd.AddCommand(routesGetCmd)
	getCmd.AddCommand(servicesGetCmd)
}
