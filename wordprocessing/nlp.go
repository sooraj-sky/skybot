package wordprocessing

import (
	"strings"
	"skybot/os-exec"
	"os"
	util "skybot/utility"
)

func NaturalLanguageProcessor(inputword string) {
	if strings.HasPrefix(inputword, "open") {
		applicationName := strings.Fields(inputword)
		skybotexec.CommandExecutor(applicationName[1])
	}
	if inputword == "exit" {
		util.TriggerSpeachClean("Thank you")
		os.Exit(3)
	}
}