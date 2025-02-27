package receivers

import "net/http"

type HTTPReciever struct {
	mux http.ServeMux
}
