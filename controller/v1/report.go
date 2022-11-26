package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/truongnh28/environment-be/client"
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/helper"
	"github.com/truongnh28/environment-be/helper/common"
	"io"
	"log"
	"os"
)

type RepostHandlerImpl struct {
	cloudinaryService client.CloudinaryAPI
	reportService     services.reportService
}

func NewReportHandler(cloudinaryService client.CloudinaryAPI) *RepostHandlerImpl {
	return &RepostHandlerImpl{
		cloudinaryService: cloudinaryService,
	}
}

func (s *RepostHandlerImpl) PostReport(context *gin.Context) {
	var (
		out = &dto.ReportResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	file, header, err := context.Request.FormFile("image")
	filename := header.Filename
	outImg, err := os.Create("./tmp/" + filename + ".png")
	fmt.Println(header.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outImg.Close()
	_, err = io.Copy(outImg, file)
	if err != nil {
		log.Fatal(err)
	}

	response, code := s.reportService.GetAllSong()
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Reports = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
