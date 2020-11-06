package bus

import (
	"bytes"
	"errors"
	"sync"
)

var (
	msgBus sync.Map
)

func formatKey(seqNum, nodeName, key string) string {
	var buff bytes.Buffer

	buff.WriteString(seqNum)
	buff.WriteString(nodeName)
	buff.WriteString(key)

	return buff.String()
}

func PutMsg(seqNum, nodeName, key string, value interface{}) {
	msgKey := formatKey(seqNum, nodeName, key)
	msgBus.Store(msgKey, value)
}

func GetMsg(seqNum, nodeName, key string) (value interface{}, err error) {
	msgKey := formatKey(seqNum, nodeName, key)
	value, ok := msgBus.Load(msgKey)
	if !ok {
		return nil, errors.New("key not found")
	}

	return
}
