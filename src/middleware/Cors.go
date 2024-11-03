package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var origins = strings.Split(os.Getenv("ORIGINS"), ",")

/*Esta funcion pasarla a un helpers y renombrarle como Some para reusarlas*/
func contains(list []string, value string) bool {
	for _, s := range list {
		if s == value {
			return true
		}
	}
	return false
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		originalUrl := r.URL.Path
		method := r.Method
		origin := r.Header.Get("Origin")

		log.Printf("Enter on Custom cors with the url: %s", originalUrl)

		if origin != "" {
			isAllowed := contains(origins, origin)

			if !isAllowed {
				http.Error(w, "Blocked by CORS policy.", http.StatusUnauthorized)
				return
			}

			headers := map[string]string{
				"Access-Control-Allow-Origin":      origin,
				"Access-Control-Allow-Methods":     "GET, POST, OPTIONS, PUT, PATCH, DELETE",
				"Access-Control-Allow-Headers":     "Access-Control-Allow-Headers,access-control-allow-origin, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, token",
				"Access-Control-Allow-Credentials": "true",
			}

			for k, v := range headers {
				w.Header().Set(k, v)
			}
		}

		if method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
