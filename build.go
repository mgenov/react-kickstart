package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {

	if !isBableInstalled() {
		log.Println("bible was not installed, you have to set-up environment first.")
		os.Exit(-1)
	}

	go babelWatch()

	go func() {
		log.Fatal(http.ListenAndServe(":3000", http.FileServer(http.Dir("frontend"))))
	}()

	select {}
}
func isBableInstalled() bool {
	_, err := exec.LookPath("babel")
	if err != nil {
		return false
	}
	return true
}

// babelWatch is executing bable for continuous compilation of frontend sources.
func babelWatch() {
	cmd := exec.Command("babel", "--presets", "react", "src", "--watch", "--out-dir", "build")
	cmd.Dir = "frontend"
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Command finished.")
}
