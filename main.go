package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	//type:=os.Getenv("STRESS_TEST_TYPE")
	stressType := os.Getenv("STRESS_TEST_TYPE")
	os.Setenv("STRESS_WORKERS", "1")
	Workers := os.Getenv("STRESS_WORKERS")
	app := "stress-ng"
	//workers := 1
	switch stressType {
	case "CPU":
		{
			cpuPercentage := os.Getenv("STRESS_CPU_PERCENTAGE")
			arg1 := "-v"
			arg2 := "--cpu"
			arg3 := "--cpu-load"
			cmd := exec.Command(app, arg1, arg2, Workers, arg3, cpuPercentage)
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("CPU STRESS TEST")
		}
	case "MEM":
		{
			fmt.Println("MEMORY STRESS TEST")
			memBytes := os.Getenv("STRESS_MEM_BYTES")
			arg1 := "-v"
			arg2 := "--vm"
			arg3 := "--vm-bytes"
			cmd := exec.Command(app, arg1, arg2, Workers, arg3, memBytes)
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}

		}
	case "IO":
		{
			fmt.Println("IO STRESS TEST")
			arg1 := "-v"
			arg2 := "--io"
			cmd := exec.Command(app, arg1, arg2, Workers)
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}

		}

	}
}
