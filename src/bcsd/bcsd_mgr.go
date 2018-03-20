package main

import (
	"BCSD/node"
	"BCSD/block"
)

type BCSDMgr struct {
	node_mgr	NodeMgr
	block_mgr	BlockMgr
	service		Service
}

func (bcsd_mgr* BCSDMgr) Start() {

}