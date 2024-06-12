package api

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/leandrofars/oktopus/internal/bridge"
	"github.com/leandrofars/oktopus/internal/cwmp"
	n "github.com/leandrofars/oktopus/internal/nats"
	"github.com/leandrofars/oktopus/internal/utils"
	"github.com/nats-io/nats.go"
)

func (a *Api) cwmpGetParameterNamesMsg(w http.ResponseWriter, r *http.Request) {
	sn := getSerialNumberFromRequest(r)

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.Marshall(err.Error()))
		return
	}

	data, _, err := cwmpInteraction[cwmp.GetParameterNamesResponse](sn, payload, w, a.nc)
	if err != nil {
		return
	}

	w.Write(data)
}

func (a *Api) cwmpGetParameterAttributesMsg(w http.ResponseWriter, r *http.Request) {
	sn := getSerialNumberFromRequest(r)

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.Marshall(err.Error()))
		return
	}

	data, _, err := cwmpInteraction[cwmp.GetParameterAttributesResponse](sn, payload, w, a.nc)
	if err != nil {
		return
	}

	w.Write(data)
}

func (a *Api) cwmpGetParameterValuesMsg(w http.ResponseWriter, r *http.Request) {
	sn := getSerialNumberFromRequest(r)

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.Marshall(err.Error()))
		return
	}

	data, _, err := cwmpInteraction[cwmp.GetParameterValuesResponse](sn, payload, w, a.nc)
	if err != nil {
		return
	}

	w.Write(data)
}

func (a *Api) cwmpSetParameterValuesMsg(w http.ResponseWriter, r *http.Request) {
	sn := getSerialNumberFromRequest(r)

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.Marshall(err.Error()))
		return
	}

	data, _, err := cwmpInteraction[cwmp.SetParameterValuesResponse](sn, payload, w, a.nc)
	if err != nil {
		return
	}

	w.Write(data)
}

func (a *Api) cwmpAddObjectMsg(w http.ResponseWriter, r *http.Request) {
	sn := getSerialNumberFromRequest(r)

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.Marshall(err.Error()))
		return
	}

	data, _, err := cwmpInteraction[cwmp.AddObjectResponse](sn, payload, w, a.nc)
	if err != nil {
		return
	}

	w.Write(data)
}

func cwmpInteraction[T cwmp.SetParameterValuesResponse | cwmp.GetParameterAttributesResponse | cwmp.GetParameterNamesResponse | cwmp.GetParameterValuesResponse | cwmp.AddObjectResponse](
	sn string, payload []byte, w http.ResponseWriter, nc *nats.Conn,
) ([]byte, T, error) {

	var response T

	data, err := bridge.NatsCwmpInteraction(
		n.NATS_CWMP_ADAPTER_SUBJECT_PREFIX+sn+".api",
		payload,
		w,
		nc,
	)
	if err != nil {
		return data, response, err
	}

	err = xml.Unmarshal(data, &response)
	if err != nil {
		err = json.Unmarshal(data, &response)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.Marshall(err))
		}
	}
	return data, response, err
}
