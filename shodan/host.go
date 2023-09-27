package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HostLocation struct {
	City         string  `json:"city"`
	RegionCode   string  `json:"region_code"`
	AreaCode     int     `json:"area_code"`
	Longitude    float64 `json:"longitude"`
	CountryCode3 string  `json:"country_code3"`
	Latitude     float64 `json:"latitude"`
	PostalCode   string  `json:"postal_code"`
	DmaCode      int     `json:"dma_code"`
	CountryCode  string  `json:"country_code"`
	CountryName  string  `json:"country_name"`
}

type HostHTTPInfo struct {
	RobotsHash      interface{}            `json:"robots_hash"`
	Redirects       []interface{}          `json:"redirects"`
	Securitytxt     interface{}            `json:"securitytxt"`
	Title           string                 `json:"title"`
	SitemapHash     interface{}            `json:"sitemap_hash"`
	Robots          interface{}            `json:"robots"`
	Server          string                 `json:"server"`
	Host            string                 `json:"host"`
	HTML            string                 `json:"html"`
	Location        string                 `json:"location"`
	Components      map[string]interface{} `json:"components"`
	SecuritytxtHash interface{}            `json:"securitytxt_hash"`
	Sitemap         interface{}            `json:"sitemap"`
	HTMLHash        int                    `json:"html_hash"`
}

type HostShodanInfo struct {
	Crawler string   `json:"crawler"`
	PTR     bool     `json:"ptr"`
	ID      string   `json:"id"`
	Module  string   `json:"module"`
	Options struct{} `json:"options"`
}

type Host struct {
	Product   string         `json:"product"`
	Hash      int            `json:"hash"`
	IP        int            `json:"ip"`
	Org       string         `json:"org"`
	ISP       string         `json:"isp"`
	Transport string         `json:"transport"`
	CPE       []string       `json:"cpe"`
	Data      string         `json:"data"`
	ASN       string         `json:"asn"`
	Port      int            `json:"port"`
	Hostnames []string       `json:"hostnames"`
	Location  HostLocation   `json:"location"`
	Timestamp string         `json:"timestamp"`
	Domains   []string       `json:"domains"`
	HTTP      HostHTTPInfo   `json:"http"`
	OS        interface{}    `json:"os"`
	Shodan    HostShodanInfo `json:"_shodan"`
	IPStr     string         `json:"ip_str"`
}

type HostSearch struct {
	Matches []Host `json:"matches"`
}

func (s *Client) HostSearch(q string) (*HostSearch, error) {
	res, err := http.Get(
		fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseURL, s.apiKey, q),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret HostSearch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
