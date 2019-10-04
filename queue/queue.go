package queue

import (
	"container/list"
	"errors"
	"golang/iso8583"
	"golang/logging"
	"net"
	"sync"
	"time"
)

// MessageStatus mark a record is process/processed
type messageStatus int

const (
	Pending messageStatus = iota
	InProgress
	Completed
)

// Message store message iso from client
type Message struct {
	ClientConn   net.Conn
	RequestTime  time.Time
	ResponseTime time.Time
	RequestData  *iso8583.ISO8583Data
	ResponseData *iso8583.ISO8583Data
	Status       messageStatus
}

// var iso8583Queue *list.List

//Iso8583Queue
type Iso8583Queue struct {
	sync.Mutex
	MsgList *list.List
}

var iso8583Queue *Iso8583Queue

//NewElement return new message
func NewElement(client net.Conn, requestData *iso8583.ISO8583Data, status messageStatus) Message {
	return Message{
		ClientConn:   client,
		RequestTime:  time.Now(),
		ResponseTime: time.Now(),
		RequestData:  requestData,
		Status:       status,
	}
}

// InitQueue init queue as list
func InitQueue() {
	iso8583Queue = &Iso8583Queue{
		MsgList: list.New(),
	}
	logging.GetLog().Info("Init queue with queue len is: ", iso8583Queue.MsgList.Len())
}

// GetQueue get queue
func GetQueue() *Iso8583Queue {
	return iso8583Queue
}

// Put put element to back queue
func Put(ele Message) {
	iso8583Queue.Lock()
	defer iso8583Queue.Unlock()
	iso8583Queue.MsgList.PushBack(Message{
		ClientConn:   ele.ClientConn,
		RequestTime:  ele.RequestTime,
		ResponseTime: ele.ResponseTime,
		RequestData:  ele.RequestData,
		ResponseData: ele.ResponseData,
		Status:       ele.Status,
	})
}

// Get get first element
func Get() (Message, error) {
	if IsEmpty() {
		return Message{}, errors.New("queue is nil")
	}
	iso8583Queue.Lock()
	defer iso8583Queue.Unlock()
	for e := iso8583Queue.MsgList.Front(); e != nil; e = e.Next() {
		if e.Value.(Message).Status == Pending {
			iso8583Queue.MsgList.Remove(e)
			return e.Value.(Message), nil
		}
		continue
	}
	return Message{}, errors.New("queue is nil")
}

// IsEmpty check queue empty or not, return true mean queue empty otherwise return false
func IsEmpty() bool {
	iso8583Queue.Lock()
	defer iso8583Queue.Unlock()
	return iso8583Queue.MsgList.Len() == 0
}
