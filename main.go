package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	log "github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
	"github.com/shirou/gopsutil/mem"
)

var acq Acquisition

func main() {
	fmt.Println("                                 _ _                  ")
	fmt.Println("                                | (_)                 ")
	fmt.Println(" ___ _ __   ___   ___  _ __   __| |_  __ _  __ _      ")
	fmt.Println("/ __| '_ \\ / _ \\ / _ \\| '_ \\ / _` | |/ _` |/ _` | ")
	fmt.Println("\\__ \\ | | | (_) | (_) | |_) | (_| | | (_| | (_| |   ")
	fmt.Println("|___/_| |_|\\___/ \\___/| .__/ \\__,_|_|\\__, |\\__, |")
	fmt.Println("                      | |             __/ | __/ |     ")
	fmt.Println("                      |_|            |___/ |___/      ")
	fmt.Println("    (c) 2017-2018 Claudio Guarnieri (nex@nex.sx)      ")
	fmt.Println("                                                      ")

	// Set up the logging.
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())

	acq.Initialize()

	log.Info("Started acquisition ", acq.Folder)

	generateProfile()
	generateProcessList()
	generateAutoruns()

	virt, _ := mem.VirtualMemory()
	total := virt.Total / 1000000000

	log.Warning("Do you want to take a memory snapshot (it will take circa ", total, " GB of space) ? [y/N]")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	if choice == "y" || choice == "Y" {
		generateMemoryDump()
	} else {
		log.Info("Skipping memory acquisition.")
	}

	storeSecurely()

	log.Info("Acquisition completed.")

	log.Info("Press Enter to finish ...")
	var b = make([]byte, 1)
	os.Stdin.Read(b)
}
