package controllers

import "gin/internal/service"

type ChildController struct {
	childservice *service.ChildService
}

func NewChildController(service *service.ChildService) *ChildController {
	return &ChildController{childservice: service}
}
