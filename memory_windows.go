package main

import (
	"os/exec"
	"io/ioutil"
	"path/filepath"
	log "github.com/Sirupsen/logrus"
)

var winpmemPath string = filepath.Join(binPath, "winpmem.exe")

func dropWinpmem() error {
	err := initBinFolder()
	if err != nil {
		return err
	}

	winpmemData, err := Asset("winpmem.exe")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(winpmemPath, winpmemData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func generateMemoryDump() {
	log.Info("Taking a snapshot of the system memory...")

	err := dropWinpmem()
	if err != nil {
		log.Error("Unable to find winpmem: ", err.Error())
		return
	}

	cmdArgs := []string{"--format", "raw", "--output", acq.Memory}

	err = exec.Command(winpmemPath, cmdArgs...).Run()
	if err != nil {
		log.Error("Unable to launch winpmem (are you running as Administrator?): ", err.Error())
		return
	}

	log.Info("Memory dump generated at ", acq.Memory)
}
