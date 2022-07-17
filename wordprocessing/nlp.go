package wordprocessing

import (
	"strings"
	"skybot/os-exec"
	"skybot/wikki-search"
	"os"
	util "skybot/utility"
	"fmt"
	"net/url"
)

func NaturalLanguageProcessor(inputword string) {


	switch {
	case strings.HasPrefix(inputword, "open"):
		fmt.Println(strings.HasPrefix(inputword, "open"))
		applicationName := strings.Fields(inputword)
		skybotexec.CommandExecutor(applicationName[1])
	case strings.HasPrefix(inputword, "search"):
		applicationName := strings.Fields(inputword)
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
	  



			
	case inputword == "exit" :
			util.TriggerSpeachClean("Thank you")
		os.Exit(3)
	}

}