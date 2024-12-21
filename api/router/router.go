package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/go-playground/validator/v10"
    "github.com/ggicci/httpin"
	"gorm.io/gorm"

	q "northpole-shop/api/resource/common/query"

	"northpole-shop/config"
	"northpole-shop/api/resource/device"
	"northpole-shop/api/router/middleware"
	"northpole-shop/api/router/middleware/requestlog"
)
type QueryDeviceList = q.QueryDeviceList

func New(c *config.Conf, l *zerolog.Logger, v *validator.Validate, db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Route("/api", func(r chi.Router) {
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
