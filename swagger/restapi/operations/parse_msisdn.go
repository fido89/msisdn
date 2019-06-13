// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ParseMsisdnHandlerFunc turns a function with the right signature into a parse msisdn handler
type ParseMsisdnHandlerFunc func(ParseMsisdnParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ParseMsisdnHandlerFunc) Handle(params ParseMsisdnParams) middleware.Responder {
	return fn(params)
}

// ParseMsisdnHandler interface for that can handle valid parse msisdn params
type ParseMsisdnHandler interface {
	Handle(ParseMsisdnParams) middleware.Responder
}

// NewParseMsisdn creates a new http.Handler for the parse msisdn operation
func NewParseMsisdn(ctx *middleware.Context, handler ParseMsisdnHandler) *ParseMsisdn {
	return &ParseMsisdn{Context: ctx, Handler: handler}
}

/*ParseMsisdn swagger:route GET /parse parseMsisdn

ParseMsisdn parse msisdn API

*/
type ParseMsisdn struct {
	Context *middleware.Context
	Handler ParseMsisdnHandler
}

func (o *ParseMsisdn) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewParseMsisdnParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}