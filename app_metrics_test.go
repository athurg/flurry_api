package flurry

import (
	"net/url"
	"testing"
	"time"
)

func TestAppMetrics(t *testing.T) {
	//测试前输入有效的账户AccessCode及应用的ApiKey
	client := New("xxxxxxxxxxxxxxxxxxxx")
	q := url.Values{}
	q.Set("apiKey", "xxxxxxxxxxxxxxxxxxxx")
	q.Set("startDate", "2016-05-01")
	q.Set("endDate", "2016-05-10")

	//接口调用间隔不低于1秒
	time.Sleep(time.Second)

	info, err := client.AppMetrics("NewUsers", q)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("报告日期:", info.GeneratedDate)
	t.Log("数据范围:", info.StartDate, "=>", info.EndDate)
	for _, m := range info.Day {
		t.Logf("    %s: %s\n", m.Date, m.Value)
	}
}

func TestRetainedUsersMetrics(t *testing.T) {
	//测试前输入有效的账户AccessCode及应用的ApiKey
	client := New("xxxxxxxxxxxxxxxxxxxx")
	apiKey := "xxxxxxxxxxxxxxxxxxxx"
	start := time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2016, 5, 10, 0, 0, 0, 0, time.UTC)

	//接口调用间隔不低于1秒
	time.Sleep(time.Second)

	info, err := client.RetainedUsersMetrics(apiKey, start, end)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("报告日期:", info.GeneratedDate)
	t.Log("数据范围:", info.StartDate, "=>", info.EndDate)
	for _, m := range info.Day {
		t.Logf("    %s: %s\n", m.Date, m.Value)
	}
}
