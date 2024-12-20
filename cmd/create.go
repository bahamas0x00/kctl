/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources (services, routes...)",
	Long:  `Create resources (services, routes...)`,
}

func init() {
	createCmd.AddCommand(routesCreateCmd)
	createCmd.AddCommand(servicesCreateCmd)
	createCmd.AddCommand(upstreamsCreateCmd)
	createCmd.AddCommand(targetsCreateCmd)
	createCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Read from file (json)")

}
