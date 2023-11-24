package cache

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
)

const defaultTTL = 400

// RedisCacheStore is cache store implementation on redis
type RedisCacheStore struct {
	conn *redis.Client
}

// CacheStore is cache definition
type CacheStore interface {
	GetRedis() *redis.Client
	SetBytesWithTTL(key string, data []byte, ttl time.Duration) error
	GetBytes(key string) ([]byte, error)
	Del(key string) error
	SetIntWithTTL(key string, data int, ttl time.Duration) error
	SetUint64WithTTL(key string, data uint64, ttl time.Duration) error
	SADD(key string, args ...interface{}) error
	GetInt(key string) (int, error)
	GetUint64(key string) (uint64, error)
	GetSMembersUint64(key string) ([]uint64, error)
	TTL(key string) (time.Duration, error)
	HSet(key, field string, data string) error
	HGet(key, field string) (string, error)
	HDel(key, field string) error
	HmSet(ctx context.Context, key string, values map[string]interface{}, ttl time.Duration) error
	HmGet(ctx context.Context, key string, fields []string) (map[string]interface{}, error)
	Expire(key string, duration time.Duration) error
	SetStringWithTTL(key string, text string, ttl time.Duration) error
	GetString(key string) (string, error)
}

// NewCacheStore initialize connection for redis and store it into redis pool
func NewCacheStore(redisAddr, redisPort, certBase64 string) (CacheStore, error) {
	url := fmt.Sprintf("%s:%s", redisAddr, redisPort)

	opt := &redis.Options{
		Addr: url,
	}

	if certBase64 != "" {
		certFile, err := base64.StdEncoding.DecodeString(certBase64)
		if err != nil {
			log.Printf("failed to decode cert for cache store", err)
			return nil, err
		}

		decodedCert, _ := pem.Decode(certFile)
		if decodedCert == nil {
			log.Printf("failed to decode cert for cache store", err)
			return nil, err
		}
		cert, err := x509.ParseCertificate(decodedCert.Bytes)
		if err != nil {
			log.Printf("failed to decode cert for cache store", err)
			return nil, err
		}

		var certs tls.Certificate
		certs.Certificate = append(certs.Certificate, cert.Raw)

		certsPool := x509.NewCertPool()
		certsPool.AddCert(cert)

		tlsConfig := &tls.Config{
			ClientAuth: tls.RequireAnyClientCert,
			ClientCAs:  certsPool,
			// Certificates: []tls.Certificate{certs},
			// ServerName:   redisAddr,
		}

		opt.TLSConfig = tlsConfig
	}

	db := redis.NewClient(opt)

	// test connection
	pingResp, err := db.Ping().Result()
	if err != nil {
		log.Printf("failed ping cache store", err)
		return nil, err
	}
	log.Printf("ping resp: " + pingResp)

	return &RedisCacheStore{
		conn: db,
	}, nil
}

// NewCacheStore initialize connection for redis and store it into redis pool
func NewCacheStore2(redisAddr, redisPort, password string) (CacheStore, error) {
	url := fmt.Sprintf("%s:%s", redisAddr, redisPort)

	opt := &redis.Options{
		Addr:     url,
		Password: password,
	}

	db := redis.NewClient(opt)

	// test connection
	pingResp, err := db.Ping().Result()
	if err != nil {
		log.Printf("failed ping cache store", err)
		return nil, err
	}
	log.Printf("ping resp: " + pingResp)

	return &RedisCacheStore{
		conn: db,
	}, nil
}

// NewCacheStore initialize connection for redis and store it into redis pool
func NewCacheStore3(redisAddr, redisPort string) (CacheStore, error) {
	url := fmt.Sprintf("%s:%s", redisAddr, redisPort)

	opt := &redis.Options{
		Addr: url,
	}

	db := redis.NewClient(opt)

	// test connection
	_, err := db.Ping().Result()
	if err != nil {
		log.Printf("failed ping cache store", err)
		return nil, err
	}
	//utilsLog.Info(context.TODO(), "ping resp: "+pingResp)

	return &RedisCacheStore{
		conn: db,
	}, nil
}
