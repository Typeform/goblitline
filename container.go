package goblitline

import "github.com/lann/builder"

type S3Destination struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type containerData struct {
	ImageIdentifier string         `json:"image_identifier"`
	Quality         uint           `json:"quality"`
	S3Destination   *S3Destination `json:"s3_destination,omitempty"`
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

func (b ContainerBuilder) S3Destination(identifier string, destination *S3Destination) ContainerBuilder {
	if destination.Key == "" {
		panic("You need to set S3Destination.Key")
	}
	if destination.Bucket == "" {
		panic("You need to set S3Destination.Bucket")
	}
	return builder.Set(b, "S3Destination", destination).(ContainerBuilder)
}

func (b ContainerBuilder) build() containerData {
	return builder.GetStruct(b).(containerData)
}
