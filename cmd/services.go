/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/bahamas0x00/kctl/pkg/common"
	"github.com/bahamas0x00/kctl/pkg/services"
	"github.com/spf13/cobra"
)

var servicesGetCmd = &cobra.Command{
	Use:   "services",
	Short: "Get services",
	Long:  `Get services or services in a workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var resp *common.HttpResponse
		var err error
		if common.IsStringSet(workspace) {
			resp, err = services.ListAllServicesInWorkspace(apiEndpoint, workspace)
		} else {
			resp, err = services.ListAllServices(apiEndpoint)
		}

		if err != nil {
			return fmt.Errorf("failed to get services: %v", err)
		}

		// if output flag is set , write the content to file
		if common.IsStringSet(output) {
			err := common.SaveResponseToFile(resp, output)
			if err != nil {
				return fmt.Errorf("failed to write to file: %v", err)
			}
		} else {
			// 如果没有设置 output 参数，则打印到控制台
			fmt.Printf("Response:\n%s", resp.Body)
		}

		return nil
	},
}

var servicesCreateCmd = &cobra.Command{
	Use:   "services",
	Short: "Create services",
	Long:  `Create services or services in a workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func init() {

}
