package main

import (
	"errors"
	"fmt"
	"luas_persegi/repository"
	"time"
)

func main() {
	fmt.Println("Luas Persegi : ", LuasPersegi(4))
}

func LuasPersegi(sisi int) int {
	return sisi * sisi
}

func Register(username, password string) error {
	if username == "" {
		return errors.New("username tidak boleh kosong")
	}
	if password == "" {
		return errors.New("password tidak boleh kosong")
	}

	// ceritanya masukin ke db

	// kalo sukses return nil
	return nil
}

func RegisterToDB(userRepo repository.IUser, username, password string) error {
	if err := userRepo.RegisterWithTImestamp(username, password, time.Now()); err != nil {

	}
}
