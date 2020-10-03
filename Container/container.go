package Container

import (
	"os"
	"os/exec"
	"time"
)

type Environment struct {
	Exec  string
	Path  string
	Port  string
	Index string
	Env   string
}

func killExec(process *os.Process) {
	time.Sleep(time.Second * 2)
	err2 := process.Kill()
	if err2 != nil {
		return
	}
}

func RunExec(srv Environment, feed *bool) {
	cmd := exec.Command(srv.Exec, append([]string{}, srv.Index)...)
	cmd.Dir = srv.Path
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	b := false

	err := cmd.Start()
	if err != nil {
		feed = &b
	}

	c := true
	feed = &c

	go killExec(cmd.Process)


}
