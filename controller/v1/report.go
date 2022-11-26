package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

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

func (s *RepostHandlerImpl) UploadImage(context *gin.Context) {
	var (
		out = &dto.UploadImageResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	file, header, err := context.Request.FormFile("image")
	filename := header.Filename
	path := "./tmp/" + filename
	outImg, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer outImg.Close()
	_, err = io.Copy(outImg, file)
	if err != nil {
		log.Fatal(err)
	}

	const cldUrl = "cloudinary://512616158545567:mClhxuKZ9F-EsP4Kjm_s5qccdvk@dbk0cmzcb"
	cld, err := cloudinary.NewFromURL(cldUrl)
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}
	// upload img to cld
	uploadResult, err := cld.Upload.Upload(
		context,
		path,
		uploader.UploadParams{})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}
	url := uploadResult.URL
	out.Url = url
	os.Remove(path)
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

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

func (s *RepostHandlerImpl) CreateReport(context *gin.Context) {
	var (
		out = &dto.CreateReportResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	var body dto.CreateReportRequest
	if err := json.NewDecoder(context.Request.Body).Decode(&body); err != nil {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	response, code := s.reportService.Create(context, body)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	out.Report = response.Report
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *RepostHandlerImpl) GetReportByID(context *gin.Context) {
	var (
		out = &dto.GetReportByIDResponse{}
	)
	_reportID := context.Request.URL.Query().Get("id")
	fmt.Println("data:, ", _reportID)
	id, _ := strconv.ParseInt(_reportID, 10, 64)
	defer func() {
		context.JSON(200, out)
	}()
	response, code := s.reportService.GetByID(context, dto.GetReportByIDRequest{
		ID: int(id),
	})
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, code)
		return
	}
	out.Report = response.Report
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *RepostHandlerImpl) ListReport(context *gin.Context) {
	var (
		out = &dto.ListReportsResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	var body dto.ListReportsRequest
	if err := json.NewDecoder(context.Request.Body).Decode(&body); err != nil {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	response, code := s.reportService.List(context, body)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	out = &response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *RepostHandlerImpl) GetAllReport(context *gin.Context) {
	var (
		out = &dto.GetAllResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	response, code := s.reportService.GetAll(context)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	out = &response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *RepostHandlerImpl) UpdateReport(context *gin.Context) {
	var (
		out = &dto.UpdateReportResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	var body dto.UpdateReportRequest
	if err := json.NewDecoder(context.Request.Body).Decode(&body); err != nil {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	code := s.reportService.Update(context, body)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}

func (s *RepostHandlerImpl) TopResolver(context *gin.Context) {
	var (
		out = &dto.TopResolverResponse{}
	)
	defer func() {
		context.JSON(200, out)
	}()
	response, code := s.reportService.TopResolver(context)
	if code != common.OK {
		helper.BuildResponseByReturnCode(out, common.Fail, common.SystemError)
		return
	}
	out = &response
	helper.BuildResponseByReturnCode(out, common.Success, common.OK)
}
