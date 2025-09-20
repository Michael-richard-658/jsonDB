package operationdb

import (
	"fmt"
	"strings"
)

func (u *DBoperations) QueryParser(query string) ([]string, []string) {
	queryLength := len(query)
	semiColonIndex := queryLength - 1

	querySemicolonASCIIValue := query[semiColonIndex]
	if querySemicolonASCIIValue != ';' {
		fmt.Println("Expected ';' Found ' ' ")
		return []string{"; error"}, []string{}
	}

	queryParts := strings.Fields(
		strings.TrimSpace(
			strings.TrimSuffix(query, ";"),
		),
	)
	actionType := strings.ToLower(queryParts[0])
	switch actionType {
	case "select":

		parts := strings.Fields(query)

		if parts[1] == "*" {
			return parts, []string{}
		}

		attributes := []string{}
		i := 1
		for i < len(parts) && strings.ToUpper(parts[i]) != "FROM" {
			attributes = append(attributes, parts[i])
			i++
		}

		result := []string{}
		result = append(result, parts[0])
		result = append(result, strings.Join(attributes, ","))
		result = append(result, parts[i:]...)

		return result, []string{}
	case "desc":
		return queryParts, []string{}

	case "create":
		if strings.ToLower(queryParts[1]) != "table" {
			fmt.Println(queryParts[1] + " is not a valid command, did you mean TABLE?")
			return []string{"Error"}, nil
		}
		if len(queryParts) < 4 {
			fmt.Println(" TABLE doesnt have any fields(Not Allowed)!	")
			return []string{"Error"}, nil
		}

		attributes := queryParts[3:]
		return queryParts, attributes
	case "drop":
		return queryParts, []string{}
	default:
		return []string{"Error"}, nil

	}

}
