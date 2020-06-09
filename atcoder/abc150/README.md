# abc150

# C

- next_permutationで順列を生成し、比較を行えばACすることができる
- スライスなどの等しい確認は`reflect.DeepEqual`がよさそう
- 辞書順に並ぶため、各数列が辞書順で何番目かを直接計算する事もできる
- 考え方
  - その数列が3から始まると考えた場合、1から始まる数列と2から始まる数列の後ろであることが確定している
  - 数列が仮に3桁だった場合、1から始まる数列は(3-1)!通り考えられる
  - このようにそれより前に出現する数列の通りを計算し、各桁行うことでその数列の出現順序が計算できる
  - `c2.go`は辞書順を利用して回答してみた例









