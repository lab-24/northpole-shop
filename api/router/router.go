package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/go-playground/validator/v10"
    "github.com/ggicci/httpin"
	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"

	q "northpole-shop/api/resource/common/query"

	"northpole-shop/config"
	"northpole-shop/api/resource/device"
	"northpole-shop/api/router/middleware"
	"northpole-shop/api/router/middleware/requestlog"
)
type QueryDeviceList = q.QueryDeviceList

func New(c *config.Conf, l *zerolog.Logger, v *validator.Validate, db *gorm.DB) *chi.Mux {
	tokenAuth := jwtauth.New("HS256", []byte(c.Auth.JwtSecret), nil)
	r := chi.NewRouter()

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	if c.Auth.Debug{
		_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
		l.Warn().Msgf("JWT_DEBUG=true")
		l.Info().Msgf("A sample jwt is %s", tokenString)
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Route("/api", func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))
		// Handle valid / invalid tokens
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Use(middleware.RequestID)
		r.Use(middleware.ContentTypeJSON)
		deviceAPI := device.New(l, v, db)
		r.With(httpin.NewInput(QueryDeviceList{})).Method(http.MethodGet, "/devices", requestlog.NewHandler(deviceAPI.List, l))
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome anonymous"))
		})
	})

	return r
}
