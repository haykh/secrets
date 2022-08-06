package cobra

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all keys with associated values in secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := goencode.File(encodingKey, secretsPath())
		keys, err := v.List()
		if err != nil {
			fmt.Println("Failed to retrieve keys")
			return
		}
		if len(keys) == 0 {
			fmt.Println("There are currently no key value pairs in secret storage")
			return
		}
		keysRep := strings.Join(keys, "\n- ")
		fmt.Print("Keys:\n- ")
		fmt.Println(keysRep)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
