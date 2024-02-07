protoc -I. --gogofast_out=. --rpcx_out=. ./pbgate/gate.proto
protoc -I. --gogofast_out=plugins=rpcx:. ./httpgate/httpgate.proto
protoc -I. --gogofast_out=. ./logic/logic.proto
protoc --proto_path=proto --gogofast_out=proto --gogofast_opt=paths=source_relative pbcommon/pbcommon.proto
protoc --proto_path=proto --gogofast_out=proto --gogofast_opt=paths=source_relative --simple_out=model --simple_opt=paths=source_relative pbcmdService/pbcmdService.proto
