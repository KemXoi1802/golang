package server

import (
	"bufio"
	"golang/iso8583"
	"golang/logging"
	"golang/queue"
	"net"
	"runtime"
	"sync"
	"time"
)

const (
	maxBufferSize = 4096
	maxByteLength = 2
)

// ISO8583Client information to handle stream connection, for each connection will be created new ISO8583client
type ISO8583Client struct {
	sync.WaitGroup
	mClientCon net.Conn
	mServer    *Server
}

//Listen listen data from client and put it to queue
func (client *ISO8583Client) Listen() {
	r := bufio.NewReader(client.mClientCon)
	for client.mServer.mIsRunning {
		var err error
		bytesLen := make([]byte, maxByteLength)
		iso8583data := make([]byte, maxBufferSize)
		if bytesLen, err = ReadByte(r, 2); err == nil {
			if length, err := iso8583.MessageLengthToInt(client.mServer.mLengthType, bytesLen); err == nil {
				iso8583data, err = ReadByte(r, length)
				if len(iso8583data) != length {
					logging.GetLog().Info("invalid length")
					continue
				}
				msg := iso8583.NewIso8583Data(iso8583data, length)
				if msg != nil {
					ele := queue.NewElement(client.mClientCon, msg, queue.Pending)
					queue.Put(ele)
				} else {
					logging.GetLog().Debug("MSG IS NILL")
				}

			}
		}
	}
	client.Done()
}

//ReadByte read byte specific length
func ReadByte(r *bufio.Reader, bytesRead int) ([]byte, error) {
	output := make([]byte, bytesRead, bytesRead)
	var err error
	for i := 0; i < bytesRead; i++ {
		output[i], err = r.ReadByte()
	}
	return output, err
}

//ProcessMessage run process msg
func (client *ISO8583Client) ProcessMessage() {
	for true {
		if queue.IsEmpty() == false {
			if message, err := queue.Get(); err == nil {
				if message.RequestData.IsRequest() {
					if runtime.NumGoroutine() > 50 {
						time.Sleep(30 * time.Second)
						queue.Put(message)
					}
					go func(msg interface{}) {
						message.RequestData.Unpack()
						message.ResponseData = message.RequestData.Clone()
						message.ResponseData.SetResponseMTI()
						message.ResponseData.SwapNII()
						message.ResponseData.PackField(39, "00")
						message.ResponseData.PackField(62, "12345678901213141512133456870923")
						message.ResponseData.Pack()
						data, _ := message.ResponseData.BuildMsg()
						message.ClientConn.Write(data)
					}(message)
				}
			}
		}
		time.Sleep(1 / 1000 * time.Second)
	}
	client.Done()
}
