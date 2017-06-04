// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "blog": blog Resource Client
//
// Command:
// $ goagen
// --design=git.si-media.biz/takutoarao/start/design
// --out=$(GOPATH)/src/git.si-media.biz/takutoarao/start
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ShowBlogPath computes a request path to the show action of blog.
func ShowBlogPath(blogID int) string {
	param0 := strconv.Itoa(blogID)

	return fmt.Sprintf("/blogs/%s", param0)
}

// Get blog by id
func (c *Client) ShowBlog(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBlogRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBlogRequest create the request corresponding to the show action endpoint of the blog resource.
func (c *Client) NewShowBlogRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
