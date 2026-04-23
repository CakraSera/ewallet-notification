package repository

import (
	"context"
	"ewallet-notification-service/internal/models"

	"ewallet-notification/internal/models"

	"gorm.io/gorm"
)

type EmailRepo struct {
	DB *gorm.DB
}

func (r *EmailRepo) GetTemplate(ctx context.Context, templateName string) (*models.EmailTemplate, error) {
	var resp models.NotificationTemplate
	err := r.DB.Where("template_name = ?", templateName).Last(&resp).Error
	return resp, err
}

func (r *EmailRepo) InsertNotificationHistory(ctx context.Context, templateName string, notif *models.NotificationHistory) error {
	return r.DB.Create(notif).Error
}
