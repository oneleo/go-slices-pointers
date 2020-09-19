# go-slices-pointers
 golang 語言的 slice 與儲存 pointer 的 slice

## Slice A 範例
- 建置一個包含 3 個 string 元素的 slice
- 呼叫傳遞 pointer 的 func

```go
var a = []string{"1", "2", "3"}

changeSliceA(&a)
```

## Slice B 範例
- 建置一個包含 3 個 string pointer 元素的 slice
- 呼叫傳遞 value 的 func

```go
s1 := "1"
s2 := "2"
//s3 := "3"
var b = []*string{&s1, &s2, new(string)}

changeSliceB(b)
```

## Slice C 範例
- 建置一個包含 3 個 string pointer 元素的 slice
- 呼叫傳遞 pointer 的 func

```go
c := make([]*string, len(b))
copy(c, b)

changeSliceC(&c)
```

## 下載 & 執行
```sh
$> git clone git@github.com:oneleo/go-slices-pointers.git
$> cd ./go-slices-pointers

$> go run ./main.go
```