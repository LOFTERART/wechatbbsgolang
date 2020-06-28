package utils

import (
	"github.com/astaxie/beego/orm"
	"math/rand"
	"strconv"
	"time"
)


func RandStringRunes(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}

func String2Int(val string) int {

	goodsId_int, err := strconv.Atoi(val)
	if err != nil {
		return -1
	} else {
		return goodsId_int
	}
}

func String2Int64(val string) int64 {
	i, _ := strconv.ParseInt(val, 10, 64)

	return i
}

func Int2String(val int) string {
	return strconv.Itoa(val)
}

func Int642String(val int64) string {
	return strconv.FormatInt(val, 10)
}

func Float642String(val float64) string {
	return strconv.FormatFloat(val, 'E', -1, 64)
}

//func GetUUID() string {
//	uuid, err := uuid.NewV4()
//	if err != nil {
//		return ""
//	} else {
//		return uuid.String()
//	}
//}

//the result likes 1423361979
func GetTimestamp() int64 {
	return time.Now().Unix()
}

//the result likes 2015-02-08 10:19:39 AM
func FormatTimestamp(timestamp int64, format string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

func ExactMapValues2Int64Array(maparray []orm.Params, key string) []int64 {

	var vals []int64
	for _, value := range maparray {
		vals = append(vals, value[key].(int64))
	}
	return vals
}

func ExactMapValues2StringArray(maparray []orm.Params, key string) []string {

	var vals []string
	for _, value := range maparray {
		vals = append(vals, value[key].(string))
	}
	return vals
}

type PageData struct {
	NumsPerPage int         `json:"pageSize"`
	CurrentPage int         `json:"currentPage"`
	Count       int         `json:"count"`
	TotalPages  int         `json:"totalPages"`
	Data        interface{} `json:"data"`
}

func GetPageData(rawData []orm.Params, page int, size int) PageData {

	count := len(rawData)
	totalpages := (count + size - 1) / size
	var pagedata []orm.Params

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pagedata = append(pagedata, rawData[idx])
	}

	return PageData{NumsPerPage: size, CurrentPage: page, Count: count, TotalPages: totalpages, Data: pagedata}
}

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
