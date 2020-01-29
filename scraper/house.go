package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type HouseInfo struct {
	title      string
	unitPrice  int
	totalPrice int
	area       float64
	address    string
	detail     DetailInfo
}

type DetailInfo struct {
	houseType string
	area      string
	toward    string
	level     string
	floor     string
	buildTime string
	buildType string
}

//解析房子数据

func main() {
	log.Println("start..")

	c := colly.NewCollector(
		colly.Async(true),
		//模拟浏览器
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*.lianjia.*", Parallelism: 5})

	houses := make([]HouseInfo, 0, 200)
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("ul.sellListContent>li", func(element *colly.HTMLElement) {
		element.ForEach("div.info.clear", func(i int, element *colly.HTMLElement) {
			totalPrice, err := strconv.Atoi(element.DOM.Find(".priceInfo>.totalPrice>span").Text())
			if err != nil {
				panic(err)
			}

			unitPrice := getUnitPrice(element.DOM.Find(".priceInfo>.unitPrice>span").Text())

			areaFl := float64(totalPrice*10000) / float64(unitPrice)
			areaFloat, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", areaFl), 64)
			detailInfo := ParseHouseInfo(element.DOM.Find(".address>.houseInfo").Text())
			houseInfo := HouseInfo{
				title:      element.DOM.Find(".title").Text(),
				unitPrice:  unitPrice,
				totalPrice: totalPrice,
				area:       areaFloat,
				address:    element.DOM.Find(".flood>.positionInfo>a").Eq(0).Text(),
				detail:     detailInfo,
			}
			houses = append(houses, houseInfo)
			log.Println(houseInfo)
		})
	})

	c.OnResponse(func(response *colly.Response) {
		log.Println(response.StatusCode)
	})

	for i := 1; i < 8; i++ {
		url := "https://wh.lianjia.com/ershoufang/chuhehanjie/pg" + strconv.Itoa(i) + "/"
		go c.Visit(url)
		c.Wait()
	}
	//保存到csv里
	saveCsv(houses)
}

func saveCsv(houses []HouseInfo) {
	f, err := os.Create("house.csv")
	checkErr(err)
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()
	err = writer.Write([]string{"title", "均价", "总价", "面积", "地址", "户型", "area", "朝向", "装修", "楼层", "建造时间", "建筑材料"})
	checkErr(err)
	for _, info := range houses {
		detail := info.detail
		err = writer.Write([]string{
			info.title,
			strconv.Itoa(info.unitPrice),
			strconv.Itoa(info.totalPrice),
			strconv.FormatFloat(info.area, 'E', -1, 64),
			info.address,
			detail.houseType,
			detail.area,
			detail.toward,
			detail.level,
			detail.floor,
			detail.buildTime,
			detail.buildType,
		})
	}
}

//解析房子得属性
func ParseHouseInfo(houseInfo string) DetailInfo {
	var detailInfo DetailInfo
	infoArr := strings.Split(houseInfo, "|")
	infoLen := len(infoArr)
	if infoLen == 7 {
		detailInfo = DetailInfo{
			houseType: infoArr[0],
			area:      infoArr[1],
			toward:    infoArr[2],
			level:     infoArr[3],
			floor:     infoArr[4],
			buildTime: infoArr[5],
			buildType: infoArr[6],
		}
	} else if infoLen == 6 {
		detailInfo = DetailInfo{
			houseType: infoArr[0],
			area:      infoArr[1],
			toward:    infoArr[2],
			level:     infoArr[3],
			floor:     infoArr[4],
			buildTime: "无",
			buildType: infoArr[5],
		}
	} else {
		log.Println("houseInfo Parse Err:")
		log.Println(infoArr)
	}
	return detailInfo
}

func getUnitPrice(str string) int {
	var valid = regexp.MustCompile("[0-9]")
	res := valid.FindAllStringSubmatch(str, -1)
	var priceStr string
	for _, res := range res {
		priceStr += res[0]
	}
	unitPrice, err := strconv.Atoi(priceStr)
	checkErr(err)
	return unitPrice
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
