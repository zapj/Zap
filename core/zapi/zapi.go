package zapi

var Client *Zapi = &Zapi{}

type Zapi struct {
}

func NewZapi() *Zapi {
	return &Zapi{}
}
