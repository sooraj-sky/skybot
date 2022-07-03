package main

import (
"github.com/hegedustibor/htgo-tts"
userinit "skybot/user"
)

func main() {
    currenWho := userinit.Whoami()
    userinit.UserInvok()
    speech := htgotts.Speech{Folder: "audio", Language: "en"}
    welcomeSpeach := "Hello" + currenWho + "Welcome to Skywalks" 
    speech.Speak(welcomeSpeach)
}
