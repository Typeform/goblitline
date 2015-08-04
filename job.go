package goblitline

import (
	"bytes"
	"encoding/json"

	"github.com/lann/builder"
)

type jobData struct {
	ApplicationID       string            `json:"application_id"`
	Hash                string            `json:"hash,omitempty"`
	Src                 string            `json:"src"`
	Functions           []functionData    `json:"functions"`
	ImaggaTag           bool              `json:"imagga_tag,omitempty"`
	WaitRetryDeplay     uint              `json:"wait_retry_delay,omitempty"`
	RetryPostback       bool              `json:"retry_postback,omitempty"`
	ExtendedMetadata    bool              `json:"extended_metadata,omitempty"`
	GetExif             bool              `json:"get_exif,omitempty"`
	PassthroughMetadata map[string]string `json:"passthrough_metadata,omitempty"`
	IncludeIPTC         bool              `json:"include_iptc,omitempty"`
	SuppressAutoOrient  bool              `json:"supress_auto_orient,omitempty"`
	SrcType             string            `json:"src_type,omitempty"`
	PostbackURL         string            `json:"postback_url,omitempty"`
	PostbackHeaders     map[string]string `json:"postback_headers,omitempty"`
	WaitForS3           bool              `json:"wait_for_s3,omitempty"`
	ContentTypeJson     bool              `json:"content_type_json,omitempty"`
	V                   string            `json:"v,omitempty"`
	LongRunning         bool              `json:"long_running,omitempty"`
}

type JobBuilder builder.Builder

func init() {
	builder.Register(JobBuilder{}, jobData{})
}

func (b JobBuilder) ApplicationID(id string) JobBuilder {
	return builder.Set(b, "ApplicationID", id).(JobBuilder)
}

func (b JobBuilder) Hash(hash string) JobBuilder {
	return builder.Set(b, "Hash", hash).(JobBuilder)
}

func (b JobBuilder) Src(src string) JobBuilder {
	return builder.Set(b, "Src", src).(JobBuilder)
}

func (b JobBuilder) Functions(functions ...FunctionBuilder) JobBuilder {
	for _, function := range functions {
		f := function.build()
		b = builder.Append(b, "Functions", f).(JobBuilder)
	}
	return b
}

func (b JobBuilder) ImaggaTag(v bool) JobBuilder {
	return builder.Set(b, "ImaggaTag", v).(JobBuilder)
}

func (b JobBuilder) WaitRetryDeplay(delay uint) JobBuilder {
	return builder.Set(b, "WaitRetryDeplay", delay).(JobBuilder)
}

func (b JobBuilder) RetryPostback(v bool) JobBuilder {
	return builder.Set(b, "RetryPostback", v).(JobBuilder)
}

func (b JobBuilder) ExtendedMetadata(v bool) JobBuilder {
	return builder.Set(b, "ExtendedMetadata", v).(JobBuilder)
}

func (b JobBuilder) GetExif(v bool) JobBuilder {
	return builder.Set(b, "GetExif", v).(JobBuilder)
}

func (b JobBuilder) PassthroughMetadata(k, v string) JobBuilder {
	var hash map[string]string
	meta, ok := builder.Get(b, "PassthroughMetadata")
	if ok {
		hash = meta.(map[string]string)
	} else {
		hash = make(map[string]string)
	}
	hash[k] = v
	return builder.Set(b, "PassthroughMetadata", hash).(JobBuilder)
}

func (b JobBuilder) IncludeIPTC(v bool) JobBuilder {
	return builder.Set(b, "IncludeIPTC", v).(JobBuilder)
}

func (b JobBuilder) SuppressAutoOrient(v bool) JobBuilder {
	return builder.Set(b, "SuppressAutoOrient", v).(JobBuilder)
}

func (b JobBuilder) SrcType(src_type string) JobBuilder {
	return builder.Set(b, "SrcType", src_type).(JobBuilder)
}

func (b JobBuilder) PostbackURL(src_type string) JobBuilder {
	return builder.Set(b, "PostbackURL", src_type).(JobBuilder)
}

func (b JobBuilder) WaitForS3(v bool) JobBuilder {
	return builder.Set(b, "WaitForS3", v).(JobBuilder)
}

func (b JobBuilder) ContentTypeJson(v bool) JobBuilder {
	return builder.Set(b, "ContentTypeJson", v).(JobBuilder)
}

func (b JobBuilder) PostbackHeaders(k, v string) JobBuilder {
	var hash map[string]string
	meta, ok := builder.Get(b, "PostbackHeaders")
	if ok {
		hash = meta.(map[string]string)
	} else {
		hash = make(map[string]string)
	}
	hash[k] = v
	return builder.Set(b, "PostbackHeaders", hash).(JobBuilder)
}

func (b JobBuilder) V(v string) JobBuilder {
	return builder.Set(b, "V", v).(JobBuilder)
}

func (b JobBuilder) LongRunning(v bool) JobBuilder {
	return builder.Set(b, "LongRunning", v).(JobBuilder)
}

func (b JobBuilder) build() jobData {
	return builder.GetStruct(b).(jobData)
}

func (b JobBuilder) ToJson() *bytes.Buffer {
	doc, _ := json.Marshal(b.build())
	return bytes.NewBuffer(doc)
}

func (b JobBuilder) Post() (*Response, error) {
	body := b.ToJson()
	return Post(body)
}
