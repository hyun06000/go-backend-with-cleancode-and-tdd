package main

import "testing"

func TestHello(t *testing.T) {
	runName := "say hello to people with name"
	t.Run(runName, testNameHello)

	runName = "say hello to world without name"
	t.Run(runName, checkEmptyName)
}

func testNameHello(t *testing.T) {
	nameList := []string{"Chris", "David"}
	checkAllNameList(t, nameList)
}

func checkAllNameList(t *testing.T, nameList []string) {
	for _, personName := range nameList {
		checkPersonName(t, personName)
	}
}

func checkPersonName(t *testing.T, personName string) {
	got := Hello(personName)
	want := "Hello, " + personName

	assertCorrectMessage(t, want, got)
}

func checkEmptyName(t *testing.T) {
	got := Hello("")
	want := "Hello, World"

	assertCorrectMessage(t, want, got)
}

func assertCorrectMessage(t *testing.T, want string, got string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
