package queue

import (
	"golang/iso8583"
	"golang/utils"
	"container/list"
	"errors"
	"net"
	"time"
)

// MessageQueueStatus mark a record is process/processed
type messageQueueStatus int

const (
	New messageQueueStatus = iota
	InProgress
	Completed
)

// MessageQueue store message iso from client
type MessageQueue struct {
	mClientConn   net.Conn
	mRequestTime  time.Time
	mResponseTime time.Time
	mRequestData  *iso8583.ISO8583Data
	mResponseData iso8583.ISO8583Data
	mStatus       messageQueueStatus
}

var iso8583Queue *list.List

//NewElement return new message
func NewElement(client net.Conn, requestData *iso8583.ISO8583Data, status messageQueueStatus) MessageQueue {
	return MessageQueue{
		mClientConn:   client,
		mRequestTime:  time.Now(),
		mResponseTime: time.Now(),
		mRequestData:  requestData,
		mStatus:       status,
	}
}

// InitQueue init queue as list
func InitQueue() {
	iso8583Queue = list.New()
	utils.GetLog().Info("Init queue with queue len is: ", iso8583Queue.Len())
}

// GetQueue get queue
func GetQueue() *list.List {
	return iso8583Queue
}

// Put put element to back queue
func Put(ele MessageQueue) {
	iso8583Queue.PushBack(MessageQueue{
		mClientConn:   ele.mClientConn,
		mRequestTime:  ele.mRequestTime,
		mResponseTime: ele.mResponseTime,
		mRequestData:  ele.mRequestData,
		mResponseData: ele.mResponseData,
		mStatus:       ele.mStatus,
	})
	utils.GetLog().Info("Length of Queue after PUT new element: ", iso8583Queue.Len())
}

// GetFront get first element
func GetFront() (MessageQueue, error) {

	for true {
		if iso8583Queue.Len() == 0 {
			continue
		}
		utils.GetLog().Info("Get front from queue")
		ele := iso8583Queue.Front()
		utils.GetLog().Info("Length of Queue after POP front element: ", iso8583Queue.Len())
		return ele.Value.(MessageQueue), nil
	}
	return MessageQueue{}, errors.New("queue is nil")
}
