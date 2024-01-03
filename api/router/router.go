package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"fuchsrot/notification-hub/api/requestlog"
	"fuchsrot/notification-hub/api/resource/channel"
	"fuchsrot/notification-hub/api/resource/health"
	"fuchsrot/notification-hub/api/resource/message"
	"fuchsrot/notification-hub/api/router/middleware"
)

func New(l *zerolog.Logger, db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/live", health.Read)

	r.Route("/v1", func(r chi.Router) {
		r.Use(middleware.ContentTypeJson)

		channelAPI := channel.New(l, db)
		messageAPI := message.New(l, db)
		r.Method("GET", "/channels", requestlog.NewHandler(channelAPI.List, l))
		r.Method("POST", "/channels", requestlog.NewHandler(channelAPI.Create, l))

		r.Method("POST", "/channels/{channelId}/messages", requestlog.NewHandler(messageAPI.Create, l))
	})

	return r
}
