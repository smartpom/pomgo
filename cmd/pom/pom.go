//go:build linux || darwin || windows
// +build linux darwin windows

package main

import (
	"fmt"
	"github.com/looplab/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"idle",
		fsm.Events{
			{Name: "challenger", Src: []string{"idle"}, Dst: "challenger"},
			{Name: "target", Src: []string{"challenger"}, Dst: "target"},
			{Name: "challenge", Src: []string{"target"}, Dst: "challenge"},
			{Name: "witness", Src: []string{"challenge"}, Dst: "witness"},
			{Name: "finish", Src: []string{"witness"}, Dst: "idle"},
		},
		fsm.Callbacks{
			"challenger": func(e *fsm.Event) {
				fmt.Println("challenger: " + e.FSM.Current())
			},
			"target": func(e *fsm.Event) {
				fmt.Println("target: " + e.FSM.Current())
			},
			"challenge": func(e *fsm.Event) {
				fmt.Println("challenge: " + e.FSM.Current())
			},
			"witness": func(e *fsm.Event) {
				fmt.Println("witness: " + e.FSM.Current())
			},
			"finish": func(e *fsm.Event) {
				fmt.Println("finish: " + e.FSM.Current())
			},
		},
	)

	fmt.Println(fsm.Current())

	err := fsm.Event("challenger")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("1:" + fsm.Current())

	err = fsm.Event("target")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("2:" + fsm.Current())

	err = fsm.Event("challenge")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("3:" + fsm.Current())

	err = fsm.Event("witness")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("4:" + fsm.Current())

	err = fsm.Event("finish")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("5:" + fsm.Current())

}
