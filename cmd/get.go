package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"keyvaluecli/models"
	"log"
	"net/http"
)

var get = &cobra.Command{
	Use:   "get",
	Short: "get",
	Long:  "get",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for the greet command
	},
}

var getKey = &cobra.Command{
	Use:   "key <key>",
	Short: "Get Secret from Vault",
	Long:  "Get Secret from Vault.",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for the greet command
		resp, err := http.Get("http://localhost:8089/keys/" + args[0] + "?showSecret=True")
		if err != nil {
			log.Fatalln(err)
		}

		//We Read the response body on the line below.
		body, err := io.ReadAll(resp.Body)
		res := models.Response{}
		err = json.Unmarshal(body, &res)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		fmt.Println(res.Data.Value)
	},
}

var getKeys = &cobra.Command{
	Use:   "keys",
	Short: "Get Secret from Vault",
	Long:  "Get Secret from Vault.",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for the greet command
		resp, err := http.Get("http://localhost:8089/keys")
		if err != nil {
			log.Fatalln(err)
		}

		//We Read the response body on the line below.
		body, err := io.ReadAll(resp.Body)
		res := models.ResponseList{}
		err = json.Unmarshal(body, &res)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert the body to type string
		fmt.Println(res.Data)
	},
}
