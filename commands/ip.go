package commands

import(
	"github.com/levigross/grequests"
	"log"
	"strings"
)


func GetIp() (int, error) {
	
	resp, err := grequests.Get("http://ipinfo.io/ip", nil)
	
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
		return -1, err
	}

	srcIp := strings.Replace(resp.String(), "\n", "", -1)
	log.Printf("curl ipinfo.io/ip \n")
	log.Printf("Outside IP: %s \n", srcIp)
	return 0, nil
}
