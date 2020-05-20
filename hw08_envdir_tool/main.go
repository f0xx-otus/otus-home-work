package main

import (
	"log"
	"os"
)

func main() {
	dir := os.Args[1]
	envDir, err := ReadDir(dir)
	if err != nil {
		log.Fatal("can't get environment variabes ", err)
	}
	err = changeOsEnv(envDir)
	if err != nil {
		log.Fatal("can't update OS variables", err)
	}
	returnCode := RunCmd(os.Args[2:], envDir)
	os.Exit(returnCode)
}
func changeOsEnv(envList Environment) error {
	for k, v := range envList {
		if v == "" {
			err := os.Unsetenv(k)
			if err != nil {
				return err
			}
			delete(envList, k)
			break
		}
		_, ok := os.LookupEnv(k)
		if ok {
			err := os.Unsetenv(k)
			if err != nil {
				return err
			}
			err = os.Setenv(k, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
