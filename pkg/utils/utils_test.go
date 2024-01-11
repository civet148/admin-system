package utils

import (
	"fmt"
	"github.com/civet148/log"
	"testing"
)

func TestNowRandom(t *testing.T) {
	fmt.Printf("time now random [%s]", NowRandom())
}

func TestStrMap2Uint(t *testing.T) {
	//strIn := "f2te32o3nkh7mxtsnts8nyum1pc6njygdm87sgurw7tywwxgfyl2wtb56xk6fwwrje567kaxcv456hygda3behda" //index=0
	//strIn := "f3te32o3nkh3mxtsntscnyumkpc6njygdmb7sgurw7tywwxgfyl2wtbmixk6fwwrjeufakaxcvwhygda3behda" //index=1
	strIn := "f3te32o3nk33mxtsntscnyumkpc6nj8gdmb7sgurw7tywwxgfyl2wtbmixk6fwwrjeufakaxcvwhygda3behdc" //index=2
	n := StrMap2Uint(strIn, 3)
	fmt.Printf("index = %d\n", n)
}

func TestUrlKey(t *testing.T) {
	router, key := UrlSuffix("/rpc/v1/xk6fwwrjeufakaxcvwhygda3behdc")
	log.Infof("router [%s] key [%s]", router, key)
}
