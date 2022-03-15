/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/10/2022
    @UpdateDate: 3/10/2022
    @Note:       判断文件是否存在
**/

package utils

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
