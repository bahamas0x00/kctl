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

		resp, err = services.ListAllServices(apiEndpoint, workspace)
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

/*
package main

import (
	"fmt"
	"github.com/bahamas0x00/kctl/pkg/services"
	"github.com/bahamas0x00/kctl/pkg/common"
)

func main() {
	apiEndpoint := "http://localhost:8001" // Kong Admin API 地址
	workspace := "my-workspace"

	// 示例服务数据，假设需要更新的服务
	servicesToUpdate := []services.Service{
		{Name: "service1", Host: "localhost", Port: 8081, Enabled: true},
		{Name: "service2", Host: "localhost", Port: 8082, Enabled: false},
		{Name: "service3", Host: "localhost", Port: 8083, Enabled: true},
	}

	// 创建 Services 对象
	serviceObj := services.Services{Data: servicesToUpdate}

	// 批量更新服务
	responses, errs := serviceObj.BatchUpdateServices(apiEndpoint)
	if len(errs) > 0 {
		fmt.Println("There were some errors during update:")
		for _, err := range errs {
			fmt.Println(err)
		}
	} else {
		// 打印每个服务的更新响应
		for _, resp := range responses {
			fmt.Printf("Service updated with status code: %d\n", resp.StatusCode)
		}
	}
}

*/
