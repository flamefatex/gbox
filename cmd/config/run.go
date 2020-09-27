package config

import (
	"encoding/json"
	"fmt"

	"github.com/flamefatex/config"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	config.Config().AllSettings()

	b, err := json.MarshalIndent(config.Config().AllSettings(), "", "    ")
	if err != nil {
		fmt.Printf("json.MarshalIndent err:%s\n")
		return
	}
	fmt.Printf("%s\n", string(b))
}
