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
	"github.com/bahamas0x00/kctl/pkg/upstreams"
	"github.com/spf13/cobra"
)

// kctl get upstreams
// options:
//
//	--workspace		-w 		get in a workspace
//	--output 		-o		output to file
var upstreamsGetCmd = &cobra.Command{
	Use:   "upstreams",
	Short: "Get upstreams",
	Long:  `Get upstreams or upstreams in a workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := upstreams.ListAllUpstreams(apiEndpoint, workspace)
		if err != nil {
			return fmt.Errorf("failed to get upstreams: %v", err)
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

// kctl create upstreams -r [readFilePath]
// must set readFilePath
var upstreamsCreateCmd = &cobra.Command{
	Use:   "upstreams",
	Short: "Create upstreams",
	Long:  `Create upstreams or upstreams in a workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s upstreams.Upstreams
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

			// batch create upstreams
			_, errs := s.BatchCreateUpstreams(apiEndpoint, workspace)
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

// kctl delete upstreams
// options:
//
//	--workspace		-w 		delete in a workspace
//	--read          -r      delete list of upstreams from file
var upstreamsDeleteCmd = &cobra.Command{
	Use:   "upstreams",
	Short: "Delete upstreams",
	Long:  `Delete upstreams or upstreams in workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s upstreams.Upstreams
		if common.IsStringSet(filePath) {
			data, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s error: %v", filePath, err)
			}
			err = json.Unmarshal(data, &s)
			if err != nil {
				return fmt.Errorf("not json file %v", err)
			}

			// batch delete upstreams
			_, errs := s.BatchDeleteUpstreams(apiEndpoint, workspace)
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

// kctl update upstreams
// options:
//
// --workspace  	-w 		update in a workspace
// --read 			-r 		update upstreams from json file
var upstreamsUpdateCmd = &cobra.Command{
	Use:   "upstreams",
	Short: "Update upstreams",
	Long:  `Update upstreams or upstreams in workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s upstreams.Upstreams
		if common.IsStringSet(filePath) {
			data, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s error: %v", filePath, err)
			}
			err = json.Unmarshal(data, &s)
			if err != nil {
				return fmt.Errorf("not json file %v", err)
			}

			// batch update upstreams
			_, errs := s.BatchUpdateUpstreams(apiEndpoint, workspace)
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

}
