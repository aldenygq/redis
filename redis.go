package redis

import (
    "github.com/go-redis/redis/v8"
    "context"
    "time"
)
type RClient struct {
    Client *redis.Client
}


func NewRedisClient(url,pwd string) *RClient {
    client := redis.NewClient(&redis.Options{
        Addr:     url, // Redis服务器地址和端口号
        Password: pwd,               // Redis密码（如果有）
        DB:       0,                // Redis数据库编号
    })
    return &RClient{
        Client: client,
    }
}

func (r *RClient) SetValue(key string, value string, expiration time.Duration) error {
    err := r.Client.Set(context.Background(), key, value, expiration).Err()
    if err != nil {
        return err
    }
    return nil
}

func (r *RClient) GetValue(key string) (string, error) {
    val, err := r.Client.Get(context.Background(), key).Result()
    if err == redis.Nil {
        return "", nil // 如果键不存在，则返回空字符串而不是错误
    } else if err != nil {
        return "", err
    }
    return val, nil
}

func (r *RClient) Close() {
    r.Close()
}
