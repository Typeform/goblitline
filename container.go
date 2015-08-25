package goblitline

import "github.com/lann/builder"

type S3Destination struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type containerData struct {
	ImageIdentifier string         `json:"image_identifier"`
	Quality         uint           `json:"quality"`
	Extension       string         `json:"extension"`
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

func (b ContainerBuilder) Extension(ext string) ContainerBuilder {
	return builder.Set(b, "Extension", ext).(ContainerBuilder)
}

func (b ContainerBuilder) S3Destination(key string, bucket string) ContainerBuilder {
	destination := &S3Destination{Key: key, Bucket: bucket}
	return builder.Set(b, "S3Destination", destination).(ContainerBuilder)
}

func (b ContainerBuilder) build() containerData {
	return builder.GetStruct(b).(containerData)
}
