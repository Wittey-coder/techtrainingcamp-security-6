package session

import (
	"encoding/json"
	"errors"
	"sync"
)
import "github.com/garyburd/redigo/redis"

const (
	NONE     = false
	MODIFIED = true
)

type RedisSession struct {
	sessionId string
	lock      sync.RWMutex
	pool      *redis.Pool
	//设置session, 可以先放在内存中
	//延迟加载，提升性能
	//session的数据有很多，放在表里
	sessionData map[string]interface{}
	//记录map是否被操作
	flag bool
}

// NewRedisSession 新建一个Session
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionId:   id,
		pool:        pool,
		sessionData: make(map[string]interface{}, 16),
		flag:        NONE,
	}
	return s
}

// Set 设置Session中某个key的数据
func (r *RedisSession) Set(key string, value interface{}) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.sessionData[key] = value
	r.flag = MODIFIED
	return nil
}

// Get 获取Session中某个key的数据
func (r *RedisSession) Get(key string) (interface{}, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	// 先在内存里找数据
	result, ok := r.sessionData[key]
	// 如果内存没有，就把Redis里的数据加载上来
	if !ok {
		err := r.loadFromRedis()
		if err != nil {
			return nil, err
		}
		result, ok = r.sessionData[key]
		if !ok {
			return nil, errors.New("the key is not exists")
		}
	}
	return result, nil
}

// Delete 删除Session中的某个键
func (r *RedisSession) Delete(key string) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	// 先从Redis加载出来
	err := r.loadFromRedis()
	if err != nil {
		return err
	}
	// 删除
	delete(r.sessionData, key)
	// 直接写入
	err = r.Save()
	if err != nil {
		return err
	}
	return nil
}

// Save 将内存中的key保存到Redis
func (r *RedisSession) Save() error {
	r.lock.Lock()
	defer r.lock.Unlock()
	// 查看是否被修改
	if !r.flag {
		return nil
	}
	// 序列化
	data, err := json.Marshal(r.sessionData)
	if err != nil {
		return err
	}

	// 获取连接然后设置值和过期时间
	conn := r.pool.Get()
	_, err = conn.Do("SET", r.sessionId, string(data))
	if err != nil {
		return err
	}
	_, err = conn.Do("expire", r.sessionId, 600)
	if err != nil {
		return err
	}

	// 清除内存中的数据然后把修改标记改回去
	r.flag = NONE
	r.loadFromRedis()
	return nil
}

// loadFromRedis 把redis里的数据加载到内存
func (r *RedisSession) loadFromRedis() error {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return err
	}
	data, err := redis.String(reply, err)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), &r.sessionData)
	if err != nil {
		return err
	}

	return nil
}
