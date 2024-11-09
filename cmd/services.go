/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bahamas0x00/kctl/pkg/common"
	"github.com/bahamas0x00/kctl/pkg/services"
	"github.com/spf13/cobra"
)

var serviceName string

// kctl get services
// options:
//
//	--workspace		-w 		get in a workspace
//	--output 		-o		output to file
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

		return err
	},
}

// kctl create services -r [readFilePath]
// must set readFilePath
var servicesCreateCmd = &cobra.Command{
	Use:   "services",
	Short: "Create services",
	Long:  `Create services or services in a workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s services.Services
		if common.IsStringSet(read) {
			// read json from file
			data, err := os.ReadFile(read)
			if err != nil {
				return fmt.Errorf("failed to read file %s error: %v", read, err)
			}
			err = json.Unmarshal(data, &s)
			if err != nil {
				return fmt.Errorf("not json file %v", err)
			}

			// batch create services
			responses, errs := s.BatchCreateServices(apiEndpoint, workspace)
			if len(errs) > 0 {
				fmt.Println("there were some erros during create:")
				for _, err := range errs {
					return err
				}
			} else {
				for _, resp := range responses {
					fmt.Printf("Create with code:%d\n", resp.StatusCode)
				}
			}

			return err

		}

		return fmt.Errorf("read file path is not set , use -r [filepath] ")
	},
}

// kctl delete services
// options:
//
//	--workspace		-w 		delete in a workspace
//	--all 			-a  	delete all services
//	--serviceName   -sn 	Choose which service to delete
var servicesDeleteCmd = &cobra.Command{
	Use:   "services",
	Short: "Delete services",
	Long:  `Delete services or services in workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// kctl update services
// options:
//
// --workspace  	-w 		update in a workspace
// --read 			-r 		update services from json file
var servicesUpdateCmd = &cobra.Command{
	Use:   "services",
	Short: "Update services",
	Long:  `Update services or services in workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	servicesDeleteCmd.PersistentFlags().StringVarP(&serviceName, "serviceName", "s", "", "Choose which service to delete")
}
