/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var read string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources (services, routes...)",
	Long:  `Create resources (services, routes...)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choose resources ")
	},
}

func init() {
	createCmd.AddCommand(routesCreateCmd)
	createCmd.AddCommand(servicesCreateCmd)
	createCmd.PersistentFlags().StringVarP(&read, "read", "r", "", "Read from file (json)")

}
