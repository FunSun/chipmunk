# 题解

1. 游戏玩家的初始⽤户名是随机的，遵从以下格式：
两个⼤写字母后⾯跟着四个数字，⽐如 RX8372 或
者 BC8115 。随机意味着，⽤户名不能符合某个可
以预测的序列。此外，随机⽤户名意味着存在冲突
风险，如有可能，解决⽅案中应避免重复使⽤⽤户
名。请按以上要求编写代码。

A: 根据通用的LCG随机数生成方法(https://en.wikipedia.org/wiki/Linear_congruential_generator)在(26*26*10000)中不重复的生成一个整形, 然后取尾部4位作为用户名的数字部分, 头部/26作为第一个字幕, 头部%26作为第二个字母

2. 请为以上程序编写测试代码。

A: 见lcg.go usergen_test.go

3. 请估算程序⽣成 10,000 个⽤户名所需要的时间，并
给出分析过程。

A: 因为生成随机数和将随机数转换成用户名的时间是固定的, 所以每次操作的时间复杂度是O(1), 令单次时间为T
生成10000个用户的时间为10000T

4. 如需在多个进程中，独⽴⽣成符 合以上规则的、不
重复的随机⽤户名，应该如何调整算法

A: 这里的关键点在于对lcg中n的访问是原子的, 所以lcg.Next中n的访问使用了原子操作, 所以在多个goroutine
中调用UsernameGenerator.Generate() 都是并发安全的
