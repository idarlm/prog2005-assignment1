package endpoints

import (
	"assignment1/internal/endpoints/diag"
	"assignment1/internal/endpoints/neighbour"
	"assignment1/internal/endpoints/stub"
	"assignment1/internal/endpoints/uniinfo"
	"fmt"
	"net/http"
	"time"
)

// Assign handler func to path
func SetHandle(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, handler)
	fmt.Println("Set endpoint:", path)
}

// Only allow GET METHOD for handler
func HandleGET(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is allowed.", http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	}
}

func SetApiEndpoints() {
	SetHandle(ROOT_PATH_V1+DIAG_PATH, HandleGET(diag.DiagHandler))
	SetHandle(ROOT_PATH_V1+UNIINFO_PATH, HandleGET(uniinfo.UniinfoHandler))
	SetHandle(ROOT_PATH_V1+NEIGHBOURUNI_PATH, HandleGET(neighbour.NeighbourHandler))

	diag.Timestamp = time.Now() // set timestamp for diag endpoint
}

func SetDebugEndpoints() {
	SetHandle("/stub/hipo", HandleGET(stub.HipoHandler))
	SetHandle("/name/Norway", HandleGET(stub.CountryHandler))
	SetHandle("/alpha", HandleGET(stub.CountryHandler))
}
