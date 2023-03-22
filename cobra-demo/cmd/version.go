package cmd

import "github.com/spf13/cobra"

/*
   @Auth: menah3m
   @Desc: 每个子命令都使用单独的文件
*/

func init() {
	// 配置命令的层级关系
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number of cobra-d",
	Run: func(cmd *cobra.Command, args []string) {
		println("cobra-d version: go 1.18")
	},
}
