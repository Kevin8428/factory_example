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
		// localhost:8080/?entity=cities&max_count=3&filter=isInternational
		// mc := r.FormValue("max_count")

		e := r.FormValue("entity")
		f := url.QueryEscape(r.FormValue("filter"))

		results, err := a.transformer.Get(e)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		shouldFilter, err := a.transformer.ShouldFilter(e, f)
		fmt.Println("shouldFilter: ", shouldFilter)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		if shouldFilter {
			fmt.Println("running filter")
			v, err := a.transformer.Filter(e, results, f)
			if err == nil {
				results = v
			}
		}
		jData, err := json.Marshal(results)
		if err != nil {
			fmt.Println("error marshalling data ", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	})
}
