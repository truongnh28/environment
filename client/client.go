package client

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(ctx context.Context, cld *cloudinary.Cloudinary, imagePath string) (*uploader.UploadResult, error) {
	resp, err := cld.Upload.Upload(ctx, imagePath, uploader.UploadParams{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
