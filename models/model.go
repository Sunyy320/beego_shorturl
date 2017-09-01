package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"encoding/binary"
	"bytes"
	"math/rand"
	"time"
)
type response struct {
	Code string `json:"code"`
	Msg string `json:"message"`
	Data string `json:"data"`
}

// 62个字符, 需要6bit做索引(2 ^ 6 = 64)
var charTable = [...]rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k',
	'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
	'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6',
	'7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S',
	'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

// go中以大小写定义是否为公有还是私有方法
func GetMD5(lurl string) string {
	h := md5.New()
	salt1 := "salt4shorturl"
	io.WriteString(h, lurl+salt1)
	urlmd5 := fmt.Sprintf("%x", h.Sum(nil))
	return urlmd5
}

// 生成短地址
/**
比如 https://lengzzz.com 的 md5 值是 3CD4B16B4855CF2DBEB9051867665045 ，第一个字节是 0x3c ，十进制值为 60 那么在表中第 60 个是 Y ，那么用 Y 来表示 0x3c 。

但是这个思路还略有不妥，万一字节超了61就不行了。所以我们对字节使用 % 0x3d 取余运算，因为 0x3d 即为62，所以得到的数不会超出。

例如 0xd4 ，十进制为 212 显然超出了，但是 0xd4 % 0x3d 等于 26 ，第 26 个字符是 '0' ，所以可以用 '0' 来表示。

因此，我们可以把 md5 的输出分为 4 份，每份 4 个字节（byte，1个字节8位）。4 个字节又可以分成 6 份，每份 5 位（bit）共 30 位（bit），这样正好是在使用一个 uint32 计算，效率较高。
*/
func GenerateShortUrl(lurl string) string {
	// [90 16 94 139 157 64 225 50 151 128 214 46 162 38 93 138]
	urlmd5 := md5.Sum([]byte(lurl))// 返回的是对应的十进制
	var shortUrlList [] string

	// 分成4份，每一份4个字节
	for i := 0; i < 4; i++ {
		part:= urlmd5[i*4:i*4+4]
		// 把四个字节当做一个整数 1511022219 | 2638274866 | 2541803054 |2720423306
		partUnit:=binary.BigEndian.Uint32(part)

		shortURlBuffer :=&bytes.Buffer{}

		// 将30bit 分成6份，每份5bit
		for j:=0;j<6;j++{
			index:=partUnit % 62
			shortURlBuffer.WriteRune(charTable[index])
			partUnit=partUnit>>5
		}
		shortUrlList=append(shortUrlList,shortURlBuffer.String())
	}
	rand.Seed(time.Now().UnixNano())
	n:=rand.Intn(4)
	return shortUrlList[n]
}
