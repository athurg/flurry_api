package flurry

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	apiAccessCode string
}

func New(accessCode string) *Client {
	return &Client{apiAccessCode: accessCode}
}

/*
NOTE:
    接口解析采用的是XML而不是JSON，原因是day节点返回数据格式没有一致性。
    当返回数据只有一天时，格式是：
		{"day":{"@date":"2016-04-01","@value":"52"}}
    当返回数据包含多天时，格式是：
		{"day":[{"@date":"2016-04-01","@value":"52"}]}
*/
func (cli *Client) request(path string, q url.Values, result interface{}) error {
	httpClient := &http.Client{}
	q.Set("apiAccessCode", cli.apiAccessCode)

	uri := "http://api.flurry.com" + path + "?" + q.Encode()
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Accept", "application/xml")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		return fmt.Errorf("错误码[%d]，请求过于频繁(间隔不得低于一秒)")
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("错误码[%d]", resp.StatusCode)
	}

	return xml.NewDecoder(resp.Body).Decode(&result)
}