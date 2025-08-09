package main

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

func main() {
	urls := "httpminio.portworx.co/path:909"
	strTemp := urls
	if !strings.Contains(urls, ":") && !strings.Contains(urls, "://") {
		strTemp = "http://" + urls
	} else if strings.Contains(urls, ":") {
		strTemp = "http://" + urls
	}

	fmt.Printf(" strTemp: %v", strTemp)
	u, err := url.ParseRequestURI(strTemp)
		if err != nil {
			panic(err)
		}
	fmt.Printf("line 17 host: %v", u.Host)
		
	if host, port, err := net.SplitHostPort(u.Host); err == nil {
		fmt.Println("line 28 Host:", host)
		fmt.Println("line 29 Port:", port)
	} else {
		fmt.Println("\n Host:", u.Host)
	}
}