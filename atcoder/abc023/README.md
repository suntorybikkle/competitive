# abc023

# D

- どの順番で風船を割れば良いかと考えても、当然間に合わない
- 知りたい答えは、最小のペナルティだが、それ自体を探索する方針
- あるペナルティ以下で風船すべてを割ることができるかを判断することは、各風船がそのペナルティまでに何秒かかるかを計算することで確認できる
  - 計算量は、N * log N (ビンソートを用いることで N でも可能)
- あるペナルティの最小値を2部探索を用いることで時間内に実行することができる
  - 最終的な計算量 N * log N * log(H + NS)