package controllers

import (
	"net/http"

	"dv/pkg/htemplate"
)

func Index(ht *htemplate.HTemplate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Title": "Data Vault",
			"Body":  "Am on the way up!",
		}

		err := ht.Execute(w, "index.html", data)
		if err != nil {
			// Log or handle error if needed
		}
	}
}
