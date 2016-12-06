package gotunnelme

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	
	proxyUrl, err := url.Parse("http://www-proxy.ericsson.se:8080")
	
	if err != nil {
		return nil, err
	}
	
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	request, _ := http.NewRequest("GET", tunnelserverUrl, nil)
	response, httpErr := myClient.Do(request)
	
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
