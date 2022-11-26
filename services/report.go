package services

import (
	"context"

	"github.com/golang/glog"
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/helper/common"
	"github.com/truongnh28/environment-be/pkg/converter"
	"github.com/truongnh28/environment-be/repositories"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type ReportService interface {
	Create(ctx context.Context, message *dto.CreateReportRequest) (*dto.CreateReportResponse, error)
	GetByID(ctx context.Context, message *dto.GetReportByIDRequest) (*dto.GetReportByIDResponse, error)
	List(ctx context.Context, message *dto.ListReportsRequest) (*dto.ListReportsResponse, error)
	// UpdateWithMap(ctx context.Context, *dto.UpdateReportRequest) error
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

func (r *reportService) Create(ctx context.Context, message *dto.CreateReportRequest) (*dto.CreateReportResponse, error) {
	record := converter.FromReportDTO(message)
	report, err := r.reportRepo.Create(ctx, record)
	if err != nil {
		return nil, err
	}

	resp := &dto.CreateReportResponse{
		Report: *converter.ToReportDTO(report),
	}
	return resp, nil
}

func (r *reportService) List(ctx context.Context, message *dto.ListReportsRequest) (*dto.ListReportsResponse, error) {

}

// func (r *reportService) UpdateWithMap(ctx context.Context, message *dto.UpdateReportRequest) error {
// 	params := map[string]interface{}
// 	for v := range message.FieldMask {
// 		params[v] =
// 	}
// 	err := r.reportRepo.UpdateWithMap(octx, record.)
// }
