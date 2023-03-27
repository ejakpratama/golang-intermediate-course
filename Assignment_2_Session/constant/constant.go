package constant

const (
	SESSION_ID   = "test-id"
	POSTGRES_URL = "postgres://postgre:postgresss@127.0.0.1:5432/postgres?sslmode=disable"
)

var (
	SESSION_AUTH_KEY       = []byte("my-auth-key-very-secret")
	SESSION_ENCRYPTION_KEY = []byte("my-encryption-key-very-secret123")
)
