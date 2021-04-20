package query

type IQuery interface {
	SetParameter(name string, property interface{})
	Execute(codes []PositionCode) []PositionCode
}
