package main

import "plugin"

type NodeProcessor interface {
	Process(seqNum string) error
}

func main() {

	decoder, err := plugin.Open("./lib/decoder.so")
	if err != nil {
		panic(err)
	}

	profile, err := plugin.Open("./lib/profile.so")
	if err != nil {
		panic(err)
	}

	sym, err := decoder.Lookup("Node")
	if err != nil {
		panic(err)
	}

	process, ok := sym.(NodeProcessor)
	if !ok {
		panic("convert failed.")
	}
	process.Process("1")

	sym, err = profile.Lookup("Node")
	if err != nil {
		panic(err)
	}

	process, ok = sym.(NodeProcessor)
	if !ok {
		panic("convert failed.")
	}
	process.Process("1")
}
