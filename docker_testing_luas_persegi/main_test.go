package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestLuasPersegi(t *testing.T) {
	t.Run("case sisi 4", func(t *testing.T) {
		luas := LuasPersegi(4)
		require.Equal(t, 16, luas)
	})
	t.Run("case sisi 6", func(t *testing.T) {
		luas := LuasPersegi(6)
		require.Equal(t, 36, luas)
	})
}

func TestRegister(t *testing.T) {
	t.Run("case username kosong", func(t *testing.T) {
		err := Register("", "password")
		require.Error(t, err)
	})
	t.Run("case password kosong", func(t *testing.T) {
		err := Register("username", "")
		require.Error(t, err)
	})
	t.Run("case sukses", func(t *testing.T) {
		err := Register("username", "password")
		require.NoError(t, err)
	})
}

// func TestLuasSamaDengan64(t *testing.T) {
// 	assert.Equals(t, LuasPersegi(8), 64, "Luas Persegi Tidak Sama Dengan 64")
// }

func BenchmarkLuasPersegi(b *testing.B) {
	b.Log("Test Benchmark Luas Persegi")
	for i := 1; i < b.N; i++ {
		b.Log("Test Benchmark Luas Persegi ke : ", i)
		LuasPersegi(100)
	}
}

func TestRegisterToDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Case Error", func(t *Testing.T) {
		userRepo := mock_repository.NewMockUser(ctrl)
		userRepo.Expect().Register("username", "password").Return(errors.New("Error DB Mati"))

		err := RegisterToDB(userRepo, "username", "password")
		require.Error(t, err)
	})

}

func TestRegisterToDBWithTimestamp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

}
