package goblitline

type Image struct {
	ImageIdentifier string `json:"image_identifier"`
	S3Url           string `json:"s3_url"`
}

type Results struct {
	Error  string  `json:"error"`
	JobID  string  `json:"job_id"`
	Images []Image `json:"images"`
}

type Response struct {
	Results Results `json:"results"`
}
