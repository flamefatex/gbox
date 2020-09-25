package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

type VersionInfo struct {
	ServiceName string
	Version     string
	GitCommit   string
}

var versionInfo *VersionInfo

func SetVersionInfo(vi *VersionInfo) {
	versionInfo = vi
}

func Do(cmd *cobra.Command, args []string) {
	fmt.Printf("ServiceName: %s\n", versionInfo.ServiceName)
	fmt.Printf("Version:     %s\n", versionInfo.Version)
	fmt.Printf("GitCommit:   %s\n", versionInfo.GitCommit)
}
