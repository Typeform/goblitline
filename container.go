package goblitline

import "github.com/lann/builder"

type containerData struct {
	ImageIdentifier string `json:"image_identifier"`
	Quality         uint   `json:"quality"`
}

type ContainerBuilder builder.Builder

func init() {
	builder.Register(ContainerBuilder{}, containerData{})
}

func (b ContainerBuilder) ImageIdentifier(id string) ContainerBuilder {
	return builder.Set(b, "ImageIdentifier", id).(ContainerBuilder)
}

func (b ContainerBuilder) Quality(quality uint) ContainerBuilder {
	return builder.Set(b, "Quality", quality).(ContainerBuilder)
}

func (b ContainerBuilder) build() containerData {
	return builder.GetStruct(b).(containerData)
}
