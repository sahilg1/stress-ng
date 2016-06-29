package main

import (
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
			print("CPU STRESS TEST")
		}
	}
}
