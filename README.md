# go-slices-pointers
 golang 語言的 slice 與儲存 pointer 的 slice

## Slice A 範例
- 建置一個包含 3 個 string 元素的 slice
- 呼叫傳遞 value 的 func

```go
var a = []string{"1", "2", "3"}

changeSliceA(a)
```

## Slice B 範例
- 建置一個包含 3 個 string 元素的 slice
- 呼叫傳遞 pointer 的 func

```go
b := make([]string, len(a))
copy(b, a)

changeSliceB(&b)
```

## Slice C 範例
- 建置一個包含 3 個 string pointer 元素的 slice
- 呼叫傳遞 value 的 func

```go
s1 := "1"
s2 := "2"
//s3 := "3"
var c = []*string{&s1, &s2, new(string)}

changeSliceC(c)
```

## Slice D 範例
- 建置一個包含 3 個 string pointer 元素的 slice
- 呼叫傳遞 pointer 的 func

```go
d := make([]*string, len(c))
copy(d, c)

changeSliceD(&d)
```

## 下載 & 執行
```sh
$> git clone git@github.com:oneleo/go-slices-pointers.git
$> cd ./go-slices-pointers

$> go run ./main.go
```

## 輸出
```sh
a at 0xc000004480       a = [3 2 3]
b at 0xc0000044a0       b = [3 2 3 4]
c at 0xc0000044c0       c = [0xc0000421f0 0xc000042200 0xc000042210]
d at 0xc000004500       d = [0xc0000421f0 0xc000042200 0xc000042210 0xc000042230]
```