package services

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"spotify/cache"
	"spotify/config"
	"spotify/dto"
	"spotify/helper/common"
	"spotify/models"
	"spotify/repositories"
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
	Update(ctx context.Context, username string, req dto.UpdateAccountRequest) common.SubReturnCode
	Create(ctx context.Context, message dto.CreateReportRequest) dto.CreateReportResponse
	FindByUserName(ctx context.Context, username string) (dto.Account, common.SubReturnCode)
}


type SongService interface {
	GetAllSong() ([]dto.Song, common.SubReturnCode)
}

func NewSongService(songRepo repositories.SongRepository) SongService {
	return &songServiceImpl{
		songRepo: songRepo,
	}
}

type songServiceImpl struct {
	songRepo repositories.SongRepository
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
