package cobra

import (
	"fmt"

	"github.com/haykh/secrets"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "sets a key value pair in secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secrets.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			fmt.Println("Failed to set key value pair")
		}
		fmt.Println("Key value pair successfully set:")
		fmt.Printf("%s=%s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
