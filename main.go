package main

import (
	"bufio"
	"fmt"
	"os"
	"skybot/os-exec"
	userinit "skybot/user"
	util "skybot/utility"
	"strings"
)

func main() {
	currenWho := userinit.Whoami()
	userinit.UserInvok()
	welcomeSpeach := "Hello" + currenWho + "Welcome to Skywalks"
	util.TriggerSpeachClean(welcomeSpeach)

	for {
		util.TriggerSpeachClean(util.RandomQueryText())
		fmt.Println(" listening... ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			userCmdInput := scanner.Text()
			if strings.HasPrefix(userCmdInput, "open") {
				applicationName := strings.Fields(userCmdInput)
				skybotexec.CommandExecutor(applicationName[1])
			}
			if userCmdInput == "exit" {
				util.TriggerSpeachClean("Thank you")
				break
			}
		}

	}
}
