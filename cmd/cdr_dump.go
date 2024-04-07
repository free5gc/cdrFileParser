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
	"fmt"
	"os"
	"strconv"

	"github.com/haoli000/tttns/cdr"
	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump [input_file|-] [index] [output_file]",
	Short: "Dump the row content of CDR to output file or stdout",
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.ParseUint(args[1], 10, 32)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(4)
		}
		if index < 1 || index > (1<<32)-1 {
			fmt.Println("Error: Index must be an integer starting from 1")
			os.Exit(4)
		}
		content := cdr.GetContent(args[0])

		if len(args) == 2 {
			cdr.DumpCdr(content, uint32(index), os.Stdout)
		} else if len(args) == 3 {
			filename := args[2]
			file, err := os.Create(filename)
			if err != nil {
				fmt.Println("Error opening file:", err)
				os.Exit(4)
			}
			defer file.Close()
			cdr.DumpCdr(content, uint32(index), file)
		}
	},
}

func init() {
	dumpCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
	cdrCmd.AddCommand(dumpCmd)
}
