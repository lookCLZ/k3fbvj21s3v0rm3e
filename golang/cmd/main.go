package main

import "os/exec"

func main() {
	ctx,cancelFunc := context.WithCancel(context.TODO())

	go func(){
		cmd:=exec.CommandContext(ctx,"/bin/bash","-c","sleep 2;echo hello")
		output,err:=cmd.CombinedOutput()
	}()
	time.Sleep(1*time.Second)
	cancelFunc()
}
