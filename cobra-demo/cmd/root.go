package cmd

/*
   @Auth: menah3m
   @Desc: rootCmd 作为根命令，其他命令都是在rootCmd 的基础上添加得到的
*/

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use: "cobra-d",
	Run: func(cmd *cobra.Command, args []string) {
		println("this content is printed by cobra demo.")
	},
}

func Excute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
