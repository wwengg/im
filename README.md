# IM

## install

## cobra-cli
```
go install github.com/spf13/cobra-cli@latest
```

## gate

```
// HTLV+CRC, H header code, T function code, L data length, V data content
//+------+-------+-------+--------+--------+
//| H    | T     | L     | V      | CRC    |
//| 1Byte| 1Byte | 1Byte | NBytes | 2Bytes |
//+------+-------+-------+--------+--------+

// HeaderCode FunctionCode DataLength Body                         CRC
// A2         10           0E         0102030405060708091011121314 050B
```


H  头部码
T  方法码 = http里的路由 user/createUser [post]   =   1   1= json  2 = protobuf
L  长度
V  内容包体
CRC算法 数据
