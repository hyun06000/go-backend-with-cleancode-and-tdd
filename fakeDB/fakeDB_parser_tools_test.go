package fakeDB

import (
	"testing"
)

func TestStartWith_SHOW(t *testing.T) {
	got := StartWith_SHOW("SHOW Query")
	assertBoolDifference(t, got, true)

	got = StartWith_SHOW("Query")
	assertBoolDifference(t, got, false)
}

func TestStartWith_CREATE(t *testing.T) {
	got := StartWith_CREATE("CREATE Query")
	assertBoolDifference(t, got, true)

	got = StartWith_CREATE("Query")
	assertBoolDifference(t, got, false)
}

func TestStartWith_USE(t *testing.T) {
	got := StartWith_USE("USE Query")
	assertBoolDifference(t, got, true)

	got = StartWith_USE("Query")
	assertBoolDifference(t, got, false)
}
func TestStartWith_INSERT_INTO(t *testing.T) {
	got := StartWith_INSERT_INTO("INSERT INTO Query")
	assertBoolDifference(t, got, true)

	got = StartWith_INSERT_INTO("Query")
	assertBoolDifference(t, got, false)
}

func TestStartWith_SELECT(t *testing.T) {
	got := StartWith_SELECT("SELECT Query")
	assertBoolDifference(t, got, true)

	got = StartWith_SELECT("Query")
	assertBoolDifference(t, got, false)
}

func TestStartWith_DATABASE(t *testing.T) {
	got := StartWith_DATABASE("DATABASE Query")
	assertBoolDifference(t, got, true)

	got = StartWith_DATABASE("Query")
	assertBoolDifference(t, got, false)
}

func TestStartWith_TABLE(t *testing.T) {
	got := StartWith_TABLE("TABLE Query")
	assertBoolDifference(t, got, true)

	got = StartWith_TABLE("Query")
	assertBoolDifference(t, got, false)
}

func TestStartWith_TABLES(t *testing.T) {
	got := StartWith_TABLES("TABLES Query")
	assertBoolDifference(t, got, true)

	got = StartWith_TABLES("Query")
	assertBoolDifference(t, got, false)
}

func TestStartWith_ASTERISK(t *testing.T) {
	got := StartWith_ASTERISK("* Query")
	assertBoolDifference(t, got, true)

	got = StartWith_ASTERISK("Query")
	assertBoolDifference(t, got, false)
}

func TestStartWith_ALHPABET(t *testing.T) {
	got := StartWith_ALPHABET("Query")
	assertBoolDifference(t, got, true)

	got = StartWith_ALPHABET("* Query")
	assertBoolDifference(t, got, false)
}

func TestSplit(t *testing.T) {
	got := Split("A - B - C", " - ")
	want := []string{"A", "B", "C"}

	for i, got_val := range got {
		assertStringDifference(t, "Split", got_val, want[i])
	}

}

func TestGetSplitedWord(t *testing.T) {
	got := GetSplitedWord("A B C D E", 2)
	want := "C"

	assertStringDifference(t, "SplitedWord", got, want)
}

func TestTrimPreSuf(t *testing.T) {
	got := TrimPreSuf("This is a fakeDB", "This i", "eDB")
	want := "s a fak"

	assertStringDifference(t, "TrimPreSuf", got, want)
}

func TestMergeWordList(t *testing.T) {
	A := []string{"A", "B", "C", "D"}
	B := []string{"a", "b", "c", "d"}
	got := MergeWordList(A, B, "_:_")
	want := []string{"A_:_a", "B_:_b", "C_:_c", "D_:_d"}

	for i, got_val := range got {
		assertStringDifference(t, "MergeWordList", got_val, want[i])
	}
}

func TestStringQuotesCheck(t *testing.T) {
	got := HasSingleQuotesPreSuf("'ABC'")
	assertBoolDifference(t, got, true)

	got = HasSingleQuotesPreSuf("ABC")
	assertBoolDifference(t, got, false)

}
