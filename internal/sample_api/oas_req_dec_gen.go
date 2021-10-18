// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

func decodeFoobarPostRequest(r *http.Request) (req *Pet, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Pet
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}
		if err := request.validate(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}

		return &request, nil
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodePetCreateRequest(r *http.Request) (req PetCreateRequest, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request Pet
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}
		if err := request.validate(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}

		return &request, nil
	case "text/plain":
		var request PetCreateTextPlainRequest
		_ = request
		return req, fmt.Errorf("text/plain decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}