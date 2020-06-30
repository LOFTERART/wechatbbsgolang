package serializer

import (
	"QUZHIYOU/models"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"sort"
)

//序列化1
type Community struct {
	Id      uint   `json:"id"`
	Name    string `json:"cityName"`
	KeyWord string `json:"keyword"`
	Letter  string `json:"letter"`
}

//返回结果序列化2
type ResComs struct {
	Letter string      `json:"letter"`
	Data   []Community `json:"data"`
}

type Lists struct {
	List []ResComs `json:"list"`
}

func BuildCommunity(item models.Communitys) Community {
	return Community{
		Id:      item.ID,
		Name:    item.Name,
		KeyWord: item.KeyWord,
		Letter:  item.Letter,
	}
}

type Animals []models.Communitys

//noinspection ALL
func (a Animals) Len() int { return len(a) }

//noinspection ALL
func (s Animals) Less(i, j int) bool {
	a, _ := UTF82GBK(s[i].Letter)
	b, _ := UTF82GBK(s[j].Letter)
	bLen := len(b)
	for idx, chr := range a {
		if idx > bLen-1 {
			return false
		}
		if chr != b[idx] {
			return chr < b[idx]
		}
	}
	return true
}

//noinspection ALL
func (a Animals) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//UTF82GBK : transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

//GBK2UTF8 : transform  GBK byte array into UTF8 string
func GBK2UTF8(src []byte) (string, error) {
	GB18030 := simplifiedchinese.All[0]
	bytes, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
	return string(bytes), err
}

func BuildCommunitys(item []models.Communitys) (list Lists) {

	//强转类型
	sort.Sort(Animals(item))

	//格式化后数据
	//var cityList []Community
	//for _, v := range item {
	//	cityList = append(cityList, BuildCommunity(v))
	//}

	//二维数组
	var arrays []ResComs
	//一维数组
	var temp ResComs

	item = append(item, models.Communitys{})

	for i := 0; i < len(item)-1; i++ {

		if item[i].Letter == item[i+1].Letter {
			temp.Letter = item[i].Letter
			temp.Data = append(temp.Data, BuildCommunity(item[i]))
		} else {
			temp.Letter = item[i].Letter
			temp.Data = append(temp.Data, BuildCommunity(item[i]))
			arrays = append(arrays, temp)
			//	置空temp
			temp.Letter = ""
			temp.Data = make([]Community, 0)

		}

	}

	list.List = arrays

	temp.Data = nil

	//待封装数据
	//res := ResComs{}
	//
	//for i := 0; i < len(cityList); i++ {
	//	//取出第一个数组
	//	item := cityList[i]
	//
	//	var tempList []Community
	//	tempList = append(tempList, item)
	//
	//	//循环比较余下的数据
	//	for j := i + 1; j < len(cityList); j++ {
	//		temp := cityList[j]
	//		//如果相等就扔进tempList []
	//		if temp.Letter == item.Letter {
	//			tempList = append(tempList, temp)
	//			//根据下标删除已经加入的元素
	//			cityList = append(cityList[:j], cityList[j+1:]...)
	//			//移后一位避免跳过当前位置
	//			j = j - 1
	//		}
	//	}
	//
	//	//外层循环赋值 一个letter 和多个[]Community
	//	res.Letter = item.Letter
	//	res.Data = tempList
	//	//每次循环加入待返回数组
	//	list.List = append(list.List, res)
	//
	//}

	return

}
