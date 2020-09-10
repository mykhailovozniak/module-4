package module_4

import "fmt"

type Handler interface {
	SendRequest(message int) string
}

type InternalErrorHandler struct {
	next Handler
}

func (h *InternalErrorHandler) SendRequest(code int) (result string) {
	if code == 500 {
		result = "Internal Error Happened"
	} else if h.next != nil {
		result = h.next.SendRequest(code)
	}
	return
}

type UnauthorizedHandler struct {
	next Handler
}

func (h *UnauthorizedHandler) SendRequest(code int) (result string) {
	if code == 401 {
		result = "Unauthorized Error Happened"
	} else if h.next != nil {
		result = h.next.SendRequest(code)
	}
	return
}

func chainOfResponsibilityClientCode() {
	handler := &UnauthorizedHandler{
		next: &InternalErrorHandler{},
	}

	unauthorizedError := handler.SendRequest(401)
	fmt.Println(unauthorizedError)

	internalError := handler.SendRequest(500)
	fmt.Println(internalError)
}
