package problems

import (
	"net/http"
	"strconv"
)

type Filters struct {
	search string
    rated bool
    maxRating *int
    minRating *int
}

func stringToOptionalInt(s string) *int {
    v, err := strconv.Atoi(s)
    if err != nil {
        return nil
    }
    return &v
}

func filtersFromRequest(r *http.Request) Filters {
    maxRating := stringToOptionalInt(r.FormValue("max_rating"))
    minRating := stringToOptionalInt(r.FormValue("min_rating"))

	return Filters{
        search: r.FormValue("search"),
        rated: r.FormValue("rated") == "on",
        maxRating: maxRating,
        minRating: minRating,
	}
}
