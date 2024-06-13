package cmd

import (
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

const (
	HOST = "localhost"
	PORT = "7836"
	TYPE = "tcp"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Purge a site's page cache",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		dir = "/sites/ashleyrich.com/"

		if err != nil {
			color.Error.Println(err.Error())
			os.Exit(1)
		}

		parts := strings.Split(dir, string(filepath.Separator))

		if !strings.HasPrefix(dir, "/sites") || len(parts) < 2 {
			color.Warn.Tips("This does not seem to be a SpinupWP site")
			os.Exit(1)
		}

		domain := parts[2]

		tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)

		if err != nil {
			color.Error.Println(err.Error())
			os.Exit(1)
		}

		conn, err := net.DialTCP(TYPE, nil, tcpServer)

		if err != nil {
			color.Error.Println(err.Error())
			os.Exit(1)
		}

		_, err = conn.Write([]byte("/cache/" + domain))

		if err != nil {
			color.Error.Println(err.Error())
			os.Exit(1)
		}

		color.Info.Tips("Cache purged for " + domain)

	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}
