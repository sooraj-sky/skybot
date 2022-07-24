package wordprocessing

import (
	"fmt"
	"net/url"
	"os"
	"skybot/os-exec"
	util "skybot/utility"
	skyweather "skybot/weather-search"
	"skybot/wikki-search"
	"strings"
	hardcoded "skybot/hc-intelligence"
	"bufio"
	"flag"
	"log"
	"encoding/base64"
)


func base64Decode(str string) (string, bool) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", true
	}
	return string(data), false
}


func NaturalLanguageProcessor(inputword string) {

	fptr := flag.String("fpath", "bwords.txt", "file path to read from")
	flag.Parse()
  
	f, err := os.Open(*fptr)
	if err != nil {
	  log.Fatal(err)
	}
	defer func() {
	  if err = f.Close(); err != nil {
	  log.Fatal(err)
	}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
	  decoded_str, decode_err := base64Decode(s.Text())
	  if decode_err {
		  fmt.Println("Decoding failed.")
	  }
	  if (strings.Contains(inputword, decoded_str)) {
		util.TriggerSpeachClean("i won't respond to that")
		os.Exit(3)
		
	  }
	  
	}
	err = s.Err()
	if err != nil {
	  log.Fatal(err)
	}


	switch {
	case strings.HasPrefix(inputword, "open"):
		applicationName := strings.Fields(inputword)
		skybotexec.CommandExecutor(applicationName[1])

	case strings.HasPrefix(inputword, "temperature"):
		applicationName := strings.Fields(inputword)
		skyweather.WetherSearch(applicationName[2])

	case strings.HasPrefix(inputword, "search") || strings.HasPrefix(inputword, "who"):
		applicationName := strings.Fields(inputword)

		if applicationName[1] == "is" {
			w, err := wikimedia.New("http://en.wikipedia.org/w/api.php")
			if err != nil {
				fmt.Println(err)
			}

			findwho := (strings.TrimPrefix(inputword, "who is"))
			capsWho := (strings.Title(strings.ToLower(findwho)))
			f := url.Values{
				"action": {"query"},
				"prop":   {"extracts"},
				"titles": {capsWho},
			}
			res, err := w.Query(f)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res.Query.Pages)
			for _, v := range res.Query.Pages {
				result := v.Extract[:200]
				resultSecond := v.Extract[200:400]
				util.TriggerSpeachClean(result)
				util.TriggerSpeachClean(resultSecond)

			}

		} else if applicationName[1] == "are" && applicationName[2] == "you" {
			hardcoded.HcIntell("me")

		}else{

			w, err := wikimedia.New("http://en.wikipedia.org/w/api.php")
			if err != nil {
				fmt.Println(err)
			}
			f := url.Values{
				"action": {"query"},
				"prop":   {"extracts"},
				"titles": {applicationName[1]},
			}
			res, err := w.Query(f)
			if err != nil {
				fmt.Println(err)
			}
			for _, v := range res.Query.Pages {
				result := v.Extract[:200]
				resultSecond := v.Extract[200:400]
				util.TriggerSpeachClean(result)
				util.TriggerSpeachClean(resultSecond)

			}
		}

	case inputword == "exit":
		util.TriggerSpeachClean("Thank you")
		os.Exit(3)
	}

}
