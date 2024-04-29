package pkg

type Shape interface {
	getType() string
	Accept(Visitor)
}
