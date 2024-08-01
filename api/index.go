package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/xynrgys/googlesearch"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0 // or any default value you prefer
	}
	return i
}

func parseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false // or any default value you prefer
	}
	return b
}

func search(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	r.ParseForm() // Parse form data

	SearchTerm := r.FormValue("SearchTerm")
	opts := googlesearch.SearchOptions{
		CountryCode:    r.FormValue("CountryCode"),
		LanguageCode:   r.FormValue("LanguageCode"),
		Limit:          parseInt(r.FormValue("Limit")),
		Start:          parseInt(r.FormValue("Start")),
		UserAgent:      r.FormValue("UserAgent"),
		OverLimit:      parseBool(r.FormValue("OverLimit")),
		ProxyAddr:      r.FormValue("ProxyAddr"),
		FollowNextPage: parseBool(r.FormValue("FollowNextPage")),
	}

	results, err := googlesearch.Search(r.Context(), SearchTerm, opts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	jsonData, err := convertSearchResultsToJson(results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return jsonData, nil
}

func convertSearchResultsToJson(results []googlesearch.Result) ([]byte, error) {
	jsonData, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
