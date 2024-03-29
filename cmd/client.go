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

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Consumes the restful Quiz api ",
	Long: `
	Provides a CLI to intreact with a restful api .The api needs to be started before using this command.
	The api server can be started by calling this command with a server attribute ex:"quiz-cli server"`,
	Run: func(cmd *cobra.Command, args []string) {
		currentScore, currentQuestionIndex := 0, 1

		currentScore, err := client.ShowNextQuestion(currentQuestionIndex, currentScore)
		if err == nil {
			var percentageScore float32 = (float32(currentScore) / float32(client.TotalNumberOfQuestions)) * 100

			client.SaveScoreToFile(percentageScore)

			allUsersScores := client.GetScoresList()

			rank := client.GetScoreRanked(allUsersScores, percentageScore)

			performancePercentage := (float32(rank) / float32(len(allUsersScores))) * float32(100)

			fmt.Printf("Your score is  %.2f percent , Your rank is %v \n", percentageScore, rank)
			fmt.Printf("---- You where better then %.2f percent of all quizer ----- \n\n", performancePercentage)
		}

		// fmt.Println("client called")
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
