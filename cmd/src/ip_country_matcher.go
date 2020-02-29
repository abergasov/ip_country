package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var ips map[string]string

func init() {
	ips = make(map[string]string)
}

func CheckIP(w http.ResponseWriter, req *http.Request) {
	ip := getIP(req)
	if val, ok := ips[ip]; ok {
		if val != "Undefined" {
			finishRequest(ip+"-"+val, w)
			return
		}
	}
	country := getCountry(ip)
	ips[ip] = country
	finishRequest(ip+"-"+country, w)
}

func getCountry(ip string) (country string) {
	url := "https://ipapi.co/" + ip + "/country/"
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return "Undefined"
	}

	fmt.Println(string(body))
	return string(body)
}

func finishRequest(row string, w http.ResponseWriter) {
	w.Write([]byte(row))
}

func getIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
