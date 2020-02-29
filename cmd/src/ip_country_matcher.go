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
		if val != "UNDEFINED" {
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
		return "UNDEFINED"
	}

	fmt.Println(string(body))
	return string(body)
}

func finishRequest(row string, w http.ResponseWriter) {
	w.Write([]byte(row))
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
