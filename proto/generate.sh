protoc -I. --gogofast_out=. --rpcx_out=. ./pbgate/gate.proto
protoc -I. --gogofast_out=plugins=rpcx:. ./httpgate/httpgate.proto
protoc -I. --gogofast_out=plugins=rpcx:. --rpcx_out=. ./logic/logic.proto