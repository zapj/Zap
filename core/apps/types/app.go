package types

type IApp interface {
	Stop() error
	Start() error
	Restart() error
	Reload() error
}
