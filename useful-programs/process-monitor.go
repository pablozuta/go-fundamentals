package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func monitorProcess(pid int, results chan<- string) {  
	cmd := exec.Command("ps", "-o", "%cpu,%mem", strconv.Itoa(pid))  
	output, err := cmd.CombinedOutput()  
	  
	if err != nil {  
	results <- fmt.Sprintf("Error monitoring process %d: %v", pid, err)  
	return  
	}  
	  
	lines := strings.Split(string(output), "\n")  
	if len(lines) < 2 {  
	results <- fmt.Sprintf("Error monitoring process %d: no output", pid)  
	return  
	}  
	  
	parts := strings.Fields(lines[1])  
	if len(parts) < 2 {  
	results <- fmt.Sprintf("Error monitoring process %d: invalid output", pid)  
	return  
	}  
	  
	cpu, err := strconv.ParseFloat(parts[0], 64)  
	if err != nil {  
	results <- fmt.Sprintf("Error monitoring process %d: %v", pid, err)  
	return  
	}  
	  
	mem, err := strconv.ParseFloat(parts[1], 64)  
	if err != nil {  
	results <- fmt.Sprintf("Error monitoring process %d: %v", pid, err)  
	return  
	}  
	  
	results <- fmt.Sprintf("Process %d: CPU %.2f%%, memory %.2f%%", pid, cpu, mem)  
}  
  
func main() {  
pids := []int{1, 2, 3, 4, 5}  
results := make(chan string)  
  
for _, pid := range pids {  
go monitorProcess(pid, results)  
}  
  
for i := 0; i < len(pids); i++ {  
fmt.Println(<-results)  
}  
}