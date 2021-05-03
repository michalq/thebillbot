package arcgis

type EsriFieldType string

const (
	EsriFieldTypeInteger EsriFieldType = "esriFieldTypeInteger"
	EsriFieldTypeOID     EsriFieldType = "esriFieldTypeOID"
	EsriFieldTypeDate    EsriFieldType = "esriFieldTypeDate"
	EsriFieldTypeString  EsriFieldType = "esriFieldTypeString"
)

type ArcGisResponse struct {
	ObjectIdFieldName string `json:"objectIdFieldName"`
	UniqueIdField     struct {
		Name               string `json:"name"`
		IsSystemMaintained bool   `json:"isSystemMaintained"`
	} `json:"uniqueIdField"`
	GlobalIdFieldName string `json:"globalIdFieldName"`
	Fields            []struct {
		Name         string        `json:"name"`
		Type         EsriFieldType `json:"type"`
		Alias        string        `json:"alias"`
		SqlType      string        `json:"sqlType"`
		Domain       interface{}   `json:"domain"`
		DefaultValue interface{}   `json:"defaultValue"`
	} `json:"fields"`
	Features []struct {
		Attributes map[string]interface{}
	} `json:"features"`
}
