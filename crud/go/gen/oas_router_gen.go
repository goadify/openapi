// Code generated by ogen, DO NOT EDIT.

package gen

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [2]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/entity"
			if l := len("/entity"); len(elem) >= l && elem[0:l] == "/entity" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case '/': // Prefix: "/"
				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "name"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleGetRecordsRequest([1]string{
							args[0],
						}, elemIsEscaped, w, r)
					case "POST":
						s.handleCreateRecordRequest([1]string{
							args[0],
						}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[1] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "DELETE":
							s.handleDeleteRecordByIdRequest([2]string{
								args[0],
								args[1],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleGetRecordByIdRequest([2]string{
								args[0],
								args[1],
							}, elemIsEscaped, w, r)
						case "PUT":
							s.handleUpdateRecordByIdRequest([2]string{
								args[0],
								args[1],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PUT")
						}

						return
					}
				}
			case '_': // Prefix: "_mappings"
				if l := len("_mappings"); len(elem) >= l && elem[0:l] == "_mappings" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleGetEntitiesMappingsRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	pathPattern string
	count       int
	args        [2]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/entity"
			if l := len("/entity"); len(elem) >= l && elem[0:l] == "/entity" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case '/': // Prefix: "/"
				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "name"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "GetRecords"
						r.operationID = "GetRecords"
						r.pathPattern = "/entity/{name}"
						r.args = args
						r.count = 1
						return r, true
					case "POST":
						r.name = "CreateRecord"
						r.operationID = "CreateRecord"
						r.pathPattern = "/entity/{name}"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[1] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							// Leaf: DeleteRecordById
							r.name = "DeleteRecordById"
							r.operationID = "DeleteRecordById"
							r.pathPattern = "/entity/{name}/{id}"
							r.args = args
							r.count = 2
							return r, true
						case "GET":
							// Leaf: GetRecordById
							r.name = "GetRecordById"
							r.operationID = "GetRecordById"
							r.pathPattern = "/entity/{name}/{id}"
							r.args = args
							r.count = 2
							return r, true
						case "PUT":
							// Leaf: UpdateRecordById
							r.name = "UpdateRecordById"
							r.operationID = "UpdateRecordById"
							r.pathPattern = "/entity/{name}/{id}"
							r.args = args
							r.count = 2
							return r, true
						default:
							return
						}
					}
				}
			case '_': // Prefix: "_mappings"
				if l := len("_mappings"); len(elem) >= l && elem[0:l] == "_mappings" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: GetEntitiesMappings
						r.name = "GetEntitiesMappings"
						r.operationID = "GetEntitiesMappings"
						r.pathPattern = "/entity_mappings"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			}
		}
	}
	return r, false
}
