package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("blog", func() { // API defines the microservice endpoint and
	Title("The virtual wine cellar")    // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
	Scheme("http")                      // the design.
	Host("localhost:8080")
})

var _ = Resource("blog", func() { // Resources group related API endpoints
	BasePath("/blogs")      // together. They map to REST resources for REST
	DefaultMedia(BlogMedia) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get blog by id") // with its path, parameters (both path
		Routing(GET("/:blogID"))      // parameters and querystring values) and payload
		Params(func() {               // (shape of the request body).
			Param("blogID", Integer, "Blog ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

// BlogMedia defines the media type used to render blogs.
var BlogMedia = MediaType("application/vnd.goa.example.blog+json", func() {
	Description("A blog of wine")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique blog ID")
		Attribute("href", String, "API href for making requests on the blog")
		Attribute("name", String, "Name of wine")
		Required("id", "href", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
	})
})
