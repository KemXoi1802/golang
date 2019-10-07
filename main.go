package main

import (
	"fmt"
	"golang/config"
	"golang/logging"
	"golang/queue"
	"golang/server"
	"strings"
	"time"

	"math/rand"

	"github.com/spf13/viper"
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

// type Transactions struct {
// 	XMLName xml.Name `xml:"Transactions"`
// 	Transactions []Transaction `xml:"Transaction"`
// }

// type Transaction struct {
// 	XMLName xml.Name `xml:"Transaction"`
// 	Fields []Field `xml:"Field"`
// }

// type Field struct {
// 	XMLName xml.Name `xml:"Field"`
// 	Typ string `xml:"type,attr"`
// 	ID string `xml:"id,attr"`
// 	Value string `xml:"value,attr"`
// }

//GenerateRandomBytes generate random number specific length
// func GenerateRandomBytes(n int) ([]byte, error) {
// 	b := make([]byte, n)
// 	_, err := rand.Read(b)
// 	// Note that err == nil only if we read len(b) bytes.
// 	if err != nil {
// 		return nil, err
// 	}

// 	return b, nil
// }

// //GenerateRandomString random string
// func GenerateRandomString(s int) (string, error) {
// 	b, err := GenerateRandomBytes(s)
// 	return base64.URLEncoding.EncodeToString(b), err
// }

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

const charset = "0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//StringWithCharset test
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//String test
func String(length int) string {
	return StringWithCharset(length, charset)
}

func main() {
	// logging.Init("log.txt", "debug")
	// xmlFile, _ := os.Open("response.xml")
	// defer xmlFile.Close()
	// v, _ := ioutil.ReadAll(xmlFile)

	// var txn Transactions
	// xml.Unmarshal(v, &txn)

	// // logging.GetLog().Info(txn.Transactions[0].Fields[0].ID)
	// // logging.GetLog().Info(txn.Transactions[0].Fields[0].Value)

	// // logging.GetLog().Info(txn.Transactions[0].Fields[1])
	// // logging.GetLog().Info(txn.Transactions[0].Fields[1])

	// for i:=0; i< len(txn);i++ {
	// 	if txn.Transactions[i].Fields[] == "810" {
	// 		logging.GetLog().Info(txn.Transactions[i])
	// 	}
	// }

	n := random(100000000000, 999999999999)

	logging.GetLog().Info("random numbers: ", n)

	str := String(6)

	logging.GetLog().Info(str)

	viper.SetConfigFile("response.json")
	if err := viper.ReadInConfig(); err != nil {
		logging.GetLog().Info("Error reading config file, %s", err)
	}

	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	// v := viper.Get("810")
	v := viper.Sub("810")

	fmt.Printf("Value: %v, Type: %T\n", v, v)
	s1 := v.Get("fields")
	fmt.Println(strings.Split(fmt.Sprintf("%v", s1), ","))
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
