package main

import (
  "os"
	"os/exec"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)
func Keys() {
  keyboard.Listen(func(key keys.Key) (stop bool, err error) {
    if key.Code == keys.RuneKey {
      switch key.String() {
      case "q" :
        cmd := exec.Command("clear") 
        cmd.Stdout = os.Stdout
        cmd.Run()
        return true, nil 
      }
    } else {
        switch key.Code {
        case keys.CtrlC, keys.Escape:
          Esc()
      }
    }
    return false, nil 
  })
}
