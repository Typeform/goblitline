package goblitline

func Job(AppId string) JobBuilder {
	return JobBuilder{}.ApplicationID(AppId)
}

func Function(name string) FunctionBuilder {
	return FunctionBuilder{}.Name(name)
}

func Container(imageId string) ContainerBuilder {
	return ContainerBuilder{}.
		ImageIdentifier(imageId).
		Quality(75)
}
