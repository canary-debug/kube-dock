// cmd/yaml.go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "操作 Kubernetes YAML 文件",
	Long:  `用于修改、检查或更新 Kubernetes 配置文件的命令`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Yaml")
	},
}

func init() {
	rootCmd.AddCommand(yamlCmd)
}
