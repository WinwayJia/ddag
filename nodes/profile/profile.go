package main

import (
	"errors"
	"fmt"

	"github.com/WinwayJia/ddag/bus"
	"github.com/WinwayJia/ddag/nodes/decoder/msgtype"
)

var (
	NodeName = "profile"
)

type NodeStruct struct{}

var Node NodeStruct

type UserProfile struct {
	Age    int
	Gender int
}

func (n *NodeStruct) Process(seqNum string) error {
	fmt.Printf("enter node: %s\n", NodeName)

	msg, err := bus.GetMsg(seqNum, msgtype.NodeName, "key")
	if err != nil {
		fmt.Printf("node msg not found. seqNum: %s nodeName: %s key: %s\n", seqNum, msgtype.NodeName, "key")
		return err
	}

	user, ok := msg.(msgtype.UserInfo)
	if !ok {
		return errors.New("convert type failed")
	}

	fmt.Printf("get register value from bus: %+v\n", user)

	profile := UserProfile{
		Age:    1,
		Gender: 1,
	}

	bus.PutMsg(seqNum, NodeName, "key", profile)
	fmt.Printf("exit node: %s\n", NodeName)

	return nil
}
