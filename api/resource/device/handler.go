package device

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
    "github.com/ggicci/httpin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	q "northpole-shop/api/resource/common/query"
	e "northpole-shop/api/resource/common/err"
	l "northpole-shop/api/resource/common/log"
	ctxUtil "northpole-shop/util/ctx"
)

type QueryDeviceList = q.QueryDeviceList

type API struct {
	logger     *zerolog.Logger
	validator  *validator.Validate
	repository *Repository
}

func New(logger *zerolog.Logger, validator *validator.Validate, db *gorm.DB) *API {
	return &API{
		logger:     logger,
		validator:  validator,
		repository: NewRepository(db),
	}
}


// List godoc
//
//	@summary		List devices
//	@description	List devices
//	@tags			devices
//	@accept			json
//	@produce		json
//	@param			startTime	query	int	false	"Start date"
//	@param			endTime 	query	int	false	"End date"
//	@param			locationId	query	string	false	"Location UUID"
//	@success		200	{array}		DeviceDTO
//	@failure		500	{object}	err.Error
//	@router			/devices [get]
func (a *API) List(w http.ResponseWriter, r *http.Request) {
	reqID := ctxUtil.RequestID(r.Context())
    input := r.Context().Value(httpin.Input).(*QueryDeviceList)

	// startTime, _ := strconv.ParseInt(chi.URLParam(r, "startTime"),10, 64)
	// endTime, _ := strconv.ParseInt(chi.URLParam(r, "endTime"),10, 64)
	// locationId, _ := uuid.Parse(chi.URLParam(r, "locationId"))

	a.logger.Info().Msgf("startTime: %v", input.StartTime)

	devices, err := a.repository.List(input.StartTime, input.EndTime, input.LocationId)
	if err != nil {
		a.logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	if len(devices) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(devices.ToDto()); err != nil {
		a.logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

