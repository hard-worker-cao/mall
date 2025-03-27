package models

import (
	"context"
	"fmt"
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
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err1 := RedisDB.Get(ctx, key).Result()
	if err1 != nil {
		fmt.Println(err1)
		return ""
	}
	if clear {
		err2 := RedisDB.Del(ctx, key).Err()
		if err2 != nil {
			fmt.Println(err2)
			return ""
		}
	}
	return val
}

// 验证cpatcha
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
