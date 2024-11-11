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
	"github.com/bahamas0x00/kctl/pkg/targets"
	"github.com/spf13/cobra"
)

// kctl get targets
// options:
//
//	-w 	-u 	get in a workspace associated with an upstream
//	--output 		-o		output to file
var targetsGetCmd = &cobra.Command{
	Use:   "targets",
	Short: "Get targets associated with an upstream",
	Long:  `Get targets associated with an upstream or targets in a workspace associated with an upstream`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := targets.ListAllTargets(apiEndpoint, workspace, upstreamName)
		if err != nil {
			return fmt.Errorf("failed to get targets: %v", err)
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
		fmt.Printf("workspace: %s", workspace)
		return err
	},
}

// kctl create targets -r [readFilePath]
// must set readFilePath
var targetsCreateCmd = &cobra.Command{
	Use:   "targets",
	Short: "Create targets associated with an upstream",
	Long:  `Create targets associated with an upstream or targets in a workspace associated with an upstream`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s targets.Targets
		if common.IsStringSet(filePath) {
			// read json from file
			data, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s error: %v", filePath, err)
			}
			err = json.Unmarshal(data, &s)
			if err != nil {
				return fmt.Errorf("not json file %v", err)
			}

			// batch create targets
			_, errs := s.BatchCreateTargets(apiEndpoint, workspace, upstreamName)
			if len(errs) > 0 {
				fmt.Println("there were some erros during create:")
				for _, err := range errs {
					return err
				}
			}
			fmt.Printf("workspace: %s", workspace)
			return nil

		}

		return fmt.Errorf("invalid command")
	},
}

// kctl delete targets
// options:
//
//	--workspace		-w 		delete in a workspace
//	--read          -r      delete list of targets from file
var targetsDeleteCmd = &cobra.Command{
	Use:   "targets",
	Short: "Delete targets associated with an upstream",
	Long:  `Delete targets associated with an upstream or targets in a workspace associated with an upstream`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s targets.Targets
		if common.IsStringSet(filePath) {
			data, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s error: %v", filePath, err)
			}
			err = json.Unmarshal(data, &s)
			if err != nil {
				return fmt.Errorf("not json file %v", err)
			}

			// batch delete targets
			_, errs := s.BatchDeleteTargets(apiEndpoint, workspace, upstreamName)
			if len(errs) > 0 {
				fmt.Println("there were some erros during delete:")
				for _, err := range errs {
					return err
				}
			}

			fmt.Printf("workspace: %s", workspace)

			return nil

		}
		return fmt.Errorf("invalid command")
	},
}

// kctl update targets
// options:
//
// --workspace  	-w 		update in a workspace
// --read 			-r 		update targets from json file
var targetsUpdateCmd = &cobra.Command{
	Use:   "targets",
	Short: "Update targets associated with an upstream",
	Long:  `Update targets associated with an upstream or targets in a workspace associated with an upstream`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s targets.Targets
		if common.IsStringSet(filePath) {
			data, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s error: %v", filePath, err)
			}
			err = json.Unmarshal(data, &s)
			if err != nil {
				return fmt.Errorf("not json file %v", err)
			}

			// batch update targets
			_, errs := s.BatchUpdateTargets(apiEndpoint, workspace, upstreamName)
			if len(errs) > 0 {
				fmt.Println("There were some erros during update:")
				for _, err := range errs {
					return err
				}
			}

			fmt.Printf("workspace: %s", workspace)

			return nil

		}
		return fmt.Errorf("invalid command")
	},
}

func init() {
	targetsGetCmd.PersistentFlags().StringVarP(&upstreamName, "upstream", "u", "", "set upstream name")
	targetsCreateCmd.PersistentFlags().StringVarP(&upstreamName, "upstream", "u", "", "set upstream name")
	targetsDeleteCmd.PersistentFlags().StringVarP(&upstreamName, "upstream", "u", "", "set upstream name")
	targetsUpdateCmd.PersistentFlags().StringVarP(&upstreamName, "upstream", "u", "", "set upstream name")
}
