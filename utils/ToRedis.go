package utils

import (
	"WeiYangWork/global"
	"strconv"
)

// CacheMessage 缓存消息到队伍
func CacheMessage(teamID uint, message string) error {
	key := "team:" + strconv.Itoa(int(teamID)) + ":messages"
	err := global.Redis.LPush(key, message).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetRecentMessages 获取队伍最近的消息
func GetRecentMessages(teamID uint) ([]string, error) {
	key := "team:" + strconv.Itoa(int(teamID)) + ":messages"
	messages, err := global.Redis.LRange(key, 0, 9).Result() // 获取最近10条消息
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// PublishMessage 发布消息到队伍频道
func PublishMessage(teamID uint, message string) error {
	channel := "team:" + strconv.Itoa(int(teamID))
	err := global.Redis.Publish(channel, message).Err()
	return err
}
