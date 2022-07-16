package main

import (
	userinit "skybot/user"
	util "skybot/utility"
	"fmt"
)

func main() {
	currenWho := userinit.Whoami()
	userinit.UserInvok()
	welcomeSpeach := "Hello" + currenWho + "Welcome to Skywalks"
	util.TriggerSpeachClean(welcomeSpeach)



	for {
		var input string

		util.TriggerSpeachClean(util.RandomQueryText())
		fmt.Println(" listening... ")
		fmt.Scanln(&input)

		if input == "exit" {
			util.TriggerSpeachClean("Thank you")
			break;
		} 
	}
}
