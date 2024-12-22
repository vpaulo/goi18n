package server

import (
	"context"
	"encoding/json"
	"fmt"
	"goi18n/internal/i18n"
	"goi18n/web"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

// MiddlewareStack chains multiple middlewares
func MiddlewareStack(ms ...Middleware) Middleware {
	return Middleware(func(next http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			m := ms[i]
			next = m(next)
		}

		return next
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)
		log.Println(time.Since(start), r.Method, r.URL.Path)
	})
}

func TranslationsMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")

		// If lang is empty default to English, verifying length to make sure this example works
		if lang == "" || len(lang) > 2 {
			lang = "en"
		}

		data, err := web.Files.ReadFile(fmt.Sprintf("assets/i18n/%s.json", lang))
		if err != nil {
			log.Fatalf("error retrieving translations. Err: %v", err)
		}

		translations := i18n.Translation{}
		err = json.Unmarshal(data, &translations)

		if err != nil {
			log.Fatalf("error handling JSON Unmarshal translations. Err: %v", err)
		}

		ctx := context.WithValue(r.Context(), "translations", translations)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
