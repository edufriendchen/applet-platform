package cache

import (
	"context"
	"errors"
	"fmt"
	"log"

	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
)

func (r *RedisCacheStore) GetRedis() *redis.Client {
	return r.conn
}

// GetBytes ...
func (r *RedisCacheStore) GetBytes(key string) ([]byte, error) {
	b, err := r.conn.Get(key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		log.Printf("error in retrieving redis data [key: %s] %s", key, err.Error())
		return nil, err
	}

	return b, nil
}

// SetBytesWithTTL ...
func (r *RedisCacheStore) SetBytesWithTTL(key string, data []byte, ttl time.Duration) error {
	err := r.conn.Set(key, data, ttl).Err()
	if err != nil {
		log.Printf("error storing data to redis [key: %s]: %s", key, err.Error())
	}

	return nil
}

func (r *RedisCacheStore) SetIntWithTTL(key string, data int, ttl time.Duration) error {
	err := r.conn.Set(key, data, ttl).Err()
	if err != nil {
		log.Printf("error storing data to redis [key: %s]: %s", key, err.Error())
	}

	return nil
}

func (r *RedisCacheStore) SetUint64WithTTL(key string, data uint64, ttl time.Duration) error {
	err := r.conn.Set(key, data, ttl).Err()
	if err != nil {
		log.Printf("error storing data to redis [key: %s]: %s", key, err.Error())
	}

	return nil
}

func (r *RedisCacheStore) SADD(key string, args ...interface{}) error {
	err := r.conn.SAdd(key, args...).Err()
	if err != nil {
		log.Printf("error storing data to redis [key: %s] %s", key, err.Error())
		return err
	}

	return nil
}

func (r *RedisCacheStore) GetInt(key string) (int, error) {
	i, err := r.conn.Get(key).Int()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}

		log.Printf("error in retrieving redis data [key: %s] %s", key, err.Error())
		return 0, err
	}

	return i, nil
}

func (r *RedisCacheStore) GetUint64(key string) (uint64, error) {
	i, err := r.conn.Get(key).Uint64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}

		log.Printf("error in retrieving redis data [key: %s] %s", key, err.Error())
		return 0, err
	}

	return i, nil
}

func (r *RedisCacheStore) GetSMembersUint64(key string) ([]uint64, error) {
	listMember, err := r.conn.SMembers(key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return []uint64{}, nil
		}

		log.Printf("error in retrieving redis data [key: %s] %s", key, err.Error())
		return []uint64{}, err
	}

	resultList := make([]uint64, len(listMember))

	for i, val := range listMember {
		resultList[i], _ = strconv.ParseUint(val, 10, 64)
	}

	return resultList, nil
}

// Del ...
func (r *RedisCacheStore) Del(key string) error {
	err := r.conn.Del(key).Err()
	if err != nil {
		log.Printf("error in deleting keys : %s, %s", key, err.Error())
	}
	return nil
}

func (r *RedisCacheStore) TTL(key string) (time.Duration, error) {
	i, err := r.conn.TTL(key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}

		log.Printf("error in retrieving redis data [key: %s] %s", key, err.Error())
		return 0, err
	}

	return i, nil
}

// HSet
func (r *RedisCacheStore) HSet(key, field, value string) error {
	err := r.conn.HSet(key, field, value).Err()
	if err != nil {
		log.Printf("error storing data to redis [key: %s]: %s", key, err.Error())
	}

	return nil
}

// HGet
func (r *RedisCacheStore) HGet(key, field string) (string, error) {
	value, err := r.conn.HGet(key, field).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}

		log.Printf("error in retrieving redis data [key: %s] %s", key, err.Error())
		return "", err
	}

	return value, nil
}

// HDel
func (r *RedisCacheStore) HDel(key, field string) error {
	err := r.conn.HDel(key, field).Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}

		log.Printf("error in retrieving redis data [key: %s] %s", key, err.Error())
		return err
	}

	return nil
}

// HmSet
func (r *RedisCacheStore) HmSet(ctx context.Context, key string, values map[string]interface{}, ttl time.Duration) error {
	err := r.conn.HMSet(key, values).Err()
	if err != nil {
		errMsg := fmt.Sprintf("[redis][hmSet] error in hmset [key: %s] [field: %v] ", key, values)
		log.Printf(errMsg, err)
		return err
	}
	if ttl > 0 {
		err = r.conn.Do("EXPIRE", key, int(ttl)).Err()
		if err != nil {
			errMsg := fmt.Sprintf("[redis][hmSet] error in set EXPIRE [key: %s] [field: %v] ", key, values)
			log.Printf(errMsg, err)
			return err
		}
	}
	return nil
}

// Expire
func (r *RedisCacheStore) Expire(key string, duration time.Duration) error {
	err := r.conn.Expire(key, duration).Err()
	if err != nil {
		log.Printf("error set Expire key to redis [key: %s]: %s", key, err.Error())
	}
	return nil
}

// HmGet
func (r *RedisCacheStore) HmGet(ctx context.Context, key string, fields []string) (map[string]interface{}, error) {
	values, err := r.conn.HMGet(key, fields...).Result()
	if err != nil && err != redis.Nil {
		errMsg := fmt.Sprintf("[redis][hmGet] error in retrieving redis data [key: %s] [field: %v] ", key, values)
		log.Printf(errMsg, err)
		return nil, err
	}

	respMap := make(map[string]interface{}, 0)
	for findex, key := range fields {
		for vindex, value := range values {
			if findex == vindex {
				if value != nil {
					respMap[key] = values[vindex].(interface{})
				}
				break
			}
		}
	}
	return respMap, nil
}

// GetBytes ...
func (r *RedisCacheStore) GetString(key string) (string, error) {
	value, err := r.conn.Get(key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}

		log.Printf("error in retrieving redis data [key: %s] %s", key, err.Error())
		return "", err
	}

	return value, nil
}

// SetBytesWithTTL ...
func (r *RedisCacheStore) SetStringWithTTL(key string, text string, ttl time.Duration) error {
	err := r.conn.Set(key, text, ttl).Err()
	if err != nil {
		log.Printf("error storing data to redis [key: %s]: %s", key, err.Error())
		return err
	}

	return nil
}
