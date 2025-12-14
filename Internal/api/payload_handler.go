package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PayloadHandler struct {
}

func NewPaylaodHandler() *PayloadHandler {
	return &PayloadHandler{}
}

func (p *PayloadHandler) HandleGetPayloadById(w http.ResponseWriter, r *http.Request) {
	paramsPayloadId := chi.URLParam(r, "id")

	if paramsPayloadId == "" {
		http.NotFound(w, r)
		return
	}

	payloadId, err := strconv.ParseInt(paramsPayloadId, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "PAYLOAD ID: %d IS ON THE WAY STATION 28 ON LUNA\n", payloadId)
}

func (p *PayloadHandler) HandleCreatePayload(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CREATED A NEW PAYLOAD!\n")
}
