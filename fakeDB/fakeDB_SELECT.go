package fakeDB

import (
	"sort"
	"strconv"
	"strings"
)

func SELECT(f *FakeDB, query string) DBMessage {
	query = strings.TrimPrefix(query, "SELECT ")

	switch {
	case StartWith_ASTERISK(query):
		return SELECT_ALL(f, query)
	case StartWith_ALPHABET(query):
		return SELECT_ELEMENT(f, query)
	default:
		return f.DBMsg
	}
}

type TableExtract struct {
	colName        string
	contentsString string
}

func SELECT_ALL(f *FakeDB, query string) DBMessage {
	query = strings.TrimPrefix(query, "* FROM ")

	tbName := TBName(query)
	table := f.MySQL[f.CurrentDB][tbName]

	terminal := []TableExtract{}
	for colName, contents := range table {
		cvtArgs := GenConvertArgs("SquareBraket_CommaSpace")
		cvtArgs.Sorted = false
		contentsString := ConvertListToString(*contents.contentList, cvtArgs)

		cvtArgs.Sorted = false
		rtn := TableExtract{
			colName:        string(colName),
			contentsString: contentsString,
		}
		terminal = append(terminal, rtn)
	}

	sort.Slice(terminal, func(i, j int) bool {
		return terminal[i].colName < terminal[j].colName
	})

	rtn := "|columns|contents| \n"
	cvtArgs := GenConvertArgs("BarBar_Bar")
	cvtArgs.Sorted = false
	for _, strs := range terminal {
		tmp := []string{strs.colName, strs.contentsString}
		rtn += ConvertListToString(tmp, cvtArgs) + " \n"
	}
	f.DBMsg.Terminal = rtn

	return f.DBMsg
}

func SELECT_ELEMENT(f *FakeDB, query string) DBMessage {
	tbName := TBName(GetSplitedWord(query, 2))
	table := f.MySQL[f.CurrentDB][tbName]

	colName := ColumnName(GetSplitedWord(query, 0))
	column := table[colName]
	contents := *column.contentList

	refColName := ColumnName(GetSplitedWord(query, 4))
	refColumn := table[refColName]

	refElement := GetSplitedWord(query, 6)
	var targetIndex int
	for i, element := range *refColumn.contentList {
		if element == refElement {
			targetIndex = i
		}
	}

	target := contents[targetIndex]

	rtn := "|columns|contents| \n"
	rtn += "|" + string(refColName) + "|[" + refElement + "]| \n"
	rtn += "|" + string(colName) + "|[" + target + "]| \n"

	f.DBMsg.Terminal = rtn
	f.DBMsg.SelectedValue, _ = strconv.Atoi(target)

	return f.DBMsg
}
