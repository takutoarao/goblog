package main

import (
	"fmt"

	"git.si-media.biz/takutoarao/start/app"
	"github.com/goadesign/goa"
)

// BlogController implements the blog resource.
type BlogController struct {
	*goa.Controller
}

// NewBlogController creates a blog controller.
func NewBlogController(service *goa.Service) *BlogController {
	return &BlogController{Controller: service.NewController("BlogController")}
}

// Show runs the show action.
func (c *BlogController) Show(ctx *app.ShowBlogContext) error {
	if ctx.BlogID == 0 {
		// Emulate a missing record with ID 0
		return ctx.NotFound()
	}
	// Build the resource using the generated data structure
	blog := app.GoaExampleBlog{
		ID:   ctx.BlogID,
		Name: fmt.Sprintf("Blog #%d", ctx.BlogID),
		Href: app.BlogHref(ctx.BlogID),
	}

	// Let the generated code produce the HTTP response using the
	// media type described in the design (BlogMedia).
	return ctx.OK(&blog)
}
