package console

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
)

func Error(format string, args ...interface{}) {
	println(fmt.Sprintf(format, args...))
}

func Log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func InfiniteWait() {
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}

var Clear = map[string]func(){
	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	},
	"linux": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	},
}[runtime.GOOS]

var MoveToBegin = map[string]func(){
	"windows": func() {
		fmt.Print("\033[2J\033[1;1HW")
	},
	"linux": func() {
		fmt.Print("\033[2J\033[1;1H")
	},
}[runtime.GOOS]
