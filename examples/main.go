package main

import (
	"os"

	. "github.com/gchaincl/goblitline"
)

func main() {
	con := Container("id").Quality(100)

	fun := Function("annotate").
		Params("text", "Valpo").
		Params("color", "#ffffff").
		Save(con)

	job := Job(os.Getenv("BLITLINE_APP_ID")).
		//Hash("X").
		Functions(fun).
		ImaggaTag(true).
		WillRetryDeplay(1).
		RetryPostback(true).
		ExtendedMetadata(true).
		GetExif(true).
		PassthroughMetadata("foo", "bar").
		PassthroughMetadata("xxx", "yyy").
		IncludeIPTC(true).
		SuppressAutoOrient(true).
		SrcType("source type").
		/*
			PostbackURL("post back url").
			PostbackHeaders("X-API", "foo").
			PostbackHeaders("X-TKN", "bar").
			WaitForS3(true).
			ContentTypeJson(true).
			V("1.2").
			LongRunning(true)
		*/
		Src("http://www.wondermondo.com/Images/SAmerica/Chile/Valparaiso/Valparaiso.jpg")

	job.Post()
}