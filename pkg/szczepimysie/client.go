package szczepimysie

type QueryParams struct {
	Where                      string `query="where"`
	ObjectIds                  string `query="objectIds"`
	Time                       string `query="time"`
	ResultType                 string `query="resultType"`
	OutFields                  string `query="outFields"`
	ReturnIdsOnly              string `query="returnIdsOnly"`
	ReturnUniqueIdsOnly        string `query="returnUniqueIdsOnly"`
	ReturnCountOnly            string `query="returnCountOnly"`
	ReturnDistinctValues       string `query="returnDistinctValues"`
	CacheHint                  string `query="cacheHint"`
	OrderByFields              string `query="orderByFields"`
	GroupByFieldsForStatistics string `query="groupByFieldsForStatistics"`
	OutStatistics              string `query="outStatistics"`
	Having                     string `query="having"`
	ResultOffset               string `query="resultOffset"`
	ResultRecordCount          string `query="resultRecordCount"`
	SqlFormat                  string `query="sqlFormat"`
	F                          string `query="f"`
	Token                      string `query="token"`
}
type Client struct {
	clientId string
}

func (c *Client) Query(service string) {

}
