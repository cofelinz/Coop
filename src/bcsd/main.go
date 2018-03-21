package main

import (
	"bcsd/master"
)

func main() {
	bcsd_mgr := new(master.BcsdMgr)
	bcsd_mgr.Start()
}