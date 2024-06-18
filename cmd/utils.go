package cmd

import (
	"net"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gookit/color"
)

const (
	HOST = "localhost"
	PORT = "7836"
	TYPE = "tcp"
)

func CurrentDomain() string {
	dir, err := os.Getwd()

	if err != nil {
		color.Error.Println(err.Error())
		os.Exit(1)
	}

	dirSep := string(filepath.Separator)
	parts := slices.DeleteFunc(strings.Split(dir, dirSep), func(e string) bool {
		return e == ""
	})

	if !strings.HasPrefix(dir, dirSep+"sites") || len(parts) < 2 {
		color.Warn.Tips("This does not seem to be a SpinupWP site")
		os.Exit(1)
	}

	domain := parts[1]

	return domain
}

func PurgeCache(domain string) {
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
}
