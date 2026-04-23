package interfaces

import "context"

type IEmailExternal interface {
	SendEmail() error
}

type IEmailRepo interface {
	GetTemplate(ctx context.Context, templateName string) (*models.EmailTemplate, error)
	InsertNotificationHistory(ctx context.Context, templateName string, notif *models.NotificationHistory) error
}
