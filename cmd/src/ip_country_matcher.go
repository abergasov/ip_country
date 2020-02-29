package src

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var ips map[string]string

type ipResponse struct {
	IP                 string  `json:"ip"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryCode        string  `json:"country_code"`
	CountryCodeIso3    string  `json:"country_code_iso3"`
	CountryCapital     string  `json:"country_capital"`
	CountryTld         string  `json:"country_tld"`
	CountryName        string  `json:"country_name"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           string  `json:"latitude"`
	Longitude          string  `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  float64 `json:"country_population"`
	Message            string  `json:"message"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
}

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
	url := "https://ipapi.co/" + ip + "/json/"
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "Undefined"
	}

	var resp ipResponse
	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, &resp)
	return resp.CountryName
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
