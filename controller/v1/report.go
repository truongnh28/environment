package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/truongnh28/environment-be/client"
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/helper"
	"github.com/truongnh28/environment-be/helper/common"
	"github.com/truongnh28/environment-be/services"
)

type RepostHandlerImpl struct {
	cloudinaryService client.CloudinaryAPI
	reportService     services.ReportService
}

func NewReportHandler(reportService services.ReportService) *RepostHandlerImpl {
	return &RepostHandlerImpl{
		reportService: __reportService,
	}
}

//func (s *RepostHandlerImpl) MapResponseList(context *gin.Context) {
//	var (
//		out = &dto.MapResponse{}
//	)
//	defer func() {
//		context.JSON(200, out)
//	}()
//	file, header, err := context.Request.FormFile("image")
//	filename := header.Filename
//	outImg, err := os.Create("./tmp/" + filename + ".png")
//	fmt.Println(header.Filename)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer outImg.Close()
//	_, err = io.Copy(outImg, file)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//response, code := s.reportService.GetAllSong()
//	if code != common.OK {
//		helper.BuildResponseByReturnCode(out, common.Fail, code)
//		return
//	}
//	out.Reports = response
//	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
//}

func (s *RepostHandlerImpl) MapResponseList(c *gin.Context) {
	var (
		out = &dto.MapResponse{}
		ctx = context.Background()
	)
	defer func() {
		c.JSON(200, out)
	}()
	response, code := s.reportService.MapReportList(ctx)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.MapResps = response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
