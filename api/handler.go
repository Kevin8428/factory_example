package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (a *API) InitializeHandler(server *http.ServeMux) {
	server.Handle("/", a.Handle())
}

func (a *API) Handle() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		// localhost:8080/?entity=Cities&max_count=3&filter=international%3D0
		t := r.FormValue("entity")
		// mc := r.FormValue("max_count")
		f := url.QueryEscape(r.FormValue("filter"))
		fmt.Println("form values: ", t)
		values, err := a.transformer.Get(t)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		shouldFilter := a.transformer.ShouldFilter(t)

		if shouldFilter {
			v, err := a.transformer.Filter(values, f)
			if err == nil {
				values = v
			}
		}
		jData, err := json.Marshal(values)
		if err != nil {
			fmt.Println("error marshalling data ", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	})
}
