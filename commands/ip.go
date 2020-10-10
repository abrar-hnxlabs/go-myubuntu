package commands

import(
	"github.com/levigross/grequests"
	"log"
	"strings"
)

const ipQueryUrl string = "http://ipinfo.io/ip"	

func GetIp() (int, error) {
	
	resp, err := grequests.Get(ipQueryUrl, nil)
	
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
		return -1, err
	}

	srcIp := strings.Replace(resp.String(), "\n", "", -1)
	log.Printf("Outside IP: %s \n", srcIp)
	return 0, nil
}
