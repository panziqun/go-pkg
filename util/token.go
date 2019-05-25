package util

import (
	"errors"
	"fmt"
	"strconv"
)

const TOKEN_LENGTH = 8

func ParseToken(token string) (id int, err error) {
	if token == "" {
		return 0, errors.New("empty token！")
	}
	if len(token) < TOKEN_LENGTH+1 {
		return 0, errors.New("token invalid!")
	}

	tb := []byte(token)
	id, _ = strconv.Atoi(string(tb[TOKEN_LENGTH:]))
	return id, nil
}

func GenerateToken(id int) string {
	str := RandString(TOKEN_LENGTH)
	return str + fmt.Sprintf("%d", id)
}
