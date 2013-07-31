package main
import (
	"github.com/garyburd/redigo/redis"
	"strconv"
)

func getRedisConnection() (redis.Conn, error) {
	// TODO: Replace with pool
	// TODO: Add port to config
	// TODO: Ensure that connections are being closed properly
	return redis.Dial("tcp", redisServerPort)
}


func setSecretUser(secretType string, hash string, userId string) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("SET", "secret:" + secretType + ":" + hash, userId)
	return err
}

func getSecretUser(secretType string, hash string) (string, error) {
	c, err := getRedisConnection()
	if err != nil {
		return "", err
	}
	return redis.String(c.Do("GET", "secret:" + secretType + ":" + hash))
}

func deleteSecretUser(secretType string, hash string) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("DEL", "secret:" + secretType + ":" + hash)
	return err
}

func setUserAttribute(userId string, attribute string, data string) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("HSET", userId, attribute, data)
	return err
}

func getUserAttribute(userId string, attribute string) (string, error) {
	c, err := getRedisConnection()
	if err != nil {
		return "", err
	}
	return redis.String(c.Do("HGET", userId, attribute))
}

func deleteUserAttribute(userId string, attribute string) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("HDEL", userId, attribute)
	return err
}

func setUserCache(userId string, attribute string, data string, ttl int) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("SET", userId + ":cache:" + attribute, data, "EX", strconv.Itoa(ttl))
	return err
}

func getUserCache(userId string, attribute string) (string, error) {
	c, err := getRedisConnection()
	if err != nil {
		return "", err
	}
	return redis.String(c.Do("GET", userId + ":cache:" + attribute))
}

func deleteUserCache(userId string, attribute string) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("DEL", userId + ":cache:" + attribute)
	return err
}

func pushUserListTrim(userId string, name string, data string, size int) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("LPUSH", userId + ":" + name, data)
	if err != nil {
		return err
	}
	_, err = c.Do("LTRIM", userId + ":" + name, "0", strconv.Itoa(size))
	return err
}

func getUserListFront(userId string, name string) (string, error) {
	c, err := getRedisConnection()
	if err != nil {
		return "", err
	}
	return redis.String(c.Do("LINDEX", userId + ":" + name, "0"))
}

func getUserList(userId string, name string) ([]string, error) {
	c, err := getRedisConnection()
	if err != nil {
		return nil, err
	}
	return redis.Strings(c.Do("LRANGE", userId + ":" + name, "0", "-1"))
}

func setUserMap(userId string, name string, key string, data string) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("HSET", userId + ":" + name, key, data)
	return err
}

func getUserMap(userId string, name string, key string) (string, error) {
	c, err := getRedisConnection()
	if err != nil {
		return "", err
	}
	return redis.String(c.Do("HGET", userId + ":" + name, key))
}

func deleteUserMap(userId string, name string, key string) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("HDEL", userId + ":" + name, key)
	return err
}

func setUserFlag(userId string, name string, flag string) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("SADD", userId + ":" + name, flag)
	return err
}

func getUserFlags(userId string, name string) ([]string, error) {
	c, err := getRedisConnection()
	if err != nil {
		return nil, err
	}
	return redis.Strings(c.Do("SMEMBERS", userId + ":" + name))
}

func unsetUserFlag(userId string, name string, flag string) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("SREM", userId + ":" + name, flag)
	return err
}

func hasFlag(flags []string, flag string) bool {
	for _, v := range flags {
		if v == flag {
			return true
		}
	}
	return false
}

func userSubscribe(userId string, channel string) (*redis.PubSubConn, error) {
	c, err := getRedisConnection()
	if err != nil {
		return nil, err
	}
	psc := redis.PubSubConn{c}
	err = psc.Subscribe(userId + ":" + channel)
	if err != nil {
		return nil, err
	}
	return &psc, nil
}

func userSubscribeExisting(psc *redis.PubSubConn, userId string, channel string) error {
	return psc.Subscribe(userId + ":" + channel)
}

func userSubscriber() (*redis.PubSubConn, error) {
	c, err := getRedisConnection()
	if err != nil {
		return nil, err
	}
	psc := redis.PubSubConn{c}
	return &psc, nil
}

func userPublish(userId string, channel string, data string) error {
	c, err := getRedisConnection()
	if err != nil {
		return err
	}
	_, err = c.Do("PUBLISH", userId + ":" + channel, data)
	return err
}