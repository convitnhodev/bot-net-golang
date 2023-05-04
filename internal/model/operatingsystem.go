package model

type User struct {
	Name            string `json:"name"`
	OSName          string `json:"osName"`
	OSCodename      string `json:"osCodename"`
	OSCountry       string `json:"osCountry"`
	PublicIPAddress string `json:"publicIpAddress"`
	Region          string `json:"region"`
	RegionName      string `json:"regionName"`
	City            string `json:"city"`
	Country         string `json:"country"`
	CountryCode     string `json:"countryCode"`
	TimeZone        string `json:"timezone"`
	ISP             string `json:"isp"`
}
