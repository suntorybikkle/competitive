# Golangで競技プログラミングを解く際の基本

Golangで競技プログラミングを解いている際に使った、よく使う定型文のようなものをメモする  
主に以下の特徴を持つ
- Golangの文法の基本
- コンテスト中、実際に使った
- 2度は調べたことのあること
- 煩雑

逆にアルゴリズムなどのメモはない

※ コードブロック内は文法を無視しているケースもある

## 変数

Golangでは、宣言した変数のスコープは括弧の中だけで、インデントのない宣言はパッケージレベルになる

```go
var a int
var a, b int                      // 複数同時に宣言できる
var b, f, s = true, 2.3, "string" // 異なる型の変数宣言
x := 1                            // 初期化と宣言を同時に行う
x, y := 1, 2
```

## 入力

`fmt.Scan()`を使った入力は遅く、10 ** 5 くらいが限界  
改行などを気にせず、どんどん次の入力を受け取ることができる

```go
var a int
fmt.Scan(&a)
M := make([]int, n)
for i := range M {
  fmt.Scan(&M[i])  // 複数の入力への対応 (改行でも、スペースでも)
}
```

[Go 言語で標準入力から読み込む競技プログラミングのアレ --- 改訂第二版](https://qiita.com/tnoda_/items/b503a72eac82862d30c6)
[[Go] AtCoderで標準入力するためのスニペット](https://qrunch.net/@koralle/entries/8WVq0MS18Aws4s84?ref=qrunch)

```go
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}

func main() {
    sc.Split(bufio.ScanWords)
    n := nextInt()
}
```

## スライス

スライスは配列と異なり、可変長  

宣言

```go
var q [][]int           // 多次元スライス宣言
Q := make([]int, n)     // 長さnのスライスを生成 intの場合0で初期化される
W := make([][]int, n)   // 多次元の場合
for i := range W {
  W[i] = make([]int, m) // 別途初期化が必要だが、宣言済みなので=を使う
}
```

スライスの要素への参照

```go
q[i]          // i番目
q[i:j]        // iからj含まないまでのスライスを取得
q[len(q) - 1] // 最後の要素
```

## for文

基本

```go
for i:=0; i<N; i++ {}
for i, m := range M {} // 配列Mを要素mで回す、iはインデックス
for i := range M {}    // 一つの場合は、インデックス
for _, m := range M {}
```

よく使う表現

```
for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
  ny, nx := y+d[0], x+d[1] // 格子状の空間を上下左右動く
}
for {} // 無限ループ
for len(q) != 0 {} // while文
```
