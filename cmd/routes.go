/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var routesGetCmd = &cobra.Command{
	Use:   "routes",
	Short: "Get routes",
	Long:  `Get routes `,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var routesCreateCmd = &cobra.Command{
	Use:   "routes",
	Short: "Create routes",
	Long:  `Create routes `,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var routesDeleteCmd = &cobra.Command{
	Use:   "routes",
	Short: "Delete routes",
	Long:  `Delete routes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var routesUpdateCmd = &cobra.Command{
	Use:   "routes",
	Short: "Update routes",
	Long:  `Update routes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {

}
