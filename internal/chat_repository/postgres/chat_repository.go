package postgres

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	chatrepository "github.com/westcrime/chat-server/internal/chat_repository"
	"github.com/westcrime/chat-server/internal/model"
)

type chatRepository struct {
	pool *pgxpool.Pool
}

func NewChatRepository(pool *pgxpool.Pool) chatrepository.ChatRepository {
	return &chatRepository{pool: pool}
}

func (ur *chatRepository) CreateChat(ctx context.Context, createChat *model.CreateChat) (error, int64) {
	builderInsert := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(createChat.Name).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err, -1
	}

	var id int64
	err = ur.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return err, -1
	}

	builderChatsUsersInsert := sq.Insert("chats_users").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "username")
	for _, username := range createChat.Usernames {
		builderChatsUsersInsert = builderChatsUsersInsert.Values(id, username)
	}

	query, args, err = builderChatsUsersInsert.ToSql()
	if err != nil {
		return err, -1
	}

	_, err = ur.pool.Exec(ctx, query, args...)
	if err != nil {
		return err, -1
	}

	return nil, id
}

func (ur *chatRepository) DeleteChat(ctx context.Context, chat_id int64) error {
	builderInsert := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": chat_id})

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	res, err := ur.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %d\n", res.RowsAffected())
	fmt.Printf("Command: %s\n", res.String())

	builderInsert = sq.Delete("chats_users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": chat_id})

	res, err = ur.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %d\n", res.RowsAffected())
	fmt.Printf("Command: %s\n", res.String())

	return nil
}

func (ur *chatRepository) SendMessage(ctx context.Context, message *model.Message) error {
	builderInsert := sq.Insert("messages").
		PlaceholderFormat(sq.Dollar).
		Columns("from_user", "text", "chat_id", "timestamp").
		Values(message.From, message.Text, message.ChatId, message.Timestamp)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	res, err := ur.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %d\n", res.RowsAffected())
	fmt.Printf("Command: %s\n", res.String())
	return nil
}
