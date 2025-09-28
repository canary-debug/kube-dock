package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"strings"
)

/*
存储 --config 的值
存储 --expose 的值
存储 --addenv 的值
存储 --add 的值
存储 --copyfile 的值
*/
var (
	configPath string
	expose     string
	addenv     []string
	add        []string
	copyfile   []string
)

var dockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "操作 Dockerfile 文件",
	Long:  `用于修改、检查或生成 Dockerfile 的命令`,
	Run: func(cmd *cobra.Command, args []string) {
		// 参数检查
		if len(os.Args) <= 2 {
			fmt.Println("❌ 错误：参数不足，至少需要指定 Dockerfile 配置文件")
			fmt.Println("🚀用法: ./kube-dock --config <Dockerfile> --expose <port>")
			return
		}

		// 检查文件是否存在
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("❌ 错误: Dockerfile 文件不存在: %s\n", configPath)
			os.Exit(1)
		}

		/*
			读取配置文件内容
		*/
		file, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Println("❌ 读取配置文件错误:", err)
			return
		}

		// =========================================================
		//  1. EXPOSE(只有当 --expose 参数不为空时才执行)
		// =========================================================
		if expose != "" {
			//正则匹配 EXPOSE [端口号]
			pattern := `EXPOSE\s+(\d+)`
			regex := regexp.MustCompile(pattern)

			/*
				匹配文件当中的 EXPOSE [端口] 字段
				strings.Join 把 matches 数组中的元素用空格连接起来(转换为 staring 类型)
				strings.Replace 把 matches 匹配到的内容替换为指定字段
			*/
			matches := regex.FindAllString(string(file), -1)
			// 检测文件中是否有多个 EXPOSE 字段
			if len(matches) >= 2 {
				fmt.Println("❌ 错误：Dockerfile 中有多个 EXPOSE 字段请手动修改")
				return
			}
			result := strings.Join(matches, " ")
			cleaned := strings.Replace(string(file), result, "EXPOSE "+expose, -1)

			// 写入文件
			err = os.WriteFile(configPath, []byte(cleaned), 0644)
			if err != nil {
				fmt.Println("❌ 写入文件错误:", err)
				return
			}
			fmt.Println("🚀EXPOSE字段修改成功:", expose)
		}

		// =========================================================
		//  2. ENV(只有当 --env 参数不为空时才执行)
		// =========================================================
		if strings.Join(addenv, "") != "" {
			// 以读写追加的形式打开文件
			file, err := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("❌ 读取配置文件错误:", err)
				return
			}
			defer file.Close()

			// 遍历 --env 参数
			for _, env := range addenv {
				if parts := strings.SplitN(env, "=", 2); len(parts) == 2 {
					key, value := parts[0], parts[1]

					// 文件写入
					fmtenv := fmt.Sprintf("\nENV %s %s", key, value)
					_, err := file.WriteString(fmtenv)
					if err != nil {
						fmt.Println("❌ 写入文件错误:", err)
						return
					}
				}
			}
		}

		// =========================================================
		//  3. ADD(只有当 --add 参数不为空时才执行)
		// =========================================================
		if strings.Join(add, "") != "" {
			// 以读写追加的形式打开文件
			file, err := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("❌ 读取配置文件错误:", err)
				return
			}
			defer file.Close()

			// 遍历 --add 添加的参数
			for _, add := range add {
				if parts := strings.SplitN(add, "=", 2); len(parts) == 2 {
					key, value := parts[0], parts[1]

					// 文件写入
					fmtadd := fmt.Sprintf("\nADD %s %s", key, value)
					_, err := file.WriteString(fmtadd)
					if err != nil {
						fmt.Println("❌ 写入文件错误:", err)
						return
					}
				}
			}
		}

		// =========================================================
		//  3. COPY(只有当 --copyfile 参数不为空时才执行)
		// =========================================================
		if strings.Join(copyfile, "") != "" {
			// 以读写追加的形式打开文件
			file, err := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("❌ 读取配置文件错误:", err)
				return
			}
			defer file.Close()

			// 遍历 --add 添加的参数
			for _, copyfile := range copyfile {
				if parts := strings.SplitN(copyfile, "=", 2); len(parts) == 2 {
					key, value := parts[0], parts[1]

					// 文件写入
					fmtcopy := fmt.Sprintf("\nCOPY %s %s", key, value)
					_, err := file.WriteString(fmtcopy)
					if err != nil {
						fmt.Println("❌ 写入文件错误:", err)
						return
					}
				}
			}
		}

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
		"x",       // 短选项
		"",        // 默认值（当前目录下的 Dockerfile）
		"修改暴露的端口", // 帮助信息
	)

	// 添加 --env 用于添加系统环境变量
	dockerfileCmd.Flags().StringSliceVarP(
		&addenv,
		"addenv",
		"e",
		nil,
		"添加环境变量，格式: --addenv KEY=VALUE",
	)

	// 添加 --add 用于添加系统环境变量
	dockerfileCmd.Flags().StringSliceVarP(
		&add,
		"add",
		"a",
		nil,
		"添加ADD，格式: --add KEY=VALUE",
	)

	// 添加 --add 用于添加系统环境变量
	dockerfileCmd.Flags().StringSliceVarP(
		&copyfile,
		"copyfile",
		"p",
		nil,
		"添加COPY，格式: --copyfile KEY=VALUE",
	)

	//如果你希望这个 flag 是必填的，取消下面这行注释
	//dockerfileCmd.MarkFlagRequired("config")

	// 命令注册
	rootCmd.AddCommand(dockerfileCmd)
}
