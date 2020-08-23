/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// zombieCmd represents the zombie command
var zombieCmd = &cobra.Command{
	Use:   "zombie",
	Short: "example of a nongraceful exit when signal passed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zombie called (ctrl+c twice to exit)")
		i := 1

		// this process does not handle ctx.Done()
		for {
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	},
}

func init() {
	rootCmd.AddCommand(zombieCmd)
}
