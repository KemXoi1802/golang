package main

import (
	"golang/queue"
	"golang/server"
	"golang/utils"
)

func main() {
	queue.InitQueue()
	utils.InitManager()
	s := server.NewServer()
	s.Start()

	// iso := ISO8583.DefaultIso8583Data()
	// iso.SetMTI(800)
	// iso.PackField(2, "4761080007841110")
	// iso.PackField(3, "920000")
	// iso.PackField(11, "000001")
	// iso.PackField(12, "182138")
	// iso.PackField(13, "0920")
	// iso.PackField(24, "0032")
	// iso.PackField(39, "00")
	// iso.PackField(41, "01100191")
	// msg, length, _ := iso.Pack()
	// iso.Parse()

	// fmt.Println(msg[:length])
	// fmt.Println(iso.CheckBit(2))
	// fmt.Println(hex.DecodeString("080020380100028000001647610800078411109200000000011821380920003230303031313030313931"))
	// fmt.Println(hex.EncodeToString(msg[:length]))
	// fmt.Println("080020380100028000001647610800078411109200000000011821380920003230303031313030313931")

	// "0026" + "6000100000" + "003230303132333938373435"
	// "0810 2038010002800000 920000 000001 182138 0909 0032 3030 3132333938373435"
	// "0810 2038010002800000 920000000001182138 0920 0032 30303031313030313931"
	// "0810 2038010002800000 920000 000001 182138 0920 0032 3030 3031313030313931"
}
