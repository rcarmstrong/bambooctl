// Copyright © 2018 Ryan Armstrong <cowboys6750@gmail.com>
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

	"github.com/spf13/cobra"
)

var projectKeyFlag string

// permissionsCmd represents the permissions command
var permissionsCmd = &cobra.Command{
	Use:   "permissions",
	Short: "Configure project level permissions.",
	Long: `
Add or remove project level permissions for individual users or groups.`,
}

func init() {
	projectCmd.AddCommand(permissionsCmd)

	permissionsCmd.PersistentFlags().StringVarP(&projectKeyFlag, "projectKey", "k", "", "Specifies the key of the project to operate on.")

	permissionsCmd.Run = func(cmd *cobra.Command, args []string) {
		if projectKeyFlag != "" {
			outputPermissions(projectKeyFlag)
		} else {
			log.Println("Getting projects keys...")
			projects, response, err := cli.Projects.ListProjects()
			if err != nil {
				fmt.Printf("[%d] Bamboo returned %s when getting projects to list permissions: %s\n", response.StatusCode, response.Status, err)
				os.Exit(1)
			}

			for _, p := range projects {
				outputPermissions(p.Key)
			}
		}
	}
}

func outputPermissions(key string) {
	userPermissions, response, err := cli.ProjectPlan.UserPermissionsList(key)
	if err != nil {
		fmt.Printf("[%d] Bamboo returned %s when listing user permissions: %s\n", response.StatusCode, response.Status, err)
		os.Exit(1)
	}
	if len(userPermissions) != 0 {
		fmt.Println("User Permissions:")
		for _, u := range userPermissions {
			fmt.Println(" ", u.Name)
			for _, p := range u.Permissions {
				fmt.Println("   ", p)
			}
		}
	} else {
		fmt.Printf("%s has no user permissions configured\n", key)
	}

	groupPermissions, response, err := cli.ProjectPlan.GroupPermissionsList(key)
	if err != nil {
		fmt.Printf("[%d] Bamboo returned %s when listing group permissions: %s\n", response.StatusCode, response.Status, err)
		os.Exit(1)
	}
	if len(groupPermissions) != 0 {
		fmt.Println("Group Permissions:")
		for _, g := range groupPermissions {
			fmt.Println(" ", g.Name)
			for _, p := range g.Permissions {
				fmt.Println("   ", p)
			}
		}
	} else {
		fmt.Printf("%s has no group permissions configured\n", key)
	}

	rolePermissions, response, err := cli.ProjectPlan.RolePermissionsList(key)
	if err != nil {
		fmt.Printf("[%d] Bamboo returned %s when listing role permissions: %s\n", response.StatusCode, response.Status, err)
		os.Exit(1)
	}
	if len(rolePermissions) != 0 {
		fmt.Println("Role Permissions:")
		for _, r := range rolePermissions {
			fmt.Println(" ", r.Name)
			for _, p := range r.Permissions {
				fmt.Println("   ", p)
			}
		}
	} else {
		fmt.Printf("%s has no role permissions configured\n", key)
	}
}
