package awsinfo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Get() (map[string]interface{}, error) {
	info := make(map[string]interface{})
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	log.Println("Getting document")
	resp, err := client.Get("http://169.254.169.254/latest/dynamic/instance-identity/document")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("info %#v", string(data))

	err = json.Unmarshal(data, &info)
	if err != nil {
		return nil, err
	}

	log.Println("Getting hostname")
	resp, err = client.Get("http://169.254.169.254/latest/meta-data/public-hostname")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	log.Println("public-hostname", string(data))
	if err != nil {
		return nil, err
	}
	info["publicHostname"] = string(data)
	return info, nil
}
