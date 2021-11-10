package session

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"log"
	"sync"
	"time"
)

type RedisSessionManager struct {
	addr string // 主机地址

	password string // 密码

	pool *redis.Pool // 连接池

	lock sync.RWMutex

	sessionMap map[string]Session // 放在内存里的map，保存id和session的键值对
}

func NewRedisSessionManager() SessionManager {
	r := &RedisSessionManager{
		sessionMap: make(map[string]Session, 32),
	}
	return r
}

// Init 设置连接池
func (r *RedisSessionManager) Init(addr string, options ...string) error {
	if len(options) > 0 {
		r.password = options[0]
		r.addr = addr
		r.pool = &redis.Pool{
			// 连接
			Dial: func() (redis.Conn, error) {
				dial, err := redis.Dial("tcp", addr)
				if err != nil {
					return nil, err
				}
				_, err = dial.Do("AUTH", r.password)
				if err != nil {
					dial.Close()
					return nil, err
				}
				return dial, err
			},
			// 测试连接
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
			MaxIdle:         64,
			MaxActive:       1000,
			IdleTimeout:     240 * time.Second,
			Wait:            false,
			MaxConnLifetime: 0,
		}
	} else {
		r.addr = addr
		r.pool = &redis.Pool{
			// 连接
			Dial: func() (redis.Conn, error) {
				dial, err := redis.Dial("tcp", addr)
				if err != nil {
					return nil, err
				}
				_, err = dial.Do("AUTH", r.password)
				if err != nil {
					dial.Close()
					return nil, err
				}
				return dial, err
			},
			// 测试连接
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
			MaxIdle:         64,
			MaxActive:       1000,
			IdleTimeout:     240 * time.Second,
			Wait:            false,
			MaxConnLifetime: 0,
		}
	}
	return nil
}

// CreateSession 新建
func (r *RedisSessionManager) CreateSession() (session Session, id string, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	sessionId := uuid.NewV4().String()
	memorySession := NewRedisSession(sessionId, r.pool)
	r.sessionMap[sessionId] = memorySession
	return memorySession, sessionId, nil
}

// Get 获取session
func (r *RedisSessionManager) Get(sessionId string) (session *Session, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	get, ok := r.sessionMap[sessionId]
	if !ok {
		return nil, errors.New("不存在此Session")
	}
	return &get, nil
}

// Save 把所有session保存在redis里
func (r *RedisSessionManager) Save() error {
	for _, session := range r.sessionMap {
		err := session.Save()
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}

// Delete 删除Session
func (r *RedisSessionManager) Delete(sessionId string) error {
	delete(r.sessionMap, sessionId)
	return nil
}
