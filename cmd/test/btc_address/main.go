// 创建地址
package main

import (
	"go-dc-wallet/app"
	"go-dc-wallet/hbtc"
)

func main() {
	app.EnvCreate()
	defer app.EnvDestroy()

	hbtc.CheckAddressFree()
}
