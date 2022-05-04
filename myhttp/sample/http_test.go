package myhttp_sample

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	runName := "return score with name"
	t.Run(runName, checkPlayerScoreList)
}

const prefixURLPlayer string = "/players/"

func checkPlayerScoreList(t *testing.T) {
	playerList, scoreList := getPlayerListScoreList()
	comparePlayerWithScore(t, playerList, scoreList)
}

func getPlayerListScoreList() ([]string, []string) {
	playerList := []string{"Pepper", "Martin", "Trever"}
	scoreList := []string{"20", "16", "46"}

	return playerList, scoreList
}

func comparePlayerWithScore(
	t *testing.T, playerList []string, scoreList []string) {

	for i := 0; i < len(playerList); i++ {
		responseRecorder, request := getDummyRequestAndResponse(playerList[i])
		checkPlayerScore(t, request, responseRecorder, scoreList[i])
	}
}

func getDummyRequestAndResponse(
	playerName string) (*httptest.ResponseRecorder, *http.Request) {

	request, _ := http.NewRequest(http.MethodGet, prefixURLPlayer+playerName, nil)
	responseRecorder := httptest.NewRecorder()

	return responseRecorder, request
}

func checkPlayerScore(
	t *testing.T, request *http.Request,
	responseRecorder *httptest.ResponseRecorder, score string) {

	PlayerServer(responseRecorder, request)
	assertCorrectMessage(t, score, responseRecorder.Body.String())
}

func assertCorrectMessage(t *testing.T, want string, got string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
