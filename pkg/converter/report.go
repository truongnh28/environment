package converter

import (
	"github.com/truongnh28/environment-be/dto"
	"github.com/truongnh28/environment-be/models"
)

func FromReportDTO(message *dto.CreateReportRequest) *models.Reports {
	report := &models.Reports{}
	report.Title = message.Title
	report.Description = message.Description
	report.Priority = message.Priority
	report.Author = message.UserName
	report.City = message.City
	report.District = message.District
	report.Street = message.Street
	report.Ward = message.Ward
	report.Address = *message.Address
	report.Lag = message.Lag
	report.Lng = message.Lng

	return report
}

func ToReportDTO(record *models.Reports) *dto.Report {
	message := &dto.Report{}
	message.ID = record.ID
	message.Title = record.Title
	message.Description = record.Description
	message.CreatedAt = record.CreatedAt
	message.UpdatedAt = record.UpdatedAt
	message.DeletedAt = record.DeletedAt
	message.Status = record.Status
	message.Priority = record.Priority
	message.Author = record.Author
	message.Lag = record.Lag
	message.Lng = record.Lng
	message.Resolver = record.Resolver
	message.City = record.City
	message.District = record.District
	message.Street = record.Street
	message.Ward = record.Ward
	message.Address = record.Address
	message.ImageURL = record.ImageURL

	return message
}

func FromReport(message dto.Report) *models.Reports {
	report := &models.Reports{}
	report.Title = message.Title
	report.Description = message.Description
	report.Priority = message.Priority
	report.Author = message.Author
	report.City = message.City
	report.District = message.District
	report.Street = message.Street
	report.Ward = message.Ward
	report.Address = message.Address
	report.CreatedAt = message.CreatedAt
	report.UpdatedAt = message.UpdatedAt
	report.DeletedAt = message.DeletedAt
	report.ID = message.ID
	report.Lag = message.Lag
	report.Lng = message.Lng
	report.Status = message.Status
	report.Resolver = message.Resolver
	report.ImageURL = message.ImageURL

	return report
}
