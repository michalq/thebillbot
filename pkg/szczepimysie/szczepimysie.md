# Szczepimysie

Ministerstwo Zdrowia in Poland doesn't provide an official API for vaccination status. But it uses arcgis.com to present this data, and we can query this data easily using arcgis.

https://services-eu1.arcgis.com/zk7YlClTgerl62BY/ArcGIS/rest/services/global_szczepienia_widok2/FeatureServer/0

Making query
https://services-eu1.arcgis.com/zk7YlClTgerl62BY/ArcGIS/rest/services/global_szczepienia_widok2/FeatureServer/0/query

Example query:
https://services-eu1.arcgis.com/zk7YlClTgerl62BY/ArcGIS/rest/services/global_szczepienia_widok2/FeatureServer/0/query?where=1%3D1&objectIds=&time=&resultType=none&outFields=*&returnIdsOnly=false&returnUniqueIdsOnly=false&returnCountOnly=false&returnDistinctValues=false&cacheHint=false&orderByFields=&groupByFieldsForStatistics=&outStatistics=&having=&resultOffset=&resultRecordCount=&sqlFormat=none&f=pjson&token=