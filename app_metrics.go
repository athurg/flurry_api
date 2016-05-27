package flurry

import (
	"encoding/xml"
	"net/url"
	"time"
)

type AppMetricsInfo struct {
	XMLName       xml.Name `xml:"appMetrics"`
	Metric        string   `xml:"metric,attr"`
	Version       string   `xml:"version,attr"`
	GeneratedDate string   `xml:"generatedDate,attr"`
	StartDate     string   `xml:"startDate,attr"`
	EndDate       string   `xml:"endDate,attr"`
	Day           []struct {
		Value string `xml:"value,attr"`
		Date  string `xml:"date,attr"`
	} `xml:"day"`
}

/*
获取采集数据
    metric包括如下类型：
        ActiveUsers            活跃用户
        ActiveUsersByWeek      活跃用户（周汇总）
        ActiveUsersByMonth     活跃用户（月汇总）
        NewUsers               新用户
        MedianSessionLength    在线时长（中位数）
        AvgSessionLength       在线时长（平均值）
        Sessions               在线总时长
        RetainedUsers          留存用户
        PageViews              PageView
        AvgPageViewsPerSession 每会话平均PageView
*/
func (cli *Client) AppMetrics(metric string, q url.Values) (info AppMetricsInfo, err error) {
	err = cli.request("/appMetrics/"+metric, q, &info)
	return info, err
}

//活跃用户数
func (cli *Client) ActiveUsersMetrics(appApiKey string, start, end time.Time) (info AppMetricsInfo, err error) {
	q := url.Values{}
	q.Set("apiKey", appApiKey)
	q.Set("startDate", start.Format("2006-01-02"))
	q.Set("endDate", end.Format("2006-01-02"))

	return cli.AppMetrics("ActiveUsers", q)
}

//新增用户数
func (cli *Client) NewUsersMetrics(appApiKey string, start, end time.Time) (info AppMetricsInfo, err error) {
	q := url.Values{}
	q.Set("apiKey", appApiKey)
	q.Set("startDate", start.Format("2006-01-02"))
	q.Set("endDate", end.Format("2006-01-02"))

	return cli.AppMetrics("NewUsers", q)
}

//留存用户数
func (cli *Client) RetainedUsersMetrics(appApiKey string, start, end time.Time) (info AppMetricsInfo, err error) {
	q := url.Values{}
	q.Set("apiKey", appApiKey)
	q.Set("startDate", start.Format("2006-01-02"))
	q.Set("endDate", end.Format("2006-01-02"))

	return cli.AppMetrics("RetainedUsers", q)
}
