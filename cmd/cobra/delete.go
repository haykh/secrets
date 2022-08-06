package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes a key value pair from secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := goencode.File(encodingKey, secretsPath())
		key := args[0]
		err := v.Delete(key)
		if err != nil {
			fmt.Println("No value set for this key")
			return
		}
		fmt.Println("Key successfully deleted")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
