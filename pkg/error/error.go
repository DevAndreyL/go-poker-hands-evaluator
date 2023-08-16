package pokererr

import (
	"encoding/json"
	"fmt"
)

const internalMessageKey = "message"

type Data map[string]any

type Error struct {
	Code   Code
	Data   Data
	Source error
}

func NewError(code Code, data Data) *Error {
	if data == nil {
		data = make(Data)
	}

	return &Error{Code: code, Data: data}
}

func Wrap(err error, code Code, data Data) *Error {
	newErr := NewError(code, data)
	newErr.Source = err

	return newErr
}

func (e Error) Error() string {
	return string(e.Code)
}

func (e Error) Is(err error) bool {
	return e.Error() == err.Error()
}

func (e Error) As(err any) bool {
	if ginErr, ok := err.(*Error); ok {
		return ginErr.Code == e.Code
	}

	return false
}

func (e Error) Unwrap() error {
	return e.Source
}

func (e Error) String() string {
	return fmt.Sprintf("Code: %s, Data: %s, Source: %s", e.Code, e.Data, e.Source)
}

func (e Error) MarshalJSON() ([]byte, error) {
	type inner struct {
		Code   Code   `json:"code"`
		Data   Data   `json:"data"`
		Source *Error `json:"source"`
	}

	if e.Source == nil {
		return json.Marshal(&inner{
			Code: e.Code,
			Data: e.Data,
		})
	}

	if ginErr, ok := e.Source.(*Error); ok {
		return json.Marshal(&inner{
			Code:   e.Code,
			Data:   e.Data,
			Source: ginErr,
		})
	}

	return json.Marshal(
		Wrap(
			NewError(CodeGeneralError, Data{internalMessageKey: e.Source.Error()}),
			e.Code,
			e.Data,
		),
	)
}

func (e *Error) UnmarshalJSON(data []byte) error {
	var inner struct {
		Code   Code   `json:"code"`
		Data   Data   `json:"data"`
		Source *Error `json:"source"`
	}

	if err := json.Unmarshal(data, &inner); err != nil {
		return err
	}

	e.Code = inner.Code
	e.Data = inner.Data
	if inner.Source != nil {
		e.Source = inner.Source
	}

	return nil
}
