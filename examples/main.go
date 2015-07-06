package main

import (
	. "github.com/gchaincl/goblitline"
)

func main() {
	con := Container.Quality(10)

	fun := Function.
		Name("F2").
		Params("a", "b", "c").
		Save(con)

	job := Job.
		ApplicationID("0WA9Tv18J266Y-hmy7Z-RCg").
		Hash("X").
		Src("src").
		Functions(Function.Name("F1")).
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
		PostbackURL("post back url").
		PostbackHeaders("X-API", "foo").
		PostbackHeaders("X-TKN", "bar").
		WaitForS3(true).
		ContentTypeJson(true).
		V("1.2").
		LongRunning(true)

	job = job.Functions(fun)
	job.Post()
}