## 题目

### 1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用

	1）fix length：消息统一满足固定长度，不足补零或者其他
	
	优点：简单
	
	缺点：浪费空间

	2）delimiter based：分隔符间隔消息
	
	优点：简单，不浪费空间
	
	缺点：内容本身出现分隔符需要转义,需要扫描全部内容

	3）length field based frame decoder：先解析固定长度的字段获取长度，然后读取后续内容
	
	优点：精确定位数据，内容不用转义
	
	缺点：长度理论上有限制，需提前考虑可能的最大长度从而定义长度占用字节

### 2. 实现一个从 socket connection 中解码出 goim 协议的解码器

