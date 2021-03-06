// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "blog": Application Contexts
//
// Command:
// $ goagen
// --design=git.si-media.biz/takutoarao/start/design
// --out=$(GOPATH)/src/git.si-media.biz/takutoarao/start
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// ShowBlogContext provides the blog show action context.
type ShowBlogContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	BlogID int
}

// NewShowBlogContext parses the incoming request URL and body, performs validations and creates the
// context used by the blog controller show action.
func NewShowBlogContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowBlogContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowBlogContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramBlogID := req.Params["blogID"]
	if len(paramBlogID) > 0 {
		rawBlogID := paramBlogID[0]
		if blogID, err2 := strconv.Atoi(rawBlogID); err2 == nil {
			rctx.BlogID = blogID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("blogID", rawBlogID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowBlogContext) OK(r *GoaExampleBlog) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.blog+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowBlogContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
