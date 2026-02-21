package db

import (
    "fmt"
    "github.com/jackc/pgx/v5"
    "wod-backend/config"
    "context"
    "time"
)

func getUrl() string {
    config := config.GetPSConfig()
    return fmt.Sprintf("postgres://%s:%s@%s:5432/%s", config.User, config.Password, config.Host, config.DbName)
}

func Now() (time.Time, error) {
	conn, err := pgx.Connect(context.Background(), getUrl())
	result := time.Time{}
	if err != nil {
		return result, err
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(), "SELECT NOW()").Scan(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}