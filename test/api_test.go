package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/5anthosh/nottu/db/status"
	"github.com/5anthosh/nottu/server"
)

var router, _ = server.TestBuild()

func BenchmarkGetNotes(t *testing.B) {

	req, _ := http.NewRequest("GET", "/notes", nil)
	t.ReportAllocs()
	response := executeRequest(req)
	if response.Code == status.OK {
		println(response)
	}
}
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}
