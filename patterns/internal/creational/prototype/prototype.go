package prototype

type Prototype interface {
	Clone() Prototype
}

type InfoGetter interface {
	GetInfo() string
}
