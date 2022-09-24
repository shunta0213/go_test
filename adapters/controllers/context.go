package controllers

// Gin Context

type GinContext interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	Query(string) string
	JSON(int, interface{})
}
