package message

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	e "fuchsrot/notification-hub/api/resource/common/err"
)

type API struct {
	logger     *zerolog.Logger
	repository *Repository
}

func New(logger *zerolog.Logger, db *gorm.DB) *API {
	return &API{
		logger:     logger,
		repository: NewRepository(db),
	}
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	form := &Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	channelID := chi.URLParam(r, "channelId")

	newMessage := form.ToModel()
	newMessage.ID = uuid.New()
	newMessage.Status = "NEW"
	newMessage.ChannelID = channelID

	message, err := a.repository.Create(newMessage)
	if err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}

	a.logger.Info().Str("id", message.ID.String()).Msg("new message created")
	w.WriteHeader(http.StatusCreated)
}
