package goblitline

import "github.com/lann/builder"

type functionData struct {
	Name      string         `json:"name"`
	Params    []string       `json:"params,omitempty"`
	Functions []functionData `json:"function,omitempty"`
	Container *containerData `json:"save,omitempty"`
}

type FunctionBuilder builder.Builder

func init() {
	builder.Register(FunctionBuilder{}, functionData{})
}

func (b FunctionBuilder) Name(name string) FunctionBuilder {
	return builder.Set(b, "Name", name).(FunctionBuilder)
}

func (b FunctionBuilder) Params(params ...string) FunctionBuilder {
	for _, param := range params {
		b = builder.Append(b, "Params", param).(FunctionBuilder)
	}
	return b
}

func (b FunctionBuilder) Save(c ContainerBuilder) FunctionBuilder {
	cont := c.build()
	return builder.Set(b, "Container", &cont).(FunctionBuilder)
}

func (b FunctionBuilder) build() functionData {
	return builder.GetStruct(b).(functionData)
}
