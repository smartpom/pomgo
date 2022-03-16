package designpatterns

import (
	"fmt"
	"sync"
	"github.com/looplab/fsm"
)

var lock = &sync.Mutex{}

var singlePomFsmInstance *fsm.FSM

func GetPomFsmInstance() *fsm.FSM {
	if singlePomFsmInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singlePomFsmInstance == nil {
			fmt.Println("Creting Single PomFsm Instance Now")
			singlePomFsmInstance = fsm.NewFSM(
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
						fmt.Println("challenger: " + e.Src  + " -> " + e.FSM.Current())
					},
					"target": func(e *fsm.Event) {
						fmt.Println("target: " + e.Src  + " -> " + e.FSM.Current())
					},
					"challenge": func(e *fsm.Event) {
						fmt.Println("challenge: " + e.Src  + " -> " + e.FSM.Current())
					},
					"witness": func(e *fsm.Event) {
						fmt.Println("witness: " + e.Src  + " -> " + e.FSM.Current())
					},
					"finish": func(e *fsm.Event) {
						fmt.Println("finish: " + e.Src  + " -> " + e.FSM.Current())
					},
				},
			)
		} else {
			fmt.Println("Single PomFsm Instance already created-1")
		}
	} else {
		fmt.Println("Single PomFsm Instance already created-2")
	}
	return singlePomFsmInstance
}
