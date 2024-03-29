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
	"quiz-cli/client"

	"github.com/spf13/cobra"
)

// scoresCmd represents the scores command
var scoresCmd = &cobra.Command{
	Use:   "scores",
	Short: "Get a list of all the stored scores",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scores called")

		savedScores := client.GetScoresList()
		if len(savedScores) > 0 {
			for index, score := range savedScores {
				fmt.Printf("Rank %v Score %v \n", index+1, score)
			}
		} else {
			fmt.Println("The quiz has never been played :( ")
		}
	},
}

func init() {
	rootCmd.AddCommand(scoresCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scoresCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scoresCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
