package types

// IMonitor ...
type IMonitor interface {
	Start()
	Status() (bool, error)
}
