// Goblitline is a golang client for the Blitline API.
//
// It provide a DSL for building Jobs, Functions and Containers (function saving).
//
// The DSL try to mimic Blitline API, so if you are familiar with it, use goblitline should be easy enough.
// For more information about blitline api, please refer to http://www.blitline.com/docs/api.
package goblitline

// Returns a JobBuilder
func Job(AppId string) JobBuilder {
	return JobBuilder{}.ApplicationID(AppId)
}

// Returns a FunctionBuilder
func Function(name string) FunctionBuilder {
	return FunctionBuilder{}.Name(name)
}

// Returns a ContainerBuilder
func Container(imageId string) ContainerBuilder {
	return ContainerBuilder{}.
		ImageIdentifier(imageId).
		Quality(75)
}
