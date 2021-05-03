package arcgis

type QueryParams struct {
	Where                      string `schema="where"`
	ObjectIds                  string `schema="objectIds"`
	Time                       string `schema="time"`
	ResultType                 string `schema="resultType"`
	OutFields                  string `schema="outFields"`
	ReturnIdsOnly              string `schema="returnIdsOnly"`
	ReturnUniqueIdsOnly        string `schema="returnUniqueIdsOnly"`
	ReturnCountOnly            string `schema="returnCountOnly"`
	ReturnDistinctValues       string `schema="returnDistinctValues"`
	CacheHint                  string `schema="cacheHint"`
	OrderByFields              string `schema="orderByFields"`
	GroupByFieldsForStatistics string `schema="groupByFieldsForStatistics"`
	OutStatistics              string `schema="outStatistics"`
	Having                     string `schema="having"`
	ResultOffset               string `schema="resultOffset"`
	ResultRecordCount          string `schema="resultRecordCount"`
	SqlFormat                  string `schema="sqlFormat"`
	F                          string `schema="f"`
	Token                      string `schema="token"`
}
