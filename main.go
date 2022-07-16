package main

import (
	"bufio"
	"fmt"
	"os"
	userinit "skybot/user"
	util "skybot/utility"
	"skybot/wordprocessing"
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
			wordprocessing.NaturalLanguageProcessor(userCmdInput)			
		}
	}
}
