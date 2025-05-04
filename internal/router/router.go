package router

import "my-take-out/internal/router/admin"

type RouterGroup struct {
	admin.EmployeeRouter
}

var AllRouter = new(RouterGroup)
