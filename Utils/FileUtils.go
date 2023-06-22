package Utils

import "io/ioutil"

func OpenFile(fileName string) ([]byte, error) {
        // 1. 打开文件并读取其内容到内存中
        content, err := ioutil.ReadFile(fileName)
        if err != nil {
            return nil,err
        }

            return content,nil
}
