package middleware

import (
	"github.com/FerdinaKusumah/fastcrud/utils"
	"net/http"
)

func SetRespJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(out http.ResponseWriter, in *http.Request) {
		// set cors
		out.Header().Set("Content-Type", "application/json")
		out.Header().Set("Access-Control-Allow-Headers", "*")
		out.Header().Set("Access-Control-Allow-Origin", "*")
		if in.Method == http.MethodOptions {
			return
		}

		// logging data every request
		utils.SendLogInfo(in.URL.String())

		out.WriteHeader(http.StatusOK)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(out, in)
	})
}
