# beego_shorturl
beego实现短域名服务
### 实现思路
**存储**：我们在这里直接使用memory，使用内存。用于存储短地址和长地址的对应关系

**访问short**:我们先把获取的longurl参数加密作为key,去memory取相关的地址，
如果不为nil则返回shortrul.

如果返回nil，则还未生成短地址。我们使用算法生成短地址。存入memory,
key为加密后的longurl,value为生成的短地址并设置过期时间

**访问extend**
### 短地址生成简单算法
1. 将长地址md5生成32位签名串，分成4段，每段8个字节
2. 对四段进行循环处理，取8个字节，将它看成16进制串与0x3fffffff(30位1)
与操作，即超过30位的忽略处理
3. 将30位分为6段，每5位的数字作为字母表的索引取得特定的字符，依次进行
获得6位字符串
4. 总的md5串可以获取4个6位串，任取一个就可以作为这个长url的短url地址

这种算法虽然会生成4个但是依然存在重复的几率，本项目只是练手使用
