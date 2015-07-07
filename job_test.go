package goblitline

import (
	"bytes"
	"encoding/json"
	"testing"
)

func indent(doc string) (string, error) {
	dest := bytes.NewBuffer(nil)
	err := json.Indent(dest, []byte(doc), "", "\t")
	return dest.String(), err

}

func TestJobToToJson(t *testing.T) {
	builder := Job("app_id").
		Hash("md5").
		Src("the source").
		Functions(Function("f1"), Function("f2")).
		ImaggaTag(true).
		WaitRetryDeplay(10).
		RetryPostback(true).
		ExtendedMetadata(true).
		GetExif(true).
		PassthroughMetadata("foo", "bar").
		PassthroughMetadata("blit", "line").
		IncludeIPTC(true).
		SuppressAutoOrient(true).
		SrcType("screen_shot_url").
		PostbackURL("http://localhost:8000").
		PostbackHeaders("X-TOKEN", "a token").
		PostbackHeaders("X-SRC", "blitline").
		WaitForS3(true).
		ContentTypeJson(true).
		V("1.22").
		LongRunning(true)

	json := builder.ToJson()

	expectedJson :=
		`
	{
		"application_id": "app_id",
		"hash": "md5",
		"src": "the source",
		"functions": [
			{"name": "f1"},
			{"name": "f2"}
		],
		"imagga_tag": true,
		"wait_retry_delay": 10,
		"retry_postback": true,
		"extended_metadata": true,
		"get_exif": true,
		"passthrough_metadata": {
			"blit": "line",
			"foo": "bar"
		},
		"include_iptc": true,
		"supress_auto_orient": true,
		"src_type": "screen_shot_url",
		"postback_url": "http://localhost:8000",
		"postback_headers": {
			"X-SRC": "blitline",
			"X-TOKEN": "a token"
		},
		"wait_for_s3": true,
		"content_type_json": true,
		"v": "1.22",
		"long_running": true
	}`

	json, _ = indent(json)
	expectedJson, _ = indent(expectedJson)

	if json != expectedJson {
		t.Fatalf("`%s`\n`%s`\n are not equals", json, expectedJson)
	}
}