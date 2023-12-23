// @Title
// @Description
// @Author  Wangwengang  2023/12/22 09:34
// @Update  Wangwengang  2023/12/22 09:34
package request

// httpgate.Request
// message Request{
// string token = 1;
// string v = 2;
// string sign = 3;
// string signMethod = 4;
// string timeStamp = 5;
// bytes data = 6;
// }
type RequestJson struct {
	Token      string                 `json:"token"`
	V          string                 `json:"v"`
	Sign       string                 `json:"sign"`
	SignMethod string                 `json:"signMethod"`
	TimeStamp  string                 `json:"timeStamp"`
	Data       map[string]interface{} `json:"data"`
}
