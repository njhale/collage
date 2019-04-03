// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type glueOptions struct {
	squash    bool
	out       string
	rawImages string
}

var (
	// glueOpts represents the options available
	glueOpts = &glueOptions{}

	// glueCmd represents the glue command
	glueCmd = &cobra.Command{
		Use:   "glue",
		Short: "<glue-short-description>",
		Long:  `<glue-long-description>`,
		// TODO: add argument validation
		RunE: glueOpts.run,
	}
)

func (opts *glueOptions) run(cmd *cobra.Command, args []string) error {
	fmt.Printf("glue called: %v\n", args)
	// TODO: Parse arguments into image tokens
	// TODO: Get image manifests
	// TODO: Select layers
	// TODO: Pull selected layers
	// TODO: Create new manifest with selected layers
	// TODO: Push image

	return nil
}

func init() {
	rootCmd.AddCommand(glueCmd)

	// Set flags
	glueCmd.PersistentFlags().BoolVarP(&glueOpts.squash, "squash", "s", false, "squash layers")
	glueCmd.PersistentFlags().StringVarP(&glueOpts.out, "out", "o", "", "output image name")
	// TODO: Mark output image required?
	// Flags and config settings

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// glueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// glueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
