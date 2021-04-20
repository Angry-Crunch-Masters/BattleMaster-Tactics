package query

type BasicQuery struct {
	parameters map[string]interface{}
}

func InitBasicQuery() *BasicQuery {
	return &BasicQuery{parameters: make(map[string]interface{})}
}

func (query *BasicQuery) SetParameter(name string, property interface{}) {
	query.parameters[name] = property
}

func (query *BasicQuery) Execute(codes []PositionCode) []PositionCode {
	return codes
}
