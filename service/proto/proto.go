package proto

import (
	"github.com/flamefatex/log"
	"github.com/spf13/cobra"
)

type Param struct {
	Src string
}

func Do(cmd *cobra.Command, args []string, param *Param) {
	// todo

	log.Infof("dd %s", "dd")
}
