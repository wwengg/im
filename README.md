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

H 头部码
T 方法码 = http 里的路由 user/createUser [post] = 1 1= json 2 = protobuf
L 长度
V 内容包体
CRC 算法 数据

## VSCODE dev

.vscode/launch.json

```
{
  // 使用 IntelliSense 了解相关属性。
  // 悬停以查看现有属性的描述。
  // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
  "version": "0.2.0",
  "configurations": [
    {
      "name": "gate",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/main.go", // 执行程序，这里是项目默认的根目录，可以指定文件执行，如"${workspaceFolder}\\cmd\\app.go"
      "args": ["gate"] // 程序启动参数 如 ["-env" , 'prod']
    },
    {
      "name": "http",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/main.go", // 执行程序，这里是项目默认的根目录，可以指定文件执行，如"${workspaceFolder}\\cmd\\app.go"
      "args": ["http"] // 程序启动参数 如 ["-env" , 'prod']
    },
    {
      "name": "logic",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/main.go", // 执行程序，这里是项目默认的根目录，可以指定文件执行，如"${workspaceFolder}\\cmd\\app.go"
      "args": ["logic -c ./logic.yaml"] // 程序启动参数 如 ["-env" , 'prod']
    }
  ]
}
```
