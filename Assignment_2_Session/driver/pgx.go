package driver

import (
	"assignment_2_session/constant"
	"context"

	"github.com/jackc/pgx/v4"
)

func NewPostgresConn(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, constant.POSTGRES_URL)
}
