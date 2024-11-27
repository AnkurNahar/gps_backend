package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func FetchDevices(apiKey string) ([]map[string]interface{}, error) {
	url := "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to fetch data: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		ResultList []map[string]interface{} `json:"result_list"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.ResultList, nil
}