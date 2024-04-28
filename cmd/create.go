package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"keyvaluecli/models"
	"log"
	"net/http"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "c",
	Long:  "create",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for the greet command
	},
}

var createKey = &cobra.Command{
	Use:   "key <key> <value>",
	Short: "Get Secret from Vault",
	Long:  "Get Secret from Vault.",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for the greet command
		var model models.Secret
		model = models.Secret{
			Key:   args[0],
			Value: args[1],
		}
		body := &bytes.Buffer{}
		enc := json.NewEncoder(body)
		err := enc.Encode(model)
		if err != nil {
			return
		}
		resp, err := http.Post("http://localhost:8089/keys/",
			"application/json", body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(resp.StatusCode)
	},
}
