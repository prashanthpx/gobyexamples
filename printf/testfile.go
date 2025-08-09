package main

import (
	"fmt"
	"os/exec"
	_ "runtime"
	"strings"
)

func main() {
	out, err := exec.Command("curl", "-s", "https://mirror.openshift.com/pub/openshift-v4/clients/ocp-dev-preview/latest/sha256sum.txt").Output()
	//fmt.Println(out)
	output := string(out[:])
	s := strings.Split(output, "\n")
	for _, v := range s {
		fmt.Printf(" \n line 16 v = ", v)
		if strings.Contains(v, "openshift-install-linux") {
			fmt.Println(v)
		}
		
	}

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Println(" Error in pgm")
		fmt.Printf("%s", err)
	}
}
