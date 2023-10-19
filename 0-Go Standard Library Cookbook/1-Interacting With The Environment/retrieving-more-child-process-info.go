package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	var cmd string
	// runtime.GOOS nos devuelve el nombre de nuestro sistema operativo
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}

	proc := exec.Command(cmd, "1")
	proc.Start()

	// Wait function will wait till the process ends.
	proc.Wait()

	// After the process terminates the *os.ProcessState contains
	// simple information about the process run
	fmt.Printf("PID: %d\n", proc.ProcessState.Pid())
	fmt.Printf("Process took : %dms\n", proc.ProcessState.SystemTime()/time.Microsecond)
	fmt.Printf("Exited Sucessfully: %t\n", proc.ProcessState.Success())

}
// al correr el programa obtendremos 'false' en Exited Sucessfully , si usamos un valor literal para la variable cmd(asignarle directamente el valor "sleep" sin hacer el if statement) nos retornara 'true'