package szczepimysie

import (
	"context"
	"fmt"
	"time"

	"github.com/michalq/thebillbot/pkg/arcgis"
)

type Client struct {
	arcGisClient *arcgis.Client
	serviceName  string
}

func NewClient(arcGisClient *arcgis.Client, serviceName string) *Client {
	return &Client{arcGisClient, serviceName}
}

func (c *Client) LatestStatus(ctx context.Context) (*Data, error) {

	statuses, err := c.DailyStatuses(ctx)
	if err != nil {
		return nil, err
	}
	if len(statuses) <= 0 {
		return nil, nil
	}
	latest := statuses[0]
	for _, status := range statuses {
		if status.Data.After(latest.Data) {
			latest = status
		}
	}
	return latest, nil
}

func (c *Client) DailyStatuses(ctx context.Context) ([]*Data, error) {
	resp, err := c.arcGisClient.Query(ctx, c.serviceName, arcgis.QueryParams{
		Where:                "1=1",
		F:                    "pjson",
		ResultType:           "none",
		OutFields:            "*",
		ReturnIdsOnly:        "false",
		ReturnUniqueIdsOnly:  "false",
		ReturnCountOnly:      "false",
		ReturnDistinctValues: "false",
		CacheHint:            "false",
		SqlFormat:            "none",
	})
	if err != nil {
		return nil, err
	}
	dailyStatuses := make([]*Data, 0, len(resp.Features))
	for _, feat := range resp.Features {
		dailyStatuses = append(dailyStatuses, formatData(feat.Attributes))
	}
	return dailyStatuses, nil
}

func formatData(rawData map[string]interface{}) *Data {
	return &Data{
		ObjectId:            retrieveIntValue(rawData, "OBJECTID", 0),
		Data:                retrieveDateValue(rawData, "DATA", time.Now()),
		SzczepieniaSuma:     retrieveIntValue(rawData, "SZCZEPIENIA_SUMA", 0),
		SzczepieniaDziennie: retrieveIntValue(rawData, "SZCZEPIENIA_DZIENNIE", 0),
		Dawka1Suma:          retrieveIntValue(rawData, "DAWKA_1_SUMA", 0),
		Dawka2Suma:          retrieveIntValue(rawData, "DAWKA_2_SUMA", 0),
		Dawka2Dziennie:      retrieveIntValue(rawData, "DAWKA_2_DZIENNIE", 0),
	}
}

func retrieveIntValue(rawData map[string]interface{}, key string, def int) int {

	if val, ok := rawData[key]; ok {
		if intVal, ok := val.(float64); ok {
			return int(intVal)
		}
	}
	return def
}

func retrieveDateValue(rawData map[string]interface{}, key string, def time.Time) time.Time {

	if val, ok := rawData[key]; ok {
		fmt.Println(key, val)
		if intVal, ok := val.(int64); ok {
			return time.Unix(int64(intVal), 0)
		}
	}
	return def
}
