package main

import (
	"fmt"
	"os"

	. "github.com/Typeform/goblitline"
)

func main() {
	con1 := Container("valparaiso-1").S3Destination("my-key", "my-bucket")
	con2 := Container("valparaiso-3")

	f1 := Function("annotate").
		Params("text", "Valpo").
		Params("color", "#ffffff").
		Save(con1)

	f2 := Function("annotate").
		Params("text", "github.com/Typeform/goblitline").
		Params("color", "#000000").
		Params("y", -300).
		Save(con2)

	f3 := Function("vignette").Save(con1)

	job := Job(os.Getenv("BLITLINE_APP_ID")).
		Functions(f1, f2, f3).
		/*
			Hash("md5").
			ImaggaTag(true).
			WaitRetryDeplay(1).
			RetryPostback(true).
			ExtendedMetadata(true).
			GetExif(true).
			PassthroughMetadata("foo", "bar").
			PassthroughMetadata("xxx", "yyy").
			IncludeIPTC(true).
			SuppressAutoOrient(true).
			SrcType("source type").
			PostbackURL("post back url").
			PostbackHeaders("X-API", "foo").
			PostbackHeaders("X-TKN", "bar").
			WaitForS3(true).
			ContentTypeJson(true).
			V("1.2").
			LongRunning(true)
		*/
		Src("http://www.wondermondo.com/Images/SAmerica/Chile/Valparaiso/Valparaiso.jpg")

	fmt.Printf("Posting Json to Blitline:\n`%s`\n--\n", job.ToJson().String())

	response, err := job.Post()
	if err != nil {
		panic(err)
	}

	results := response.Results
	if results.Error != "" {
		fmt.Printf("Error: %s\n", results.Error)
		return
	}

	println("Response:")
	fmt.Printf("JobID: %s\n", results.JobID)
	for _, image := range results.Images {
		fmt.Printf("%s @ %s\n", image.ImageIdentifier, image.S3Url)
	}

}
