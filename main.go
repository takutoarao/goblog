//go:generate goagen bootstrap -d git.si-media.biz/takutoarao/start/design

package main

import (
	"git.si-media.biz/takutoarao/start/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// setup()
	// Create service
	service := goa.New("blog")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "blog" controller
	c := NewBlogController(service)
	app.MountBlogController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
