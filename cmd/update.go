/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update resources (services, routes...)",
	Long:  `Update resources (services, routes...)`,
}

func init() {
	updateCmd.AddCommand(routesUpdateCmd)
	updateCmd.AddCommand(servicesUpdateCmd)
	updateCmd.AddCommand(upstreamsUpdateCmd)
	updateCmd.AddCommand(targetsUpdateCmd)
	updateCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Read from file (json)")
}
