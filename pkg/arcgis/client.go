package arcgis

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

const BasePath = "https://services-eu1.arcgis.com/%s/ArcGIS/rest/services/%s/FeatureServer/0/query"

type Client struct {
	clientId string
}

func NewClient(clientId string) *Client {
	return &Client{clientId}
}

func (c *Client) Query(ctx context.Context, serviceName string, params QueryParams) (*ArcGisResponse, error) {
	apiUrl := fmt.Sprintf(BasePath, c.clientId, serviceName)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	query := url.Values{}
	if err := schema.NewEncoder().Encode(params, query); err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()
	req = req.WithContext(ctx)
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	responsePayload := &ArcGisResponse{}
	if err := json.Unmarshal(bodyBytes, responsePayload); err != nil {
		return nil, err
	}
	return responsePayload, nil
}
