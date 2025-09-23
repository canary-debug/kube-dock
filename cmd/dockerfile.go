// cmd/dockerfile.go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// 存储 --config 的值
var (
	configPath string
	expose     string
)

var dockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "操作 Dockerfile 文件",
	Long:  `用于修改、检查或生成 Dockerfile 的命令`,
	Run: func(cmd *cobra.Command, args []string) {
		// 检查文件是否存在
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "❌ 错误: Dockerfile 文件不存在: %s\n", configPath)
			os.Exit(1)
		}

		fmt.Printf("🔧 正在操作 Dockerfile: %s\n", configPath)

		fmt.Println("🔧 修改暴露的端口...", expose)

	},
}

func init() {
	// 添加 --config 用于指定 Dockerfile 文件位置
	dockerfileCmd.Flags().StringVarP(
		&configPath,           // 存储值的变量
		"config",              // 标志名
		"c",                   // 短选项
		"Dockerfile",          // 默认值（当前目录下的 Dockerfile）
		"指定 Dockerfile 文件的路径", // 帮助信息
	)

	// 添加 --expose 用于修改暴露的端口
	dockerfileCmd.Flags().StringVarP(
		&expose,   // 存储值的变量
		"expose",  // 标志名
		"e",       // 短选项
		"80",      // 默认值（当前目录下的 Dockerfile）
		"修改暴露的端口", // 帮助信息
	)

	// 如果你希望这个 flag 是必填的，取消下面这行注释
	//dockerfileCmd.MarkFlagRequired("config")

	// 命令注册
	rootCmd.AddCommand(dockerfileCmd)
}
