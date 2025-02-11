package cmd

import (
	"cf-tool/client"
	"os"
)

// Pull command
func Pull() (err error) {
	cln := client.Instance
	info := Args.Info
	ac := Args.Accepted
	rootPath, err := os.Getwd()
	if err != nil {
		return
	}
	if err = cln.Pull(info, rootPath, ac); err != nil {
		if err = loginAgain(cln, err); err == nil {
			err = cln.Pull(info, rootPath, ac)
		}
	}
	return
}
