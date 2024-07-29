package enum

const (
	HtmlContentType = "text/html"
)

var MappingTypeUpload = map[string]string{
	"excel": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"csv":   "text/csv",
	"pdf":   "application/pdf",
}

type UploadPlatform string

const (
	S3Uploader UploadPlatform = "s3"
	SelfHosted                = "self-hosted"
)
