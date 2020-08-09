# Messages
```go
func main() {
	M:=messages.New("msg")
	err:=M.ReLoad("templates", ".tpl")
	if err!= nil{
		panic(err)
	}
	err=fmt.Errorf("error message")
	msg := M.Execute("error.tpl", struct{Err error; PID string}{err, "PID1"})
	log.Print(msg)

	M.SetDebugger(true)
	M.Debug("error_debug.tpl", struct{Err error; PID string}{err, "PID1"})
	M.SetDebugger(true)
	M.Debug("debug", struct{Err error; PID string}{err, "PID1"})
}
```
```bash
% go run cmd/debug_print/main.go
2020/08/09 07:59:22 Error message: [PID1] some error err=error message
2020/08/09 07:59:22 Debug message: [PID1] some error err=error message filename=/Users/alexsuslov/go/src/github.com/alexsuslov/messages/cmd/debug_print/main.go:20
2020/08/09 07:59:22 Debug message from err_aggr.tpl: [PID1] some error err=error message filename=/Users/alexsuslov/go/src/github.com/alexsuslov/messages/cmd/debug_print/main.go:22
```