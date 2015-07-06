package goblitline

import "github.com/lann/builder"

type functionData struct {
	Name      string                 `json:"name"`
	Params    map[string]interface{} `json:"params,omitempty"`
	Functions []functionData         `json:"function,omitempty"`
	Container *containerData         `json:"save,omitempty"`
}

type FunctionBuilder builder.Builder

func init() {
	builder.Register(FunctionBuilder{}, functionData{})
}

func (b FunctionBuilder) Name(name string) FunctionBuilder {
	return builder.Set(b, "Name", name).(FunctionBuilder)
}

func (b FunctionBuilder) Params(key string, value interface{}) FunctionBuilder {
	var hash map[string]interface{}
	params, ok := builder.Get(b, "Params")
	if ok {
		hash = params.(map[string]interface{})
	} else {
		hash = make(map[string]interface{})
	}
	hash[key] = value
	return builder.Set(b, "Params", hash).(FunctionBuilder)
}

func (b FunctionBuilder) Save(c ContainerBuilder) FunctionBuilder {
	cont := c.build()
	return builder.Set(b, "Container", &cont).(FunctionBuilder)
}

func (b FunctionBuilder) build() functionData {
	return builder.GetStruct(b).(functionData)
}
