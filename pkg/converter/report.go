package converters

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

	return report
}
