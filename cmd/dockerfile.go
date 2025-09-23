// cmd/dockerfile.go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var dockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "操作 Dockerfile 文件",
	Long:  `用于修改、检查或生成 Dockerfile 的命令`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Docker")
	},
}

func init() {
	rootCmd.AddCommand(dockerfileCmd)
}
