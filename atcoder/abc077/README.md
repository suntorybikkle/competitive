# abc077

# C

- 真ん中のパーツを基準にするとn * nで可能だが、間に合わない
- 上下のパーツをソートすることで、二分探索を利用することができる
- 問題の本質とはそれるが、Go独特の問題
  - `fmt.Scan()`は読み込みが遅く、10**5ほどで実行時間が足りなくなるそうです (今回の入力は10 ** 5ほど)

```go
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Split(bufio.ScanWords)

    sc.Scan()
    n, _ := strconv.Atoi(sc.Text())
    // ...
}
```