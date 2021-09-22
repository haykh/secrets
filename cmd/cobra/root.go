package cobra

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is a manager for API keys and other secrets",
}

var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "",
		"the key to use when encoding and decoding secrets")
}

func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}
