package utility

import (
	"github.com/hegedustibor/htgo-tts"
	"log"
	"math/rand"
	cleaner "skybot/cleaner"
	"time"
)

func TriggerSpeachClean(text string) {

	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak(text)

	err := cleaner.RemoveAudioFIles("audio/*")
	if err != nil {
		log.Fatalf("Error removing files: %+v", err)
	}
}

func RandomQueryText() string {
	combos := []string{
		"What can i do for you.",
		"How may i help you.",
	}
	combos = append(combos, getQueryStringWithTime())

	rand.Seed(time.Now().UnixNano())
	return combos[rand.Intn(len(combos))]
}

func getQueryStringWithTime() string {
	now := time.Now().Hour()

	text := "Good "
	if now >= 0 && now < 13 {
		text += "Morning "
	} else if now < 18 {
		text += "Afternoon "
	} else {
		text += "Evening "
	}
	text += ", How may i be of help."

	return text
}
