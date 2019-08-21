package main

import(
	"github.com/levigross/grequests"
	"log"
	"strings"
)

const dnsUpdateUrl string = "https://domains.google.com/nic/update"
const ipQueryUrl string = "http://ipinfo.io/ip"	
const domain string = "plex.hnxlabs.com"

func UpdateDns() (int, error) {
	cfg, err := GetConfig()
	if err != nil {
		log.Fatalln("error while reading config file")
		return -1, err
	}

	username, _ := cfg.String("plex.user")
	password, _ := cfg.String("plex.password")
	resp, err := grequests.Get(ipQueryUrl, nil)
	
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
		return -1, err
	}

	srcIp := strings.Replace(resp.String(), "\n", "", -1)
	log.Println(srcIp)
	resp, _ = grequests.Post(dnsUpdateUrl,
		&grequests.RequestOptions{ 
			Params: map[string]string{ "hostname": domain, "myip": srcIp},
			Auth: []string{username, password},
		})
	log.Printf("%s, %s\n", domain, resp.String());
	return 0, nil
}
