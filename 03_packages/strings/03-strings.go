package main

import (
	"fmt"
	"strings"
)

var collectDetailOld = ""

func main() {
	strJoin()

}

func strJoin() {
	collectDetailStr := "1,1,wx"
	if collectDetailOld != "" {
		collectDetailStr = strings.Join([]string{collectDetailOld, "wx1111111111111111" + ":" + collectDetailStr}, ";")
		fmt.Println("old detail存在的:", collectDetailStr)
	} else {
		collectDetailStr = "wx22222222222222" + ":" + collectDetailStr
		fmt.Println("old detail不存在的:", collectDetailStr)
	}

	collectDetailOld := collectDetailStr
	collectDetailStr = "2,2,zfb"

	if collectDetailOld != "" {
		collectDetailStr = strings.Join([]string{collectDetailOld, "zfb1111111111111111" + ":" + collectDetailStr}, ";")
		fmt.Println("old detail存在的:", collectDetailStr)
	} else {
		collectDetailStr = "zfb22222222222222" + ":" + collectDetailStr
		fmt.Println("old detail不存在的:", collectDetailStr)
	}

	
	strSplit2(collectDetailStr)
}

// 取全部的支付详情，未完成
func strSplit1(collectDetailStr string) {
	var channelDes string
	var payNo, collectDetail string

	collectDetailPpts := strings.Split(collectDetailStr, ";")
	for _, collectDetailPpt := range collectDetailPpts {
		collectInfo := strings.Split(collectDetailPpt, ":")
		if len(collectInfo) == 2 {
			payNo = collectInfo[0]
			collectDetail = collectInfo[1]
		}
		detailSli := strings.Split(collectDetail, ",")
		if len(detailSli) == 3 {
			if channelDes == "" {
				channelDes = detailSli[2]
			}
		}
	}
	fmt.Println(payNo, channelDes)
}

// 取最新的一次支付的详情
func strSplit2(collectDetailStr string) {
	var channelDes string
	var payNo, collectDetail string
	collectDetailPpts := strings.Split(collectDetailStr, ";")
	if len(collectDetailPpts) >= 1 {
		collectDetailPpt := collectDetailPpts[len(collectDetailPpts)-1] //拆单支付时暂只显示最新的支付记录
		collectInfo := strings.Split(collectDetailPpt, ":")
		if len(collectInfo) == 2 {
			payNo = collectInfo[0]
			collectDetail = collectInfo[1]
		}
		detailSli := strings.Split(collectDetail, ",")
		if len(detailSli) == 3 {
			if channelDes == "" {
				channelDes = detailSli[2]
			}
		}
	} else {

	}

	fmt.Println("取出的:", payNo, channelDes)
}
