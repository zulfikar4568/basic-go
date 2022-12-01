
Akan menjalankan semua testing file

```bash
cd helper
go test
 # atau
go test -v
# atau ini me running spesifik Function
go test -v -run TestHelloWorldArip

#atau ini me running seluruh folder child nya
go test ./... -v
```