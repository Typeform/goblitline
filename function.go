package goblitline

import "github.com/lann/builder"

type function struct {
	Name      string     `json:"name"`
	Params    []string   `json:"params,omitempty"`
	Functions []function `json:"function,omitempty"`
	Container *container `json:"save,omitempty"`
}

type functionBuilder builder.Builder

func (b functionBuilder) Name(name string) functionBuilder {
	return builder.Set(b, "Name", name).(functionBuilder)
}

func (b functionBuilder) Params(params ...string) functionBuilder {
	for _, param := range params {
		b = builder.Append(b, "Params", param).(functionBuilder)
	}
	return b
}

func (b functionBuilder) Save(c containerBuilder) functionBuilder {
	cont := c.build()
	return builder.Set(b, "Container", &cont).(functionBuilder)
}

func (b functionBuilder) build() function {
	return builder.GetStruct(b).(function)
}

var Function = builder.Register(
	functionBuilder{},
	function{}).(functionBuilder)
