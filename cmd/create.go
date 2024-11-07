/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources (services, routes)",
	Long:  `Use kctl create services or kctl create routes to create resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	createCmd.AddCommand(routesCreateCmd)
	createCmd.AddCommand(servicesCreateCmd)

}
