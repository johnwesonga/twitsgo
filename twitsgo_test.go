package twitsgo

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var request = struct {
	path, query       string // request
	contenttype, body string // response
}{
	path:        "/search.json?",
	query:       "q=%23Kenya",
	contenttype: "application/json",
	body:        twitterResponse,
}

var (
	twitterResponse = `{ 'results': [{'text':'hello','id_str':'34455w4','from_user_name':'bob','from_user_id_str':'345424'}]}`
)

func TestRetrieveTweets(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", request.contenttype)
		io.WriteString(w, request.body)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	checkBody(t, resp, twitterResponse)
}

func checkBody(t *testing.T, r *http.Response, body string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error("reading reponse body: %v, want %q", err, body)
	}
	if g, w := string(b), body; g != w {
		t.Errorf("request body mismatch: got %q, want %q", g, w)
	}
}
