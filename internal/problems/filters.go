package problems

import "net/http"

type Filters struct {
	search string
}

func filtersFromRequest(r *http.Request) Filters {
	return Filters{
		r.FormValue("search"),
	}
}
