/*
Copyright © 2023 NAME HERE duckhue01.tech@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/duckhue01/jwt/jwt"
	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a JSON web token",
	Long:  "Verify a JSON web token",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			jwt.VerifyHandler(args[0], cmd.Flags())
			return
		}
		fmt.Println(REQUIRE_TOKEN_ERROR)

	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	verifyCmd.Flags().Bool(jwt.IsExpireFlag, false, "Verify if token is expire or not")
	verifyCmd.Flags().StringSlice(jwt.ContainFlag, make([]string, 0), "Verify if Token contains specified claims")
	// todo: implement later
	// verifyCmd.Flags().Bool(jwt.IsValidFlag, false, "Verify if token is a valid token or not. It means that token is issued from trusted party and doesn't change")

}
