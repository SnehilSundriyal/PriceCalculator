package iomanager

type IOmanager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}