package controller

type HttpResponse struct {
	StatusCode  int
	Body        any
	ContentType string
	Headers     map[string]string
}
