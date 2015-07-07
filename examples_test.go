package goblitline

import "fmt"

func ExampleJob() {
	res, _ := Job("MyAppKey").
		Src("http://golang.org/doc/gopher/frontpage.png").
		Functions(
		Function("sepia_tone").
			Save(Container("gopher")),
	).Post()

	for _, image := range res.Results.Images {
		fmt.Printf("%s @ %s\n", image.ImageIdentifier, image.S3Url)
	}
}
