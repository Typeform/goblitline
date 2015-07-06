package goblitline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lann/builder"
)

type job struct {
	ApplicationID       string            `json:"application_id"`
	Hash                string            `json:"hash,omitempty"`
	Src                 string            `json:"src"`
	Functions           []function        `json:"functions"`
	ImaggaTag           bool              `json:"imagga_tag,omitempty"`
	WillRetryDeplay     uint              `json:"wait_retry_delay,omitempty"`
	RetryPostback       bool              `json:"retry_postback,omitempty"`
	ExtendedMetadata    bool              `json:"extended_metadata,omitempty"`
	GetExif             bool              `json:"get_exif,omitempty"`
	PassthroughMetadata map[string]string `json:"passthrough_metadata,omitempty"`
	IncludeIPTC         bool              `json:"include_iptc,omitempty"`
	SupressAutoOrient   bool              `json:"supress_auto_orient,omitempty"`
	SrcType             string            `json:"src_type,omitempty"`
	PostbackURL         string            `json:"postback_url,omitempty"`
	PostbackHeaders     map[string]string `json:"postback_headers,omitempty"`
	WaitForS3           bool              `json:"wait_for_s3,omitempty"`
	ContentTypeJson     bool              `json:"content_type_json,omitempty"`
	V                   string            `json:"V,omitempty"`
	LongRunning         bool              `json:"long_running,omitempty"`
}

type jobBuilder builder.Builder

func (b jobBuilder) ApplicationID(id string) jobBuilder {
	return builder.Set(b, "ApplicationID", id).(jobBuilder)
}

func (b jobBuilder) Hash(hash string) jobBuilder {
	return builder.Set(b, "Hash", hash).(jobBuilder)
}

func (b jobBuilder) Src(src string) jobBuilder {
	return builder.Set(b, "Src", src).(jobBuilder)
}

func (b jobBuilder) Functions(functions ...functionBuilder) jobBuilder {
	for _, function := range functions {
		f := function.build()
		b = builder.Append(b, "Functions", f).(jobBuilder)
	}
	return b
}

func (b jobBuilder) ImaggaTag(v bool) jobBuilder {
	return builder.Set(b, "ImaggaTag", v).(jobBuilder)
}

func (b jobBuilder) WillRetryDeplay(delay uint) jobBuilder {
	return builder.Set(b, "WillRetryDeplay", delay).(jobBuilder)
}

func (b jobBuilder) RetryPostback(v bool) jobBuilder {
	return builder.Set(b, "RetryPostback", v).(jobBuilder)
}

func (b jobBuilder) ExtendedMetadata(v bool) jobBuilder {
	return builder.Set(b, "ExtendedMetadata", v).(jobBuilder)
}

func (b jobBuilder) GetExif(v bool) jobBuilder {
	return builder.Set(b, "GetExif", v).(jobBuilder)
}

func (b jobBuilder) PassthroughMetadata(k, v string) jobBuilder {
	var hash map[string]string
	meta, ok := builder.Get(b, "PassthroughMetadata")
	if ok {
		hash = meta.(map[string]string)
	} else {
		hash = make(map[string]string)
	}
	hash[k] = v
	return builder.Set(b, "PassthroughMetadata", hash).(jobBuilder)
}

func (b jobBuilder) IncludeIPTC(v bool) jobBuilder {
	return builder.Set(b, "IncludeIPTC", v).(jobBuilder)
}

func (b jobBuilder) SuppressAutoOrient(v bool) jobBuilder {
	return builder.Set(b, "SupressAutoOrient", v).(jobBuilder)
}

func (b jobBuilder) SrcType(src_type string) jobBuilder {
	return builder.Set(b, "SrcType", src_type).(jobBuilder)
}

func (b jobBuilder) PostbackURL(src_type string) jobBuilder {
	return builder.Set(b, "PostbackURL", src_type).(jobBuilder)
}

func (b jobBuilder) WaitForS3(v bool) jobBuilder {
	return builder.Set(b, "WaitForS3", v).(jobBuilder)
}

func (b jobBuilder) ContentTypeJson(v bool) jobBuilder {
	return builder.Set(b, "ContentTypeJson", v).(jobBuilder)
}

func (b jobBuilder) PostbackHeaders(k, v string) jobBuilder {
	var hash map[string]string
	meta, ok := builder.Get(b, "PostbackHeaders")
	if ok {
		hash = meta.(map[string]string)
	} else {
		hash = make(map[string]string)
	}
	hash[k] = v
	return builder.Set(b, "PostbackHeaders", hash).(jobBuilder)
}

func (b jobBuilder) V(v string) jobBuilder {
	return builder.Set(b, "V", v).(jobBuilder)
}

func (b jobBuilder) LongRunning(v bool) jobBuilder {
	return builder.Set(b, "LongRunning", v).(jobBuilder)
}

func (b jobBuilder) build() job {
	return builder.GetStruct(b).(job)
}

func (b jobBuilder) Post() error {
	built := b.build()
	// TODO: validate

	job, err := json.Marshal(built)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://api.blitline.com/job", "application/json", bytes.NewReader(job))

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	response := Response{}

	json.NewDecoder(resp.Body).Decode(&response)
	fmt.Printf("%#v\n", response)

	return nil
}

var Job = builder.Register(jobBuilder{}, job{}).(jobBuilder)