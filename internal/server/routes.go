package server

import (
	"log"
	"net/http"

	"goi18n/internal/i18n"
	"goi18n/web/views"
)

func (s *Server) RegisterRoutes() http.Handler {
	stack := MiddlewareStack(
		Logger,
	)

	router := http.NewServeMux()
	router.Handle("GET /{$}", TranslationsMW(http.HandlerFunc(s.HomeViewHandler)))
	router.Handle("GET /{lang}/{$}", TranslationsMW(http.HandlerFunc(s.HomeViewHandler)))
	router.Handle("GET /{lang}/{slug}", TranslationsMW(http.HandlerFunc(s.HomeViewHandler)))

	return stack(router)
}

func (s *Server) HomeViewHandler(w http.ResponseWriter, r *http.Request) {
	translations := r.Context().Value("translations").(i18n.Translation)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := views.IndexView(translations).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in Home page: %e", err)
	}
}
