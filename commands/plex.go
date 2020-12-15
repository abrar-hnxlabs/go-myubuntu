package commands

import (
	"fmt"
	"log"
	"regexp"

	"github.com/levigross/grequests"
)

var apiResp map[string]interface{}

// Plexurl comment explaingin
func Plexurl() {
	downloadURL := "https://plex.tv/api/downloads/5.json"
	resp, err := grequests.Get(downloadURL, nil)
	if err != nil {
		log.Println("Unable to fetch plex api.")
		return
	}

	//releases := resp.JSON()["computer"]["Linus"]["releases"]
	resp.JSON(&apiResp)
	releases := apiResp["computer"].(map[string]interface{})["Linux"].(map[string]interface{})["releases"].([]interface{})
	for _, i := range releases {
		furl := i.(map[string]interface{})["url"].(string)
		matched, _ := regexp.MatchString(`^https:.*_amd64.deb$`, furl)
		if matched == true {
			fmt.Println(furl)
			return
		}
	}
}
