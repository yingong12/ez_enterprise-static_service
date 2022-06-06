package request

import "mime/multipart"

type UploadImg struct {
	Imgs  []*multipart.FileHeader `form:"imgs" binding:"required"`
	AppID string                  `form:"app_id" binding:"required"`
}
