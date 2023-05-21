package aosdk

import (
	"fmt"
	"strings"
)

func GetUrl(action string, queryMap map[string]string) string {
	query := ""
	for k, v := range queryMap {
		query += fmt.Sprintf("%s=%s&", k, v)
	}
	query = strings.TrimRight(query, "&")
	url := fmt.Sprintf("%s/%s/%s?%s", config.Endpoint, config.Version, action, query)

	return url
}
