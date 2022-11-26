package services

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/truongnh28/environment-be/cache"
	"github.com/truongnh28/environment-be/config"
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/helper/common"
	"github.com/truongnh28/environment-be/models"
	"github.com/truongnh28/environment-be/converters"
	"github.com/truongnh28/environment-be/repositories"
	"time"
)


Create(ctx context.Context, record *models.Reports) (*models.Reports, error)
UpdateWithMap(
	octx context.Context,
	record *models.Reports,
	params map[string]interface{},
) error
GetByID(ctx context.Context, id int) (*models.Reports, error)
List(
	octx context.Context,
	size int,
	page int,
	filter *models.ReportFilter,
) ([]*models.Reports, error)
Delete(octx context.Context, record *models.Reports) error
CountWithFilter(
	octx context.Context,
	filter *models.ReportFilter,
) (int64, error)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type ReportService interface {
	Create(ctx context.Context, message dto.CreateReportRequest) dto.CreateReportResponse
	GetByID(ctx context.Context, message dto.GetReportByIDRequest) dto.GetReportByIDResponse
	List(ctx context.Context, message dto.ListReportsRequest) dto.ListReportsResponse
}

func NewReportService(reportRepo repositories.ReportRepository) ReportService {
	return &reportService{
		songRepo: reportRepo,
	}
}

type reportService struct {
	reportRepo repositories.ReportRepository
}


func (s *songServiceImpl) GetAllSong() ([]dto.Song, common.SubReturnCode) {
	var resp = make([]dto.Song, 0)
	songs, err := s.songRepo.GetAllSong()
	if err != nil {
		glog.Infoln("GetAllSong service err: ", err)
		return resp, common.SystemError
	}
	for _, val := range songs {
		resp = append(resp, dto.Song{
			SongID:      val.SongID,
			Name:        val.Name,
			AlbumID:     val.AlbumID,
			ArtistID:    val.ArtistID,
			Lyrics:      val.Lyrics,
			Length:      val.Length,
			TrackNumber: val.TrackNumber,
			CreateAt:    val.CreateAt,
			UploadAt:    val.UploadAt,
			YoutubeLink: val.YoutubeLink,
		})
	}
	return resp, common.OK
}

func (r *reportService) Create(ctx context.Context, message dto.CreateReportRequest) dto.CreateReportResponse {
	record := converters.From
	report, err := r.reportRepo.Create(ctx, )
}
