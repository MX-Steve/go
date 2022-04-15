# 1. Keywords
1. Varint
2. TLV
3. Zigzag
4. protobuf
# 2. Question 
```go
// go
type User struct {
    Name string `json:name`
    Age int32 `json:age`
}
// json data
{
    "name":"HanMeimei"
    "age":30
} 
0011,0000,1010,1000
->
x1100001 x0101000
->
10101000 0110001

257
->
1,0   000,0001
x0000001, x0000001
1000000100000010

```