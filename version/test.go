package main

import (
	"fmt"
)

type UUC struct {
	total    int64
	upload   int64
	download int64
}

func splitUUC(original UUC, limitValue int64) []UUC {
	var splittedUUCs []UUC

	remainingTotal := original.total
	remainingUpload := original.upload
	remainingDownload := original.download

	for remainingTotal > limitValue {
		newUUC := UUC{total: 0, upload: 0, download: 0}

		// 先从 upload 字段分配
		if remainingUpload > (limitValue - newUUC.total) {
			newUUC.upload = limitValue - newUUC.total
			remainingUpload -= newUUC.upload
		} else {
			newUUC.upload = remainingUpload
			remainingUpload = 0
		}
		newUUC.total += newUUC.upload

		// 然后从 download 字段分配
		if remainingDownload > (limitValue - newUUC.total) {
			newUUC.download = limitValue - newUUC.total
			remainingDownload -= newUUC.download
		} else {
			newUUC.download = remainingDownload
			remainingDownload = 0
		}
		newUUC.total += newUUC.download

		// 将新的 UUC 添加到列表中
		splittedUUCs = append(splittedUUCs, newUUC)

		// 更新剩余的 total 字段
		remainingTotal -= limitValue
	}

	// 处理剩余的部分（如果有的话）
	if remainingTotal > 0 {
		splittedUUCs = append(splittedUUCs, UUC{total: remainingTotal, upload: remainingUpload, download: remainingDownload})
	}

	return splittedUUCs
}

func main() {
	originalUUC := UUC{total: 14000, upload: 8600, download: 5400}
	limitValue := 3000

	splittedUUCs := splitUUC(originalUUC, int64(limitValue))
	for i, uuc := range splittedUUCs {
		fmt.Printf("Splitted UUC %d: %+v\n", i+1, uuc)
	}
}
