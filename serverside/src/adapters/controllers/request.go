package controllers

import "github.com/TechDepa/c_tool/usecase"

type Request interface {
	RenderJSON(sc usecase.Status, body interface{})
	RenderAbortWithError(sc usecase.Status, err error)
	BindJSON(obj interface{}) error
}
