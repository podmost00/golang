package handlers

import (
	"Lab7_RestAPI/packages/templates"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(
		w,
		templates.PAGE_START,
		templates.INDEX_PAGE,
	)
	fmt.Fprint(w, templates.PAGE_END)
}
