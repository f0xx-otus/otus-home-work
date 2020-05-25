package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	envSlice := make([]string, len(env))
	app := cmd[0]
	args := cmd[1:]
	command := exec.Command(app, args...)
	for k, v := range env {
		env := k + "=" + v
		envSlice = append(envSlice, env)
	}
	command.Env = append(os.Environ(), envSlice...)
	out, err := command.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
		log.Fatal(err)
	}
	fmt.Println(string(out))
	return 0
}
