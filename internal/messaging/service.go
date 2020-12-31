package messaging

import (
	"context"
	"main/db"
	"main/internal/dto"
	"main/internal/entity"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IMessagingService interface {
	SendMessage(message dto.CreateMessageDTO) error
}

type MessagingService struct {
	pool *pgxpool.Pool
}

func New() MessagingService {
	return MessagingService{db.Connection()}
}

func (ms MessagingService) SendMessage(message dto.CreateMessageDTO) error {
	_, err := ms.pool.Exec(context.Background(), "insert into messages (sender_id, receiver_id, subject, body, created_at) values ($1, $2, $3, $4, now())", message.SenderID, message.ReceiverID, message.Subject, message.Body)
	return err
}

func (ms MessagingService) RetrieveUserMessages(userID int) ([]*entity.Message, error) {
	var m []*entity.Message
	err := pgxscan.Select(context.Background(), ms.pool, &m, "select sender_id, receiver_id, subject, body, created_at from message where receiver_id = $1", userID)
	return m, err
}

func (ms MessagingService) RetrieveUserSentMessages(userID int) ([]*entity.Message, error) {
	var m []*entity.Message
	err := pgxscan.Select(context.Background(), ms.pool, &m, "select sender_id, receiver_id, subject, body, created_at from message where sender_id = $1", userID)
	return m, err
}
