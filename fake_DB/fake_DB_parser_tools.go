package fakedb

import (
	"sort"
	"strings"
)

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

func Split(query string, spliter string) []string {
	return strings.Split(query, spliter)
}

func GetSplitedWord(query string, index int) string {
	return Split(query, " ")[index]
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

func GetTBNameAndColNameList(query string) (TBName, []ColumnName) {
	tbNameAndColNameList := strings.Split(query, "(")

	tbName := TBName(tbNameAndColNameList[0])

	colNames := strings.TrimSuffix(tbNameAndColNameList[1], ")")
	colNameList := Split(colNames, ", ")

	var rtnColNameList []ColumnName
	for _, colName := range colNameList {
		rtnColNameList = append(rtnColNameList, ColumnName(colName))
	}

	return tbName, rtnColNameList

}

type ConvertArgs struct {
	Pre    string
	Suf    string
	Sep    string
	Sorted bool
}

func ConvertListToString(
	wordList []string, convertArgs ConvertArgs) string {

	listString := convertArgs.Pre
	endIndex := len(wordList) - 1

	if convertArgs.Sorted {
		sort.Strings(wordList)
	}

	for i, Name := range wordList {
		listString += Name
		if i != endIndex {
			listString += convertArgs.Sep
		}
	}
	listString += convertArgs.Suf

	return listString
}

func ConvertStringToList(
	query string, convertArgs ConvertArgs) []string {

	queryWithoutPreSuf := TrimPreSuf(query, convertArgs.Pre, convertArgs.Suf)
	wordList := Split(queryWithoutPreSuf, convertArgs.Sep)

	if convertArgs.Sorted {
		sort.Strings(wordList)
	}

	return wordList
}

func GenConvertArgs(name string) ConvertArgs {
	cvtArgs := ConvertArgs{}
	nameList := Split(name, "_")
	switch preSuf := nameList[0]; preSuf {
	case "SquareBraket":
		cvtArgs.Pre = "["
		cvtArgs.Suf = "]"
	case "RoundBraket":
		cvtArgs.Pre = "("
		cvtArgs.Suf = ")"
	case "BarBar":
		cvtArgs.Pre = "|"
		cvtArgs.Suf = "|"
	default:
		cvtArgs.Pre = "!wrongBraket:" + preSuf + "!"
		cvtArgs.Suf = "!wrongBraket:" + preSuf + "!"
	}

	switch sep := nameList[1]; sep {
	case "WhiteSpace":
		cvtArgs.Sep = " "
	case "CommaSpace":
		cvtArgs.Sep = ", "
	case "Bar":
		cvtArgs.Sep = "|"
	default:
		cvtArgs.Sep = "!wrongSep:" + sep + "!"
	}
	cvtArgs.Sorted = true

	return cvtArgs
}
