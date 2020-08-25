package Library

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

//
//func writeToCsv(file os.File, msg string) {
//	writer := csv.NewWriter(&file)
//	defer writer.Flush()
//
//	record := []string{time.Now().String(),msg }
//	fmt.Println(record)
//	writer.Write(record)
//}
func writeToCsv(file os.File, msg string) {
	writer := csv.NewWriter(&file)
	defer writer.Flush()
	record := []string{time.Now().String(), msg}
	fmt.Println(record)
	writer.Write(record)

}
