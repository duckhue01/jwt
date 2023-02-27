/*
Copyright © 2023 DK duckhue01.tech@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/duckhue01/jwt/jwt"
	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
)

const (
	REQUIRE_TOKEN_ERROR = "this command require a JWT token"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode a JSON web token",
	Long:  "Decode a JSON web token",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			token, err := jwt.Decode(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			header, err := prettyjson.Marshal(token.Header)
			if err != nil {
				fmt.Println(err)
				return
			}

			body, err := prettyjson.Marshal(token.Claims)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Header(algorithms & token type...):\n%s\n\n", header)
			fmt.Printf("Body(claims):\n%s\n", body)

			return
		}
		fmt.Println(REQUIRE_TOKEN_ERROR)
	},
}

func init() {
	decodeCmd.AddCommand()

	rootCmd.AddCommand(decodeCmd)

	// decodeCmd.Flags().String("toggle", "t", false, "Help message for toggle")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")
}
