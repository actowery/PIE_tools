package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// initCmd is a subcommand to StoreCmd that ads a Benchmark to the store.
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an object from the SN CMDB",
	Long: `Facilitates the execution of delete to SN CMDB.

	nodes/change/relationship/ actions are supported.
	Example usage:
	  SNHttpClient delete <action>
		`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sending Delete request")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		fmt.Println("Endpoint: " + endpoint)
		username, _ := cmd.Flags().GetString("username")
		fmt.Println("Username: " + username)
		password, _ := cmd.Flags().GetString("password")
		fmt.Println("Password: " + password)
		fmt.Println(DeleteRecord(endpoint, username, password))
	},

	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// ScannerVersion is the version of the scanner associated with the benchmark.

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(nodesCmd)
}

func DeleteRecord(endpoint string, username string, password string) string {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", endpoint, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	x := string(bodyText)
	return x
}
