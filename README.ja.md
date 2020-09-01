# distro

ja | [en](./README.md)

distro パッケージは，任意の確率分布に従う疑似乱数生成器を実装します．このパッケージは次の機能を提供します．

- 与えられた引数を基に確率分布を生成します．分布は累積値が1になるように正規化されます．
- 確率密度関数（PDF）と累積分布関数（CDF）を提供します．分布中の指定インデックス値における確率値を返します．
- 確率分布に基づいて，ランダムにインデックス値を返します．

## Usage

### Installation

```
go get github.com/roamer7038/distro
```

### Example

[Zipfの法則](https://en.wikipedia.org/wiki/Zipf%27s_law)に従う疑似乱数生成器を実装します．以下の例は，Zipf則（偏りパラメータ`0.693`）に基づく偏りを持つ要素数`1000`個のコンテンツの順位を，ランダムに標準出力しています．

```
package main

import (
  "fmt"
  "github.com/roamer7038/distro"
  "math/rand"
  "time"
)

func main() {
  rs := rand.New(rand.NewSource(time.Now().UnixNano()))
  z, _ := distro.NewZipf(rs, 1000, 0.693)

  for i := 0; i < 1000; i++ {
    fmt.Println(z.Uint64())
  }
}
```

乱数生成器を宣言するには`NewZipf`関数を実行します．`rand`パッケージで宣言された疑似乱数生成器`rs`を引数に持ちます．`n`に要素数，`alpha`に偏りパラメータを指定します．

```
z, err := distro.NewZipf(rs, n, alpha)
```

任意の順位の確率を知るには，`Pdf`関数を実行します．引数に順位を指定します．

```
p, err := z.Pdf(rank)
```

任意の順位までの累積確率を知るには，`Cdf`関数を実行します．引数に順位を指定します．

```
c, err := z.Cdf(rank)
```

`Uint64`関数は分布に基づいてランダムに順位を出力します．

```
k := z.Uint64()
```

## LICENSE

[MIT](./LICENSE)
