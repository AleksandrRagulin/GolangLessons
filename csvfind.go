/*Поиск файла в заданном формате и его обработка
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath, хотя для решения может быть использован также пакет archive/zip (поскольку файл с заданием предоставляется именно в этом формате).
В тестовом архиве, который вы можете скачать из нашего репозитория на github.com, содержится набор папок и файлов. Один из этих файлов является файлом с данными в формате CSV, прочие же файлы структурированных данных не содержат.
Требуется найти и прочитать этот единственный файл со структурированными данными (это таблица 10х10, разделителем является запятая), а в качестве ответа необходимо указать число, находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).*/
// https://stepik.org/lesson/351892/step/13?unit=335849
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func mywalkFunc(path string, info os.FileInfo, err error) error {
	tmpStr := path
	file, err := os.Open(tmpStr)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if strings.Contains(sc.Text(), ",") && strings.Contains(path, ".txt") {
			fmt.Println(path)
			fmt.Println(sc.Text())
		}
	}
	return nil
}
func main() {
	const root = "."
	if err := filepath.Walk(root, mywalkFunc); err != nil {
		fmt.Printf("ошибка: %v ", err)
	}
}
