package basic

import (
	"fmt"
	"net/http"
	"bytes"
	"time"
	"testing"
)

var url = "http://192.168.50.23:8100/" //节点ip:port
var totalNum = 10000                   //测试总次数
var routineNum = 100                   //协程数


func TestSendEthTransaction(t *testing.T) {
	fmt.Println("begin time: ", time.Now())

	for i := 0; i < routineNum; i++ {
		go test_sendTX()
	}

	select {}
}


//测试发送交易
func test_sendTX() {
	//每个协程发送 totalNum/routineNum 次交易请求
	for i := 0; i < (totalNum / routineNum); i++ {
		client := &http.Client{}
		//req := `{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":83}`
		req := `{
			"jsonrpc":"2.0",
			"method":"eth_sendTransaction",
			"params":[{
				"from": "0xf9ba1df4786963fbc6c4cd41af7f43df10b4fb72",
			 	"to": "0xcfbcdbe6c0953caa47f2d0acadaf0964a74fdb87",
			 	"gas": "0xc350",
			 	"gasPrice": "0x9184e72a000",
			 	"value": "0x2386f26fc10000",
			 	"data": ""
				}],
			"id":1}`

		request, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(req)))
		request.Header.Set("Content-type", "application/json")

		//response, err := client.Do(request)
		_, err := client.Do(request)
		if err != nil {
			fmt.Println("err: ", err)
			fmt.Println("ERROR! cur time: ", time.Now())
		}
		fmt.Println("cur time: ", time.Now())
		//	if response.StatusCode == 200 {
		//		body, _ := ioutil.ReadAll(response.Body)
		//		fmt.Println(string(body))
		//		fmt.Println("cur time: ", time.Now())
		//	}
	}
}
