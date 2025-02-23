package key

import (
	"errors"
	"net/http"

	httperror "github.com/portainer/portainer/pkg/libhttp/error"
	"github.com/portainer/portainer/pkg/libhttp/response"
)

type keyInspectResponse struct {
	Key string `json:"key"`
}

func (handler *Handler) keyInspect(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	if handler.edgeManager == nil {
		return &httperror.HandlerError{StatusCode: http.StatusServiceUnavailable, Message: "Edge key management is disabled on non Edge agent", Err: errors.New("Edge key management is disabled")}
	}

	if !handler.edgeManager.IsKeySet() {
		return httperror.NotFound("No key associated to this agent", errors.New("Edge key unavailable"))
	}

	edgeKey := handler.edgeManager.GetKey()

	return response.JSON(w, keyInspectResponse{
		Key: edgeKey,
	})
}
