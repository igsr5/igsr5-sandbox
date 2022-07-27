# 1. 20個の標準正規乱数
x <- rnorm(20)

# 2.	1.で得られた標本は平均０、分散１の正規分布からの無作為抽出で得られたものか否かを、
# Ｒを用いてＺ検定(有意水準α=5%)せよ。

x <- rnorm(20)
qnorm(0.025)
qnorm(0.075)
z<-sqrt(20)*(mean(x)-0)/sqrt(1)
2*pnorm(abs(z), lower.tail=FALSE)

# 3.	1.で得られた標本は平均０の正規分布からの無作為抽出で得られたものか否かを、
# Ｒを用いてｔ検定(有意水準α=5%)せよ。
x <- rnorm(20)
qt(0.025, 19)
qt(0.975, 19)
t.test(x)
# NOTE: p値が0.05より大きいと棄却されない
