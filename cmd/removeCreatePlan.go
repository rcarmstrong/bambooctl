// Copyright © 2018 NAME Ryan Armstrong <cowboys6750@gmail.com>
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
	"log"
	"os"

	bamboo "github.com/rcarmstrong/go-bamboo"
	"github.com/spf13/cobra"
)

// removeCreatePlanCmd represents the removeCreatePlan command
var removeCreatePlanCmd = &cobra.Command{
	Use:   "removeCreatePlan",
	Short: "Remove the create plan permission for the specified role, group, or user",
	Long: `
The removeCreatePlan subcommand will remove the create plan permission for the role, group, or user on the specified project passed by the project key flag`,
	Run: func(cmd *cobra.Command, args []string) {
		if projectKeyFlag == "" {
			log.Println("You must set the key flag (-k) to specify which project's permissions will be modified.")
			cmd.Usage()
			os.Exit(1)
		}

		role, err := cmd.PersistentFlags().GetBool("role")
		if err != nil {
			panic(err)
		}
		groups, err := cmd.PersistentFlags().GetStringSlice("groups")
		if err != nil {
			panic(err)
		}
		users, err := cmd.PersistentFlags().GetStringSlice("users")
		if err != nil {
			panic(err)
		}

		if role {
			resp, err := cli.ProjectPlan.RemoveLoggedInUsersPermissions(projectKeyFlag, []string{bamboo.ReadPermission, bamboo.WritePermission, bamboo.BuildPermission, bamboo.ClonePermission})
			if err != nil {
				fmt.Printf("[%d] %s - %s", resp.StatusCode, resp.Status, err)
			}
		}

		if len(groups) > 0 {
			for _, g := range groups {
				resp, err := cli.ProjectPlan.RemoveGroupPermissions(projectKeyFlag, g, []string{bamboo.ReadPermission, bamboo.WritePermission, bamboo.BuildPermission, bamboo.ClonePermission})
				if err != nil {
					fmt.Printf("[%d] %s - %s", resp.StatusCode, resp.Status, err)
				}
				fmt.Println(g)
			}
		}

		if len(users) > 0 {
			for _, u := range users {
				resp, err := cli.ProjectPlan.RemoveUserPermissions(projectKeyFlag, u, []string{bamboo.ReadPermission, bamboo.WritePermission, bamboo.BuildPermission, bamboo.ClonePermission})
				if err != nil {
					fmt.Printf("[%d] %s - %s", resp.StatusCode, resp.Status, err)
				}
			}
		}
	},
}

func init() {
	permissionsCmd.AddCommand(removeCreatePlanCmd)

	removeCreatePlanCmd.PersistentFlags().BoolP("role", "r", false, fmt.Sprintf("Defaults to %s", bamboo.LoggedInRole))
	removeCreatePlanCmd.PersistentFlags().StringSliceP("groups", "g", []string{}, "A single or comma seperated list of groups to remove the create plan permission to.")
	removeCreatePlanCmd.PersistentFlags().StringSliceP("users", "u", []string{}, "A single or comma seperated list of users to remove the create plan permission to.")
}
