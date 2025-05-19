package router

import "my-take-out/internal/router/admin"

type RouterGroup struct {
	admin.EmployeeRouter
	admin.CommonRouter
}

var AllRouter = new(RouterGroup)
