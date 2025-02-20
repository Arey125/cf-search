package problems

import "net/http"

type Filters struct {
	search string
    rated bool
}

func filtersFromRequest(r *http.Request) Filters {
	return Filters{
        search: r.FormValue("search"),
        rated: r.FormValue("rated") == "on",
	}
}
