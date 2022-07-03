package main

import (
"github.com/hegedustibor/htgo-tts"
userinit "skybot/user"
)

func main() {
    userinit.UserInvok()
    speech := htgotts.Speech{Folder: "audio", Language: "en"}
    speech.Speak("Welcome to Skywalks")
}
