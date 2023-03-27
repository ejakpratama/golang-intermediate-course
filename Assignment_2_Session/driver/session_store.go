package driver

import (
	"assignment_2_session/constant"
	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"log"
)

func NewPostgresStore() *pgstore.PGStore {
	store, err := pgstore.NewPGStore(constant.POSTGRES_URL, constant.SESSION_AUTH_KEY, constant.SESSION_ENCRYPTION_KEY)
	if err != nil {
		log.Fatalln("ERROR", err)
	}

	return store
}

func NewCookieStore() *sessions.CookieStore {
	store := sessions.NewCookieStore(constant.SESSION_AUTH_KEY, constant.SESSION_ENCRYPTION_KEY)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true

	return store
}
