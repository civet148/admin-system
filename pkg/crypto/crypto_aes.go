package crypto

import (
	"github.com/civet148/gotools/cryptos/goaes"
	_ "github.com/civet148/gotools/cryptos/goaes/cbc" //注册CBC加解密对象创建方法
	_ "github.com/civet148/gotools/cryptos/goaes/cfb" //注册CFB加解密对象创建方法
	_ "github.com/civet148/gotools/cryptos/goaes/ctr" //注册CTR加解密对象创建方法
	_ "github.com/civet148/gotools/cryptos/goaes/ecb" //注册ECB加解密对象创建方法
	_ "github.com/civet148/gotools/cryptos/goaes/ofb" //注册OFB加解密对象创建方法
	"github.com/civet148/log"
)

var DefaultKEY = []byte("c6vgru6d9ic6gu563cyoegnzdq0klvx4") //加密KEY(16/24/32字节)
var DefaultIV = []byte("1de1c8c41007a070")                  //加密向量(16字节)

type CryptoAES struct {
	aes goaes.CryptoAES
}

func NewCryptoAES(key, iv []byte) *CryptoAES {
	return &CryptoAES{
		aes: goaes.NewCryptoAES(goaes.AES_Mode_CBC, key, iv),
	}
}

func NewCryptoAESDefault() *CryptoAES {
	return NewCryptoAES(DefaultKEY, DefaultIV)
}

func (c *CryptoAES) EncryptBase64(in []byte) (string, error) {
	enc, err := c.aes.EncryptBase64(in)
	if err != nil {
		return "", log.Errorf("[%v] encrypt to base64 error [%v]", c.aes.GetMode(), err.Error())
	}
	return enc, nil
}

func (c *CryptoAES) DecryptBase64(in string) ([]byte, error) {
	dec, err := c.aes.DecryptBase64(in)
	if err != nil {
		return nil, log.Errorf("[%v] decrypt from base64 [%s] error [%v]", c.aes.GetMode(), in, err.Error())
	}
	return dec, nil
}
