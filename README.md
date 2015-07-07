# goblitline
A [Blitline](http://www.blitline.com/) client written in Go

# usage
Goblitline can convert this: ![valpo](http://www.wondermondo.com/Images/SAmerica/Chile/Valparaiso/Valparaiso.jpg)
into this: ![valpo-blitlined](http://s3.amazonaws.com/blitline/2015070716/5643/6i5bMoMpQtr8RU2ysHyAvbg.jpg)
By writing this:
```go
	con := Container("valparaiso")

	f1 := Function("annotate").
		Params("text", "Valpo").
		Params("color", "#ffffff").
		Save(con)

	f2 := Function("annotate").
		Params("text", "github.com/gchaincl/goblitline").
		Params("color", "#000000").
		Params("y", -300).
		Save(con)

	f3 := Function("vignette").Save(con)
	
	res, err := job := Job(os.Getenv("BLITLINE_APP_ID")).
		Functions(f1, f2, f3).
		Src("http://www.wondermondo.com/Images/SAmerica/Chile/Valparaiso/Valparaiso.jpg").
		Post()
	
	results := res.Results
	if results.Error != "" {
		fmt.Printf("Error: %s\n", results.Error)
		return
	}
	
	println("Response:")
	fmt.Printf("JobID: %s\n", results.JobID)
	for _, image := range results.Images {
		fmt.Printf("%s @ %s\n", image.ImageIdentifier, image.S3Url)
	}
```

## TODO
- [x] ~~Function pipelining~~
- [ ] CLI tool
- [ ] Custom s3 buckets
- [x] Testing
- [x] documentation
