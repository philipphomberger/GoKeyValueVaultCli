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

var set = &cobra.Command{
	Use:   "set",
	Short: "set",
	Long:  "set",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for the greet command
	},
}

var putKey = &cobra.Command{
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
		req, err := http.NewRequest(http.MethodPut, "http://localhost:8089/keys/"+args[0], body)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(resp.StatusCode)
	},
}
