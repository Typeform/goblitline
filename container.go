package goblitline

import "github.com/lann/builder"

type container struct {
	ImageIdentifier string `json:"image_identifier"`
	Quality         uint   `json:"quality"`
}

type containerBuilder builder.Builder

func (b containerBuilder) ImageIdentifier(id string) containerBuilder {
	return builder.Set(b, "ImageIdentifier", id).(containerBuilder)
}

func (b containerBuilder) Quality(quality uint) containerBuilder {
	return builder.Set(b, "Quality", quality).(containerBuilder)
}

func (b containerBuilder) build() container {
	return builder.GetStruct(b).(container)
}

var Container = builder.Register(
	containerBuilder{},
	container{
		Quality: 75,
	}).(containerBuilder)
