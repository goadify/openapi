// Code generated by ogen, DO NOT EDIT.

package gen

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// EntityMappingsGet implements GET /entity_mappings operation.
	//
	// Returns a mappings of loaded entities.
	//
	// GET /entity_mappings
	EntityMappingsGet(ctx context.Context) ([]EntityMapping, error)
	// EntityNameGet implements GET /entity/{name} operation.
	//
	// Retrieves records with pagination.
	//
	// GET /entity/{name}
	EntityNameGet(ctx context.Context, params EntityNameGetParams) (*EntitiesResponse, error)
	// EntityNameIDDelete implements DELETE /entity/{name}/{id} operation.
	//
	// Deletes record.
	//
	// DELETE /entity/{name}/{id}
	EntityNameIDDelete(ctx context.Context, req *Entity, params EntityNameIDDeleteParams) error
	// EntityNameIDGet implements GET /entity/{name}/{id} operation.
	//
	// Retrieves one record by identifier.
	//
	// GET /entity/{name}/{id}
	EntityNameIDGet(ctx context.Context, params EntityNameIDGetParams) error
	// EntityNameIDPut implements PUT /entity/{name}/{id} operation.
	//
	// Updates existing record.
	//
	// PUT /entity/{name}/{id}
	EntityNameIDPut(ctx context.Context, req *Entity, params EntityNameIDPutParams) error
	// EntityNamePost implements POST /entity/{name} operation.
	//
	// Creates record.
	//
	// POST /entity/{name}
	EntityNamePost(ctx context.Context, req *Entity, params EntityNamePostParams) error
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}