package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	ErrContentType = errors.New("content-type error")
	ErrRead        = errors.New("read error")
	ErrWrite       = errors.New("write error")
	ErrUnmarshal   = errors.New("unmarshal error")
	ErrMarshal     = errors.New("marshal error")
)

var JSONType = "application/json"
var Type = "Content-Type"

func ReadJSONBody(request *http.Request, dto interface{}) (err error) {
	if request.Header.Get(Type) != JSONType {
		return ErrContentType
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return ErrRead
	}
	defer func() {
		err = request.Body.Close()
	}()

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return ErrUnmarshal
	}
	return nil
}

func WriteJSONBody(response http.ResponseWriter, dto interface{}) (err error) {
	response.Header().Set(Type, JSONType)

	body, err := json.Marshal(dto)
	if err != nil {
		return ErrMarshal
	}

	_, err = response.Write(body)
	if err != nil {
		return ErrWrite
	}

	return nil
}
