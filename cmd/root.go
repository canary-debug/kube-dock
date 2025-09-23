// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kube-dock",
	Short: "kube-dock 是一个用于修改 Dockerfile 和 Kubernetes YAML 的内部工具",
	Long:  `kube-dock 是公司内部开发运维工具，用于自动化修改容器配置文件.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// 可选：添加全局 flag
	// rootCmd.PersistentFlags().BoolP("verbose", "v", false, "启用详细日志")

	// 可选：全局运行前检查
	// rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
	//     if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
	//         fmt.Println("运行模式: 详细输出")
	//     }
	// }
}
