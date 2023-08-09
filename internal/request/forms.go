package request

import (
	"errors"
	"net/http"

	"github.com/go-playground/form/v4"
)

var decoder = form.NewDecoder()

func DecodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = decoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
	}

	return err
}

func DecodeQueryString(r *http.Request, dst any) error {
	err := decoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
	}

	return err
}
