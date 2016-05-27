package flurry

import (
	"encoding/xml"
	"net/url"
)

type Application struct {
	Name        string `xml:"name,attr"`
	Platform    string `xml:"platform,attr"`
	ApiKey      string `xml:"apiKey,attr,omitempty"`
	CreatedDate string `xml:"createdDate,attr"`
}

type ApplicationVersion struct {
	Name        string `xml:"name,attr"`
	CreatedDate string `xml:"createdDate,attr"`
}

type AllApplicationsResponse struct {
	XMLName       xml.Name      `xml:"applications"`
	CompanyName   string        `xml:"companyName,attr"`
	GeneratedDate string        `xml:"generatedDate,attr"`
	Applications  []Application `xml:"application"`
}

type ApplicationInfoResponse struct {
	XMLName       xml.Name             `xml:"appInfo"`
	GeneratedDate string               `xml:"generatedDate,attr"`
	Versions      []ApplicationVersion `xml:"version"`
	Application
}

//获取指定apiKey的应用的版本信息
func (cli *Client) ApplicationInfo(apiKey string) (info ApplicationInfoResponse, err error) {
	q := url.Values{}
	q.Set("apiKey", apiKey)
	err = cli.request("/appInfo/getApplication", q, &info)
	return info, err
}

//获取所有应用的列表
func (cli *Client) AllApplications() (info AllApplicationsResponse, err error) {
	err = cli.request("/appInfo/getAllApplications", url.Values{}, &info)
	return info, err
}
