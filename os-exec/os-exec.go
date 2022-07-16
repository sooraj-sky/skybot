package skybotexec

import (
	"log"
	"os/exec"
	util "skybot/utility"
)

const ShellToUse = "bash"

func CommandExecutor(sericeRequest string) {
	cmd := exec.Command(sericeRequest)
	util.TriggerSpeachClean("opening" + sericeRequest)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
