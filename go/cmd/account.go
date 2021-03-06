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
	"log"

	"github.com/atoonk/mysocketctl/go/internal/http"
	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use: "account",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new account or see account information",
	Run: func(cmd *cobra.Command, args []string) {
		err := http.Register(name, email, password, sshkey)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Println("account created, please check your email")
	},
}

func init() {

	createCmd.Flags().StringVarP(&email, "email", "e", "", "Email address")
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Your name")
	createCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	createCmd.Flags().StringVarP(&sshkey, "sshkey", "s", "", "SSH Key")
	createCmd.MarkFlagRequired("email")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("password")
	createCmd.MarkFlagRequired("sshkey")

	accountCmd.AddCommand(createCmd)
	rootCmd.AddCommand(accountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
