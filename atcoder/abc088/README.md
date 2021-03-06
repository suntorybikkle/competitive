# abc088

# D

- 色を変更する数の最大値を求める問題
- ゴールまでのルートは自由であるが、出来るだけ多くの箇所の色を変更したいため必然的に最短経路問題に帰着する
- スタートからゴールまでの最短経路は幅優先探索で求めることができる
  - 計算量 H * W
- ゴールまでの最短経路が求まればルートは何でもよいため、元々が壁ではない中から、最短経路分引けばよい

## BFSアルゴリズムについて

探索中、新しい探索マスをキューに入れる際、以下のようにif条件式を書いていた
以下では、その必要はないが、多くの場合境界条件とその他の問題文に依存する条件を分けていた
なぜなら、問題文に依存する条件はout of indexになる可能性があるためだ。 (griidの周囲を囲うなどの対処はある)
ただ、if文内に条件式が並ぶ場合、左から順に評価されるっぽいため、その必要はない模様  (右に書けばよい)

```go
ny, nx := y+d[0], x+d[1]
if ny < 0 || ny == h || nx < 0 || nx == w { // 境界条件
  continue
}
if grid[ny][nx] == 0 { // 何らかの問題文に関係する条件
  grid[ny][nx] = grid[y][x] + 1
  q = append(q, []int{ny, nx})
}
```

以下のように、先にキューに入れてしまうパターンもあるっぽい (BFSの書き方の違い)

```c++
while (!q.empty()) {
    auto [i, j, cost] = q.front();
    q.pop();
    if (i < 0 || i > h - 1 || j < 0 || j > w - 1 || costs[i][j] != -1 || s[i][j] == '#') {
      continue; // キューから取り出したものが各種条件に引っかかっているかは、取り出し後判断
    }
    costs[i][j] = cost;
    q.push({i - 1, j, cost + 1}); // 境界条件などは精査せず、先にキューに入れる
    q.push({i + 1, j, cost + 1});
    q.push({i, j - 1, cost + 1});
    q.push({i, j + 1, cost + 1});
}
```
