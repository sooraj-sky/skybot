package aboutme

import (
	util "skybot/utility"
)

func HcIntell(hcQuery string) {
	switch {
	case hcQuery == "me":
		util.TriggerSpeachClean("I'm skybot designed and developed by Skywalks")
	case hcQuery == "skywalker":
		util.TriggerSpeachClean("Each and every employee of skywalks is a skywalker. skywalks is a Limited liability company, however we want so see the growth of all good things")
	}

}
