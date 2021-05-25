package main

import (
	"fmt"
	. "server/state"
)

func main() {
	lightSwitchFsm := NewLightSwitchFSM()

	err := lightSwitchFsm.SendEvent(SwitchOff, nil)
	if err != nil {
		fmt.Println("Couldn't set the initial state of the state machine, err: %v", err)
	}

}
