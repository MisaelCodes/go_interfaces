package advanced_level

import "net/http"

type HelloHandler struct {
    Data string
}

func (hh *HelloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request){
    rw.Write([]byte(hh.Data))
}
