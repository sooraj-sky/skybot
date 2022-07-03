package userinit

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func Whoami() string {
	user, err := user.Current()
	username := user.Username
	if err != nil {
		log.Println("err")
	}

	return username
}

func UserInvok() {
	newpath := filepath.Join(".", "skybot-userdata")
	err := os.MkdirAll(newpath, os.ModePerm)

	if err != nil {
		log.Println("err")
	}

	user, err := user.Current()
	username := user.Username
	userDataFile := "skybot-userdata/" + username + ".json"
	if err != nil {
		log.Println("err")
	}

	f, err := os.OpenFile(userDataFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if f == nil {
		log.Fatal(err)
	}

}
