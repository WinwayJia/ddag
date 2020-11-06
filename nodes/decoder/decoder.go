package main

import (
	"fmt"

	"github.com/WinwayJia/ddag/bus"
	"github.com/WinwayJia/ddag/nodes/decoder/msgtype"
)

type NodeStruct struct{}

var Node NodeStruct

func (n *NodeStruct) Process(seqNum string) error {
	fmt.Printf("enter node: %s\n", msgtype.NodeName)
	user := msgtype.UserInfo{
		OmgID: "omgid",
	}
	bus.PutMsg(seqNum, msgtype.NodeName, "key", user)
	fmt.Printf("exit node: %s\n", msgtype.NodeName)

	return nil
}
