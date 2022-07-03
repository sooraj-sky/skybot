package main

import (
	"github.com/hegedustibor/htgo-tts"
	"log"
	cleaner "skybot/cleaner"
	userinit "skybot/user"
)

func main() {
	currenWho := userinit.Whoami()
	userinit.UserInvok()
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	welcomeSpeach := "Hello" + currenWho + "Welcome to Skywalks"
	speech.Speak(welcomeSpeach)

	// Removing Audio Directory
	err := cleaner.RemoveAudioFIles("audio/*")
	if err != nil {
		log.Fatalf("Error removing files: %+v", err)
	}
}
