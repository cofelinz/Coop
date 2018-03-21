package master

import (
	"bcsd/node"
	"bcsd/block"
	"bcsd/server"
)

type BcsdMgr struct {
	node_mgr	node.NodeMgr
	block_mgr	block.BlockChainMgr
	service		server.Service
}

func (bcsd_mgr* BcsdMgr) Start() {
	bcsd_mgr.service.Start()
}