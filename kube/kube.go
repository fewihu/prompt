package kube

import (
	"os"
	"os/exec"
	"prompt/text"
)

const (
	kubectl          = "kubectl"
	config           = "config"
	view             = "view"
	kubeconfigSwitch = "--kubeconfig"
	retOk            = 0
)

func GetKubeState(state chan<- *text.FormattedText) {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		state <- nil
		return
	}

	valid := make(chan bool)
	go validateKubeconfig(kubeconfig, valid)
	var color func(string) text.ColoredText
	if !<-valid {
		color = text.Red
	} else {
		color = text.Violett
	}
	kubeState := text.JoinSpace(text.Whale(text.Violett), text.BoldColor(color(" "+kubeconfig)), text.Newline())
	state <- &kubeState
}

func validateKubeconfig(path string, valid chan<- bool) {
	cmd := exec.Command(kubectl, config, view, kubeconfigSwitch, path)
	if err := cmd.Start(); err != nil {
		valid <- false
		return
	}

	if err := cmd.Wait(); err != nil {
		valid <- false
		return
	}

	if cmd.ProcessState.ExitCode() != retOk {
		valid <- false
		return
	}

	valid <- true
}
