package goblitline

import (
	"math/rand"

	"github.com/lann/builder"
)

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

var alphanum = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

func randString(size int) string {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return string(buf)
}

func (b ContainerBuilder) S3Destination(identifier string, destination *S3Destination) ContainerBuilder {
	if destination.Key == "" {
		destination.Key = identifier + "-" + randString(10)
	}
	if destination.Bucket == "" {
		panic("You need to set S3Destination.Bucket")
	}
	return builder.Set(b, "S3Destination", destination).(ContainerBuilder)
}

func (b ContainerBuilder) build() containerData {
	return builder.GetStruct(b).(containerData)
}
