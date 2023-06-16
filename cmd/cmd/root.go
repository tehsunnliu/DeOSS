/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/CESSProject/DeOSS/configs"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   configs.Name,
	Short: configs.Description,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init
func init() {
	rootCmd.AddCommand(
		cmd_config(),
		cmd_version(),
		cmd_run(),
		cmd_stat(),
		cmd_exit(),
	)
	rootCmd.PersistentFlags().StringP("config", "c", "conf.yaml", "custom profile")
	rootCmd.PersistentFlags().StringP("rpc", "", "wss://testnet-rpc0.cess.cloud/ws/", "rpc endpoint")
	rootCmd.PersistentFlags().StringP("ws", "", "/", "workspace")
	rootCmd.PersistentFlags().IntP("http_port", "P", 8080, "service listening port")
	rootCmd.PersistentFlags().IntP("p2p_port", "p", 4001, "p2p port")
	rootCmd.PersistentFlags().StringP("boot", "", "_dnsaddr.bootstrap-kldr.cess.cloud", "bootstap nodes")
}

func cmd_version() *cobra.Command {
	cc := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(configs.Version)
			os.Exit(0)
		},
		DisableFlagsInUseLine: true,
	}
	return cc
}

func cmd_config() *cobra.Command {
	cc := &cobra.Command{
		Use:                   "config",
		Short:                 "Generate configuration file",
		Run:                   cmd_config_func,
		DisableFlagsInUseLine: true,
	}
	return cc
}

func cmd_run() *cobra.Command {
	cc := &cobra.Command{
		Use:                   "run",
		Short:                 "Running services",
		Run:                   cmd_run_func,
		DisableFlagsInUseLine: true,
	}
	return cc
}

func cmd_stat() *cobra.Command {
	cc := &cobra.Command{
		Use:                   "stat",
		Short:                 "Query deoss information",
		Run:                   cmd_stat_func,
		DisableFlagsInUseLine: true,
	}
	return cc
}

func cmd_exit() *cobra.Command {
	cc := &cobra.Command{
		Use:                   "exit",
		Short:                 "Unregister the deoss role",
		Run:                   cmd_exit_func,
		DisableFlagsInUseLine: true,
	}
	return cc
}
