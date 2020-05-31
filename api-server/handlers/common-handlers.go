package handlers

import (
	"net/http"
)

/*
   package to store common handlers like Not found handlers or somehting that can be used across all versions of the
   api that we define.
*/

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
