package device

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
    "github.com/ggicci/httpin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	q "northpole-shop/api/resource/common/query"
	e "northpole-shop/api/resource/common/err"
	l "northpole-shop/api/resource/common/log"
	ctxUtil "northpole-shop/util/ctx"
	validatorUtil "northpole-shop/util/validator"
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

// Create godoc
//
//	@summary		Create device
//	@description	Create device
//	@tags			devices
//	@accept			json
//	@produce		json
//	@param			body	body	DeviceForm	true	"Device form"
//	@success		201
//	@failure		400	{object}	err.Error
//	@failure		422	{object}	err.Errors
//	@failure		500	{object}	err.Error
//	@router			/devices [post]
func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	reqID := ctxUtil.RequestID(r.Context())

	form := &DeviceForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		a.logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	if err := a.validator.Struct(form); err != nil {
		respBody, err := json.Marshal(validatorUtil.ToErrResponse(err))
		if err != nil {
			a.logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("")
			e.ServerError(w, e.RespJSONEncodeFailure)
			return
		}

		e.ValidationErrors(w, respBody)
		return
	}

	newDevice := form.ToModel()
	newDevice.ID = uuid.New()

	device, err := a.repository.Create(newDevice)
	if err != nil {
		a.logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}

	a.logger.Info().Str(l.KeyReqID, reqID).Str("id", device.ID.String()).Msg("new device created")
	w.WriteHeader(http.StatusCreated)
}

// Read godoc
//
//	@summary		Read device
//	@description	Read device
//	@tags			devices
//	@accept			json
//	@produce		json
//	@param			id	path		string	true	"Device ID"
//	@success		200	{object}	DeviceDTO
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		500	{object}	err.Error
//	@router			/devices/{id} [get]
func (a *API) Read(w http.ResponseWriter, r *http.Request) {
	reqID := ctxUtil.RequestID(r.Context())

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	device, err := a.repository.Read(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		a.logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	dto := device.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		a.logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

// Update godoc
//
//	@summary		Update device
//	@description	Update device
//	@tags			devices
//	@accept			json
//	@produce		json
//	@param			id		path	string	true	"Device ID"
//	@param			body	body	DeviceForm	true	"Device form"
//	@success		200
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		422	{object}	err.Errors
//	@failure		500	{object}	err.Error
//	@router			/devices/{id} [put]
func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	reqID := ctxUtil.RequestID(r.Context())

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	form := &DeviceForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		a.logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	if err := a.validator.Struct(form); err != nil {
		respBody, err := json.Marshal(validatorUtil.ToErrResponse(err))
		if err != nil {
			a.logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("")
			e.ServerError(w, e.RespJSONEncodeFailure)
			return
		}

		e.ValidationErrors(w, respBody)
		return
	}

	device := form.ToModel()
	device.ID = id

	rows, err := a.repository.Update(device)
	if err != nil {
		a.logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataUpdateFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	a.logger.Info().Str(l.KeyReqID, reqID).Str("id", id.String()).Msg("device updated")
}

// Delete godoc
//
//	@summary		Delete device
//	@description	Delete device
//	@tags			devices
//	@accept			json
//	@produce		json
//	@param			id	path	string	true	"Device ID"
//	@success		200
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		500	{object}	err.Error
//	@router			/devices/{id} [delete]
func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	reqID := ctxUtil.RequestID(r.Context())

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	rows, err := a.repository.Delete(id)
	if err != nil {
		a.logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("")
		e.ServerError(w, e.RespDBDataRemoveFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	a.logger.Info().Str(l.KeyReqID, reqID).Str("id", id.String()).Msg("device deleted")
}
