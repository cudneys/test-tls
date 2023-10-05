/*
Copyright © 2023 Scott Cudney <scott@cudneys.net>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/cudneys/test-tls/remote"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "test-tls",
	Short: "Tests TLS Certificates",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			panic(err)
		}

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			panic(err)
		}

		start := time.Now()
		err = remote.Test(host, port, "tcp")
		end := time.Now()
		delta := end.Sub(start)

		if err != nil {
			log.WithFields(log.Fields{"elapsed_time": delta, "error": err}).Error("Completed Test With Errors")
		} else {
			log.WithFields(log.Fields{"elapsed_time": delta}).Info("Completed Test Successfully")
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logLevel, _ := cmd.PersistentFlags().GetString("loglevel")

		switch logLevel {
		case "debug", "DEBUG":
			log.SetLevel(log.DebugLevel)
		case "warn", "warning", "WARN", "WARNING":
			log.SetLevel(log.WarnLevel)
		case "err", "error", "ERR", "ERROR":
			log.SetLevel(log.ErrorLevel)
		default:
			log.SetLevel(log.InfoLevel)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("loglevel", "l", "INFO", "Sets the log level")
	rootCmd.Flags().IntP("port", "p", 443, "The port to test")
	rootCmd.Flags().StringP("host", "H", "", "Remote host to test.")
	rootCmd.Flags().StringP("protocol", "P", "tcp", "The protocol to use for the test (tcp or udp)")
	rootCmd.MarkFlagRequired("host")
}
