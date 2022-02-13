package Helper

import (
	"os"
)

// 文件追加
func AppendToFile(file, str string) bool {
	// 参数是否合法
	if len(file) == 0 || len(str) == 0 {
		return false
	}
	// 打开文件 不存在创建  文件追加 文件读写
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	// 是否有错误
	if err != nil {
		return false
	}
	// 记得关闭文件
	defer f.Close()
	// 写数据
	f.WriteString(str)
	return true
}

// 文件夹不存在则创建
func DirCreateByNotExsit(dirname string) bool {
	if len(dirname) == 0 {
		return false
	}
	_, err := os.Stat(dirname)
	if err != nil {
		if os.IsNotExist(err) {
			err2 := os.Mkdir(dirname, os.ModePerm)
			if err2 != nil {
				return false
			}
		}
	}
	return true
}
