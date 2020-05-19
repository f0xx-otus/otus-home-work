package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env
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
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = command.Wait()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
		log.Fatal(err)
	}
	return 0
}
