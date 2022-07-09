/**
  @author:panliang
  @data:2022/6/22
  @note
**/
package dispatch

import (
	"github.com/gorilla/websocket"
	"im-services/config"
	"im-services/pkg/redis"
	"sync"
	"time"
)

var (
	mux sync.Mutex
)

type DispatchService struct {
}

type DispatchServiceInterface interface {
	SetDispatchNode(uid string, node string) // 设置当前节点信息
	GetDispatchNode(uid string, node string) // 获取当前节点信息
	MessageDispatch(uid string, node string) // 获取当前节点信息
	IsDispatchNode(uid string, node string)  // 获取当前节点信息
	DetDispatchNode(uid string)              //删除当前节点
}

func (Service *DispatchService) IsDispatchNode(uid string) (bool, string) {

	n, _ := redis.RedisDB.Exists(uid).Result()

	if n > 0 {
		uNode := Service.GetDispatchNode(uid)
		return true, uNode
	} else {
		return false, ""
	}
}

func (Service *DispatchService) GetDispatchNode(uid string) string {
	return redis.RedisDB.Get(uid).Val()
}

func (Service *DispatchService) DetDispatchNode(uid string) {
	redis.RedisDB.Del(uid)
}

func (Service *DispatchService) SetDispatchNode(uid string) {
	mux.Lock()
	redis.RedisDB.Set(uid, config.Conf.Server.Node, time.Hour*24)
	mux.Unlock()
}

func (Service *DispatchService) MessageDispatch(conn *websocket.Conn) {

}
