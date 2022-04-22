package fakedb

import "strings"

func ConvertListToString(NameList []string) string {
	ListString := "["
	for i, Name := range NameList {
		ListString += Name
		if i != len(NameList)-1 {
			ListString += " "
		}
	}
	ListString += "]"

	return ListString
}

func IsShowDataBase(query string) bool {
	return (query == "show database")
}

func StartWith_CREATE(query string) bool {
	return strings.HasPrefix(query, "CREATE")
}

func StartWith_use(query string) bool {
	return strings.HasPrefix(query, "use")
}

func StartWith_INSERT_INTO(query string) bool {
	return strings.HasPrefix(query, "INSERT")
}

func StartWith_DATABASE(query string) bool {
	return strings.HasPrefix(query, "DATABASE")
}

func StartWith_TABLE(query string) bool {
	return strings.HasPrefix(query, "TABLE")
}

func SplitWithWhiteSpace(query string) []string {
	return strings.Split(query, " ")
}

func SplitWithCommaSpace(query string) []string {
	return strings.Split(query, ", ")
}

func GetSplitedWord(query string, index int) string {
	return SplitWithWhiteSpace(query)[index]
}

func TrimPreSuf(query string, prefix string, suffix string) string {
	query = strings.TrimPrefix(query, prefix)
	query = strings.TrimSuffix(query, suffix)

	return query
}

func MergeWordList(A []string, B []string, sep string) []string {
	rtn := []string{}
	for i := 0; i < len(A); i++ {
		rtn = append(rtn, A[i]+sep+B[i])
	}

	return rtn
}

func ConvertListToBarSplitedString(List []string) string {
	rtn := "|"
	for _, word := range List {
		rtn += word + "|"
	}
	return rtn
}

func GetTBNameAndColNameList(query string) (TBName, []ColumnName) {
	tbNameAndColNameList := strings.Split(query, "(")

	tbName := TBName(tbNameAndColNameList[0])

	colNames := strings.TrimSuffix(tbNameAndColNameList[1], ")")
	colNameList := SplitWithCommaSpace(colNames)

	var rtnColNameList []ColumnName
	for _, colName := range colNameList {
		rtnColNameList = append(rtnColNameList, ColumnName(colName))
	}

	return tbName, rtnColNameList

}
