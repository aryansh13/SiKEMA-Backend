package class

import service "attendance-is/services"

type handler struct {
	service service.ClassService
}

func NewClassHandler(service service.ClassService) handler {
	return handler{service: service}
}
