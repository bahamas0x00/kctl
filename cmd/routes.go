/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

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
		var r routes.Routes
		if common.IsStringSet(filePath) {
			// read json from file
			data, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s error: %v", filePath, err)
			}
			err = json.Unmarshal(data, &r)
			if err != nil {
				return fmt.Errorf("not json file %v", err)
			}

			// batch create services
			_, errs := r.BatchCreateRoutes(apiEndpoint, workspace, serviceName)
			if len(errs) > 0 {
				fmt.Println("there were some erros during create:")
				for _, err := range errs {
					return err
				}
			}
			fmt.Printf("workspace: %s\n", workspace)
			fmt.Printf("associated service name: %s\n", serviceName)
			return nil

		}

		return fmt.Errorf("invalid command")
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
