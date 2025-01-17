/*
 Copyright 2021 Qiniu Cloud (qiniu.com)
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
	cli "github.com/ar0c/goc/v2/pkg/watch"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:     "watch",
	Short:   "watch for profile's real time update",
	Long:    "watch for profile's real time update",
	Example: "",

	Run: watch,
}

var (
	watchHost string
)

func init() {
	watchCmd.Flags().StringVarP(&watchHost, "host", "", "127.0.0.1:7777", "specify the host of the goc server")
	rootCmd.AddCommand(watchCmd)
}

func watch(cmd *cobra.Command, args []string) {
	cli.Watch(watchHost)
}
