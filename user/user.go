package userinit

import (
	"os"
	"os/user"
	"path/filepath"
	"log"
)

func UserInvok () {
	newpath := filepath.Join(".", "skybot-userdata")
    err := os.MkdirAll(newpath, os.ModePerm)

    if err != nil  {
        log.Println("err")
    }
  
    user, err := user.Current()
    username := user.Username
    userDataFile := "skybot-userdata/" + username + ".json"
    if err != nil  {
        log.Println("err")
    }

    f, err := os.OpenFile(userDataFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if f == nil {
    log.Fatal(err)
    }

}