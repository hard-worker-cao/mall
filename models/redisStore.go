package models

import (
	"context"
	"time"
)

type RedisStore struct {
}

const (
	CAPTCHA = "captcha"
)

var ctx = context.Background()

// 设置cpatcha，重写Set
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := RedisDB.Set(ctx, key, value, time.Minute).Err()
	return err
}

// 获取captcha的信息重写get
func (r RedisStore) Get(id string, clear bool) (string, error) {
	key := CAPTCHA + id
	val, err := RedisDB.Get(ctx, key).Result()
	if err != nil {
		ErrHandler(err)
		return "", err
	}
	if clear {
		err := RedisDB.Del(ctx, key).Err()
		if err != nil {
			ErrHandler(err)
			return "", err
		}
	}
	return val, nil
}

// 验证cpatcha
func (r RedisStore) VerifyCaptcha(id, answer string, clear bool) bool {
	v, err := RedisStore{}.Get(id, clear)
	if err != nil {
		ErrHandler(err)
	}
	return v == answer
}
