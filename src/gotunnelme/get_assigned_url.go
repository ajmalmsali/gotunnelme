package gotunnelme

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	localtunnelServer = "http://ajm.al/"
)

type AssignedUrlInfo struct {
	Id           string `json:"id,omitempty"`
	Url          string `json:"url,omitempty"`
	Port         int    `json:"port,omitempty"`
	MaxConnCount int    `json:"max_conn_count,omitempty"`
}

func GetAssignedUrl(assignedDomain string) (*AssignedUrlInfo, error) {
	if len(assignedDomain) == 0 {
		assignedDomain = "?new"
	}
	tunnelserverUrl := fmt.Sprintf(localtunnelServer+"%s", assignedDomain)
	fmt.Println(tunnelserverUrl)

	proxy := os.Getenv("HTTP_PROXY")
	if proxy == "" {
		proxy = os.Getenv("http_proxy")
	}

	httpClient := http.DefaultClient

	if len(proxy) > 0 {
		proxyUrl, err := url.Parse(proxy)

		if err != nil {
			return nil, err
		}else{
			if Debug {
				fmt.Printf("ProxyUrl Parsing Error!")
			}
		}
		httpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}

	request, _ := http.NewRequest("GET", tunnelserverUrl, nil)

	response, httpErr := http.DefaultClient.Do(request)

	if len(proxy) > 0 {
		response, httpErr = httpClient.Do(request)
	}

	if httpErr != nil {
		return nil, httpErr
	}
	defer response.Body.Close()
	bytes, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}
	if Debug {
		fmt.Printf("***GetAssignedUrl: %s\n", string(bytes))
	}

	assignedUrlInfo := &AssignedUrlInfo{}
	if unmarshalErr := json.Unmarshal(bytes, assignedUrlInfo); unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return assignedUrlInfo, nil
}
