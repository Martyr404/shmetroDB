package main

import (
	"fmt"
	"shmetroDB/orm"

	"os"
)

func main22() {
	// 创建输出文件
	filename := "/home//testreport/report.txt"
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Printf("创建输出文件失败: %v\n", err)
		return
	}
	defer outputFile.Close()

	// 生成测试用例：从"050011"到"059999"
	start := 0   // 后四位起始数字
	end := 10000 // 后四位结束数字
	//prefix := "01" // 前缀

	for i := start; i <= end; i++ {
		// 格式化后四位为4位数字（补零）
		suffix := fmt.Sprintf("%05d", i)
		// 拼接完整测试字符串
		//testCase := prefix + suffix
		testCase := suffix

		// 调用待测试函数
		tList, err := orm.ParseCarriageNumber(testCase)

		// 处理tList可能为空的情况（避免索引越界）
		var firstElement interface{} = "<empty>"
		if len(tList) > 0 {
			firstElement = tList[0]
		}

		// 生成输出内容
		outputLine0 := fmt.Sprintf("test_case: %v,\n", testCase)
		outputLine1 := fmt.Sprintf("t_list:%v\n", firstElement)
		outputLine2 := fmt.Sprintf("Error:%v\n", err)

		// 写入文件
		if _, err := outputFile.WriteString(outputLine0); err != nil {
			fmt.Printf("写入文件失败（%s）: %v\n", testCase, err)
		}
		if _, err := outputFile.WriteString(outputLine1); err != nil {
			fmt.Printf("写入文件失败（%s）: %v\n", testCase, err)
		}
		if _, err := outputFile.WriteString(outputLine2); err != nil {
			fmt.Printf("写入文件失败（%s）: %v\n", testCase, err)
		}
	}

	fmt.Println("测试完成，结果已写入", filename)
}
