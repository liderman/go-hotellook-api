package hotellook

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
	"net/url"
)

type HotellookApi struct {
	token string
	log   LoggerInterface
}

type LoggerInterface interface {
	Debug(...interface{})
}

type VariationName struct {
	Name        string `json:"name" bson:"name"`
	IsVariation int    `json:"isVariation,string" bson:"isVariation"`
}

// NewHotellookApi creates a new instance HotellookApi.
func NewHotellookApi(token string) *HotellookApi {
	return &HotellookApi{
		token: token,
	}
}

func (a *HotellookApi) SetLogger(logger LoggerInterface) {
	a.log = logger
}

func (a *HotellookApi) getJson(path string, args map[string]string, v interface{}) error {
	apiUrl, err := url.Parse("http://engine.hotellook.com/api/v2/" + path)
	if err != nil {
		return err
	}
	params := url.Values{}
	for k, v := range args {
		if v == "" {
			continue
		}
		params.Add(k, v)
	}

	params.Add("token", a.token)
	apiUrl.RawQuery = params.Encode()

	if a.log != nil {
		a.log.Debug("API Send: " + apiUrl.String())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl.String(), nil)
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.Header.Get("Content-Encoding") == "gzip" {
		// decompress data
		reader, err := gzip.NewReader(res.Body)
		if err != nil {
			return err
		}
		defer reader.Close()
		return json.NewDecoder(reader).Decode(&v)
	}
	return json.NewDecoder(res.Body).Decode(&v)
}
