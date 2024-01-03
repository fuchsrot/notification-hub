package channel

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (a *API) List(w http.ResponseWriter, r *http.Request) {
	channels, err := a.repository.List()
	if err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	messages := *&channels[0].Messages

	a.logger.Info().Msgf("%i", len(messages))

	if len(channels) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(channels.ToDto()); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	form := &Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	newChannel := form.ToModel()
	newChannel.ID = uuid.New()

	channel, err := a.repository.Create(newChannel)
	if err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}

	a.logger.Info().Str("id", channel.ID.String()).Msg("new channel created")
	w.WriteHeader(http.StatusCreated)
}
