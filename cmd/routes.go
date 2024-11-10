/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"

	"github.com/bahamas0x00/kctl/pkg/common"
	"github.com/bahamas0x00/kctl/pkg/routes"
	"github.com/spf13/cobra"
)

var routesGetCmd = &cobra.Command{
	Use:   "routes",
	Short: "Get routes",
	Long:  `Get routes `,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := routes.ListAllRoutes(apiEndpoint, workspace, serviceName)
		if err != nil {
			return fmt.Errorf("failed to get services: %v", err)
		}
		defer resp.Body.Close()

		// if output flag is set , write the content to file
		if common.IsStringSet(output) {
			err := common.SaveResponseToFile(resp, output)
			if err != nil {
				return fmt.Errorf("failed to write to file: %v", err)
			}
		} else {
			// 如果没有设置 output 参数，则打印到控制台
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read data: %v", err)
			}
			fmt.Printf("Response:\n%s\n", string(data))
		}
		fmt.Printf("workspace: %s\n", workspace)
		fmt.Printf("associated service name: %s\n", serviceName)
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
	routesGetCmd.PersistentFlags().StringVarP(&serviceName, "service", "s", "", "set service name")
	routesCreateCmd.PersistentFlags().StringVarP(&serviceName, "service", "s", "", "set service name")
	routesDeleteCmd.PersistentFlags().StringVarP(&serviceName, "service", "s", "", "set service name")
	routesUpdateCmd.PersistentFlags().StringVarP(&serviceName, "service", "s", "", "set service name")
}
