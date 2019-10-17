# poker-card

### v1.1.1 --todo
待处理：在现有算法情况下，做效率上的优化，主要有：
- 减少各方法中的重复遍历
- 仅考虑7张，增加分析判断，去掉许多种不可能的情况，以此优化算法
- 考虑使用协程（因为当前各方法间基本独立，可同时进行计算）（需要对比协程的开销）
- 改进现有算法选取判断部分
- **目前未做成评分的形式，做成评分更具备通用性，能同时支持N副手牌的比较（目前没想到合适的评分算法，同类型下可以根据5张牌以14进制的方案进行计算，但是不同类型件做通用的评分暂未想到）**

可能导致可复用性降低及增加`util`里方法间的耦合

### v1.1.0
支持包含N个癞子的N张牌，不仅仅7张，暂未做任何优化，模块具有比较强的可复用性，在i3-7100/8g/ubuntu上跑需要46-47ms

### v1.0.0
第一个版本，未做过多优化，仅支持5张无癞子，在i3-7100/8g/ubuntu上跑需要10-11ms


### 问题更新1
给定两手各7张牌，分别找出最大的牌型（选取其中5张），并比较大小
### 问题更新2
如果有一张牌是赖子（即可以当任何牌），怎么比较
### 问题
扑克牌52张，花色黑桃spades，红心hearts，方块diamonds，草花clubs各13张，2-10，J，Q，K，A

Face：即2-10，J，Q，K，A，其中10用T来表示。

Color：即S(spades)、H(hearts)、D(diamonds)、C(clubs)

用Face字母+小写Color字母表示一张牌，比如As表示黑桃A，其中A为牌面，s为spades，即黑桃，Ah即红心A，以此类推
现分别给定任意两手牌各有5张，例如：AsAhQsQhQc，vs KsKhKdKc2c，请按德州扑克的大小规则来判断双方大小。代码要求有通用性，可以任意输入一手牌或几手牌来进行比较。

**附：德州扑克的大小规则**
- 皇家同花顺
　　同花色的A, K, Q, J和10.
　　平手牌：在摊牌的时候有两副多副皇家同花顺时，平分筹码。
- 同花顺
　　五张同花色的连续牌。
　　平手牌：如果摊牌时有两副或多副同花顺，连续牌的头张牌大的获得筹码。如果是两副或多副相同的连续牌，平分筹码。
- 四条
　　其中四张是相同点数但不同花的扑克牌，第五张是随意的一张牌
　　平手牌：如果两组或者更多组摊牌，则四张牌中的最大者赢局，如果一组人持有的四张牌是一样的，那么第五张牌最大者赢局（起脚牌）。如果起脚牌也一样，平分彩池。
- 满堂彩（葫芦，三带二）
　　由三张相同点数及任何两张其他相同点数的扑克牌组成
　　平手牌：如果两组或者更多组摊牌，那么三张相同点数中较大者赢局。如果三张牌都一样，则两张牌中点数较大者赢局，如果所有的牌都一样，则平分彩池。
- 同花
　　此牌由五张不按顺序但相同花的扑克牌组成
　　平手牌：如果不止一人抓到此牌相，则牌点最高的人赢得该局，如果最大点相同，则由第二、第三、第四或者第五张牌来决定胜负，如果所有的牌都相同，平分彩池。
- 顺子
　　此牌由五张顺序扑克牌组成
　　平手牌：如果不止一人抓到此牌，则五张牌中点数最大的赢得此局，如果所有牌点数都相同，平分彩池。
- 三条
　　由三张相同点数和两张不同点数的扑克组成
　　平手牌：如果不止一人抓到此牌，则三张牌中最大点数者赢局，如果三张牌都相同，比较第四张牌，必要时比较第五张，点数大的人赢局。如果所有牌都相同，则平分彩池。
- 两对
　　两对点数相同但两两不同的扑克和随意的一张牌组成
　　平手牌：如果不止一人抓大此牌相，牌点比较大的人赢，如果比较大的牌点相同，那么较小
　　牌点中的较大者赢，如果两对牌点相同，那么第五张牌点较大者赢（起脚牌）。如果起脚牌也相同，
则平分彩池。
- 一对
　　由两张相同点数的扑克牌和另三张随意的牌组成
　　平手牌：如果不止一人抓到此牌，则两张牌中点数大的赢，如果对牌都一样，则比较另外三张牌
　　中大的赢，如果另外三张牌中较大的也一样则比较第二大的和第三大的，如果所有的牌都一样，
则平分彩池。
- 单张大牌
　　既不是同一花色也不是同一点数的五张牌组成。
　　平手牌：如果不止一人抓到此牌，则比较点数最大者，如果点数最大的相同，则比较第二、第三、第四和第五大的，如果所有牌都相同，则平分彩池。
