/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/bahamas/kctl/common"
	"github.com/spf13/cobra"
)

var serviceName, workspace string

// routesCmd represents the routes command
var routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			command: kctl get routes -s service -w workspace

		*/
		common.IsStringSet(workspace)

	},
}

func init() {
	getCmd.AddCommand(routesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routesCmd.PersistentFlags().String("foo", "", "A help for foo")
	routesCmd.PersistentFlags().StringVarP(&serviceName, "service", "s", "", "use -s to set the service name ")
	routesCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "default", "use -w to set the workspace name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
