package flurry

import (
	"net/url"
	"testing"
	"time"
)

var (
	//Replace with valid value before test
	flurryAccessCodeForTest = "xxxxxxxxxxxxxxxxxxxx"
	flurryApiKeyForTest     = "yyyyyyyyyyyyyyyyyyyy"
)

func TestAppInfo(t *testing.T) {
	cli := New(flurryAccessCodeForTest)
	info, err := cli.AllApplications()
	if err != nil {
		t.Error(err)
		return
	}
	var appApiKey string

	t.Log("All Applications of", info.CompanyName, "(UpdateAt:", info.GeneratedDate, ")")
	for _, app := range info.Applications {
		appApiKey = app.ApiKey
		t.Logf("    %s %s %8s %s", app.CreatedDate, app.ApiKey, app.Platform, app.Name)
	}

	if appApiKey != "" {
		info, err := cli.ApplicationInfo(appApiKey)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("")
		t.Log("Version list of", info.Name, " on ", info.Platform, "(CreatedAt:", info.CreatedDate, ")")
		for _, version := range info.Versions {
			t.Logf("    %s %s", version.CreatedDate, version.Name)
		}
	}
}

func TestAppMetrics(t *testing.T) {
	client := New(flurryAccessCodeForTest)
	q := url.Values{}
	q.Set("apiKey", flurryApiKeyForTest)
	q.Set("startDate", "2016-05-01")
	q.Set("endDate", "2016-05-10")

	info, err := client.AppMetrics("NewUsers", q)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("GeneratedAt:", info.GeneratedDate)
	t.Log("Duration:", info.StartDate, "=>", info.EndDate)
	for _, m := range info.Day {
		t.Logf("    %s: %s\n", m.Date, m.Value)
	}
}

func TestRetainedUsersMetrics(t *testing.T) {
	client := New(flurryAccessCodeForTest)
	apiKey := flurryApiKeyForTest
	start := time.Date(2016, 5, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2016, 5, 10, 0, 0, 0, 0, time.UTC)

	info, err := client.RetainedUsersMetrics(apiKey, start, end)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("GeneratedAt:", info.GeneratedDate)
	t.Log("Duration:", info.StartDate, "=>", info.EndDate)
	for _, m := range info.Day {
		t.Logf("    %s: %s\n", m.Date, m.Value)
	}
}
