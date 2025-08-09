package main

import (
	"fmt"
	"strings"
	"net/url"
	"net"
)

func main() {
	urlstr := []string {
		"httpminio.portworx.co:9090/path",
		"http://minio.portworx.co:80",
		//"http:/minio.portworx.co",
		"minio.portworx.co",
		"minio.portworx.co:9090",
		"s3.amazonaws.com",
		"http://s3.com/minio",
		"http://s3.com:80/minio",
		"htt://minio.portworx.co:80",
	}
	for _, val := range urlstr {
		var u *url.URL
		var err error
		var host, port string
		fmt.Printf(" \n -------------------------")
		fmt.Printf(" \n\n url: %v", val)
		if (strings.Contains(val, ":") && !strings.Contains(val, "://")) ||
			(!strings.Contains(val, ":") && !strings.Contains(val, "://")) {
			val = "http://" + val
			fmt.Printf("\n line 27 val: %v", val)
		}

		if u, err = url.ParseRequestURI(val); err != nil {
			fmt.Printf("\n invalid URL: %v", val)
			continue
		} else {
			fmt.Printf("\n no error in ParseRequestURI", err)
		}
		
		if host, port, err = net.SplitHostPort(u.Host); err == nil {
			fmt.Println("\n  Host:", host)
			fmt.Println("  Port:", port)
			fmt.Println("\n path:", u.Path)
			continue
		} 
		fmt.Printf("\n line 45 host: %v, port: %v", host, port)
		if port == "" {
			// take default port
			fmt.Printf(" default port")
			port = "80"
		}
		if host == "" {
			// For URL like "http://s3/minio", this cannot be parsed by net pkg
			// to extract host and port. So consume them directly with default port 80
			fmt.Printf(" default host")
			if u, err = url.Parse(val); err == nil {
				fmt.Printf(" final host: %v, path: %v", u.Host, u.Path)
				host = val
			} else {
				fmt.Printf(" error parsing url.Parse(val)")
			}
			
		}
	}
}