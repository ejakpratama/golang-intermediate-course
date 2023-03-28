package main

import (
	"fmt"
	"net/http"

	"ldap_auth/config"
	"ldap_auth/repository"
	"ldap_auth/router"
	"ldap_auth/service"

	"github.com/go-ldap/ldap"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

const (
	port    = "9000"
	baseURL = "0.0.0.0:" + port
)

func main() {
	log.SetReportCaller(true)

	// init ldap connection
	ldapConn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", config.LdapServer, config.LdapPort))
	if err != nil {
		panic(err)
	}
	defer ldapConn.Close()
	// bind to ldap server
	if err = ldapConn.Bind(config.LdapBindDN, config.LdapPassword); err != nil {
		panic(err)
	}

	ldapRepo := repository.NewLDAPRepo(ldapConn)
	loginService := service.NewLoginService(ldapRepo)

	r := mux.NewRouter()
	routerHandler := router.NewRouterHandler(loginService)
	r.HandleFunc("/", routerHandler.ServeHTML)
	r.HandleFunc("/login", routerHandler.HandleLogin)

	log.Infoln("Listening at ", baseURL)
	log.Fatal(http.ListenAndServe(baseURL, r))
}
