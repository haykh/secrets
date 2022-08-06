package cobra

import (
	"fmt"

	"github.com/haykh/secrets"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets the value associated with the given key from secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secrets.File(encodingKey, secretsPath())
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Println("No value set for this key")
			return
		}
		fmt.Printf("%s=%s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
