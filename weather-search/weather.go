package skyweather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	util "skybot/utility"
)

var protocol = "https://"
var baseURL = protocol + "wttr.in"

func WetherSearch(city string) {
	r := getCart(city)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Print(err)
	}
	temp := (string(body))
	util.TriggerSpeachClean("temperature in " + temp)
	fmt.Println(string(body))
}

func getCart(city string) *http.Response {
	response, err := http.Get(os.ExpandEnv(baseURL + "/" + city + "?format=3"))
	if err != nil {
		fmt.Print(err)
	}
	return response
}
