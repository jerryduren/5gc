package main

import (
	"fmt"
	
	"github.com/looplab/fsm"
)

func main() {
	locfsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed","opened"}, Dst: "opened"},
			{Name: "close", Src: []string{"opened","closed"}, Dst: "closed"},
		},
		
		fsm.Callbacks{
			"before_open": func(e *fsm.Event) {
				fmt.Println("before_open: " + e.FSM.Current())
			},
			"before_close": func(e *fsm.Event) {
				fmt.Println("before_close: " + e.FSM.Current())
			},
			"leave_open": func(e *fsm.Event) {
				fmt.Println("leave_open: " + e.FSM.Current())
			},
			"leave_close": func(e *fsm.Event) {
				fmt.Println("leave_close: " + e.FSM.Current())
			},
			"enter_open": func(e *fsm.Event) {
				fmt.Println("enter_open: " + e.FSM.Current())
			},
			"enter_close": func(e *fsm.Event) {
				fmt.Println("enter_close: " + e.FSM.Current())
			},
			"after_open": func(e *fsm.Event) {
				fmt.Println("after_open: " + e.FSM.Current())
			},
			"after_close": func(e *fsm.Event) {
				fmt.Println("after_close: " + e.FSM.Current())
			},
		},
	)
	
	locfsm.Event("open","para0","para1")
	locfsm.Event("close","para0","para1")
	
	locfsm.SetState("opened")
	locfsm.Event("open")
	
	fmt.Println(fsm.Visualize(locfsm))
	
}
