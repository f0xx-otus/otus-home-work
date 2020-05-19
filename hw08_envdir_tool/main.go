package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("not enough aguments to run")
	}
	dir := os.Args[1]
	envDir, err := ReadDir(dir)
	if err != nil {
		log.Fatal("can't get environment variabes ", err)
	}
	fmt.Println(envDir)
	for k, v := range envDir {
		_, ok := os.LookupEnv(k)
		if ok {
			if v == "" {
				err = os.Unsetenv(k)
				if err != nil {
					fmt.Println("argument is ", k)
					log.Fatal("can't unset env Varable ", err)
				}
				delete(envDir, k)
				break
			}
			err = os.Unsetenv(k)
			if err != nil {
				fmt.Println("argument is ", k)
				log.Fatal("can't unset env Varable ", err)
			}
			err = os.Setenv(k, v)
			if err != nil {
				fmt.Println("argument is ", k)
				log.Fatal("can't set env variable ", err)
			}
		}
	}
	returnCode := RunCmd(os.Args[2:], envDir)
	os.Exit(returnCode)

}
