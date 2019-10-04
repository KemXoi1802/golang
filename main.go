package main

import (
	"golang/queue"
	"golang/server"
	"golang/logging"
	"golang/config"
	"os"
	"io/ioutil"
	"encoding/xml"
)

// <?xml version="1.0" encoding="UTF-8"?>
// <Transactions>
//     <Transaction type="810">
//         <Field id="2" value="logon"/>
//     </Transaction>

//     <Transaction type="210">
//         <Field id="3" value="sale"/>
//     </Transaction>

//     <Transaction type="510">
//         <Field id="4" value="tong ket"/>
//     </Transaction>
// </Transactions>

type Transactions struct {
	XMLName xml.Name `xml:"Transactions"`
	Transactions []Transaction `xml:"Transaction"`
}

type Transaction struct {
	XMLName xml.Name `xml:"Transaction"`
	Fields []Field `xml:"Field"`
}

type Field struct {
	XMLName xml.Name `xml:"Field"`
	Typ string `xml:"type,attr"`
	ID string `xml:"id,attr"`
	Value string `xml:"value,attr"`
}
func main() {
	logging.Init("log.txt", "debug")
	xmlFile, _ := os.Open("response.xml")
	defer xmlFile.Close()
	v, _ := ioutil.ReadAll(xmlFile)

	var txn Transactions
	xml.Unmarshal(v, &txn)

	// logging.GetLog().Info(txn.Transactions[0].Fields[0].ID)
	// logging.GetLog().Info(txn.Transactions[0].Fields[0].Value)

	// logging.GetLog().Info(txn.Transactions[0].Fields[1])
	// logging.GetLog().Info(txn.Transactions[0].Fields[1])

	for i:=0; i< len(txn);i++ {
		if txn.Transactions[i].Fields[] == "810" {
			logging.GetLog().Info(txn.Transactions[i])
		}
	}
	
	queue.InitQueue()
	config.Init()
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
