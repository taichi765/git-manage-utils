package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf("failed to read directory: %v", err)
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		func() {
			os.Chdir(e.Name())
			defer os.Chdir("..")

			fmt.Printf("\033[32m%s\033[0m", e.Name())

			out, err := exec.Command("git", "status").Output()
			if err != nil {
				log.Fatalf("failed to execure 'git status': %v", err)
			}
			if strings.Contains(string(out), "nothing to commit, working tree clean") {
				fmt.Println("-- Clean")
			} else {
				fmt.Printf("\n%s\n", string(out))
			}
		}()
	}
}
