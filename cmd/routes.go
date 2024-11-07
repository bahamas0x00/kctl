/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// routesCmd represents the routes command
var routesGetCmd = &cobra.Command{
	Use:   "routes",
	Short: "Get routes",
	Long:  `Get routes `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("routes called")
	},
}

func init() {

}
