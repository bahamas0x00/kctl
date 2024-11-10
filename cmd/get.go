/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var output string

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resources (services, routes...)",
	Long:  `Get resources (services, routes...)`,
}

func init() {
	getCmd.AddCommand(routesGetCmd)
	getCmd.AddCommand(servicesGetCmd)
	getCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output to file ")
}
