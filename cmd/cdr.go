/*
Copyright © 2024 Hao Li <mr.hao.li@gmail.com>

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/haoli000/tttns/cdr"
	"github.com/spf13/cobra"
)

// cdrCmd represents the cdr command
var cdrCmd = &cobra.Command{
	Use:   "cdr [file|-]",
	Short: "Print all CDR header info",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		content := cdr.GetContent(args[0])
		if jsonOutput, _ := cmd.Flags().GetBool("json"); jsonOutput {
			cdrInfo := cdr.ToCdrInfo(content)
			jsonBytes, err := json.MarshalIndent(cdrInfo, "", "    ")
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Println(string(jsonBytes))
		} else {
			cnt := cdr.CountCdrs(content)
			fmt.Printf("Number of CDRs: %d\n", cnt)
			for i := uint32(1); i <= cnt; i++ {
				info := cdr.ToCdrHeaderInfo(content, i)
				cdr.PrintCdrHeaderInfo(info)
			}
		}
	},
}

func init() {
	cdrCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
	rootCmd.AddCommand(cdrCmd)
}
