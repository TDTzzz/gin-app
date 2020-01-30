package main

import (
	"encoding/csv"
	"gin-app/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//解析房子数据

func main() {
	log.Println("start..")
	start := time.Now()

	c := colly.NewCollector(
		colly.Async(true),
		//模拟浏览器
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*.lianjia.*", Parallelism: 5})
	houses := make([]model.HouseInfo, 0, 400)
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("ul.sellListContent>li", func(element *colly.HTMLElement) {
		element.ForEach("div.info.clear", func(i int, element *colly.HTMLElement) {
			totalPrice, err := strconv.ParseFloat(element.DOM.Find(".priceInfo>.totalPrice>span").Text(), 64)
			checkErr(err)
			unitPrice := getUnitPrice(element.DOM.Find(".priceInfo>.unitPrice>span").Text())

			//areaFl := float64(totalPrice*10000) / float64(unitPrice)
			//areaFloat, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", areaFl), 64)
			detailInfo := ParseHouseDetail(element.DOM.Find(".address>.houseInfo").Text())
			titleElement := element.DOM.Find(".title>a")
			houseId, _ := titleElement.Attr("data-housecode")

			communityDom := element.DOM.Find(".flood>.positionInfo>a").Eq(0)
			community, communityId := getCommunity(communityDom)
			//community := element.DOM.Find(".flood>.positionInfo>a").Eq(0).Text()
			houseInfo := model.HouseInfo{
				Title:       titleElement.Text(),
				UnitPrice:   unitPrice,
				TotalPrice:  totalPrice,
				HouseDetail: detailInfo,
				HouseId:     houseId,
				Community:   community,
				CommunityId: communityId,
			}
			houses = append(houses, houseInfo)
			log.Println(houseInfo)
		})
	})

	c.OnResponse(func(response *colly.Response) {
		log.Println(response.StatusCode)
	})

	for i := 2; i < 3; i++ {
		//url := "https://wh.lianjia.com/ershoufang/chuhehanjie/pg" + strconv.Itoa(i) + "/"
		url := "https://wh.lianjia.com/ershoufang/jiedaokou/pg" + strconv.Itoa(i) + "/"
		c.Visit(url)
		c.Wait()
	}
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
	//保存到csv里
	saveCsv(houses)
}

func saveCsv(houses []model.HouseInfo) {
	f, err := os.Create("house3.csv")
	checkErr(err)
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()
	err = writer.Write([]string{"title", "房子ID", "均价", "总价", "小区", "小区ID", "户型", "area", "朝向", "装修", "楼层", "建造时间", "建筑材料"})
	checkErr(err)
	for _, info := range houses {
		detail := info.HouseDetail
		err = writer.Write([]string{
			info.Title,
			info.HouseId,
			strconv.Itoa(int(info.UnitPrice)),
			strconv.FormatFloat(info.TotalPrice, 'E', -1, 64),
			info.Community,
			info.CommunityId,
			detail.HouseType,
			strconv.FormatFloat(detail.Area, 'E', -1, 64),
			detail.Toward,
			detail.Level,
			detail.Floor,
			strconv.Itoa(int(detail.BuildYear)),
			detail.BuildType,
		})
	}
}

//解析房子得属性
func ParseHouseDetail(houseInfo string) model.HouseDetail {
	var detailInfo model.HouseDetail
	infoArr := strings.Split(houseInfo, "|")
	infoLen := len(infoArr)

	var valid = regexp.MustCompile("[0-9]|\\.")
	res := valid.FindAllString(infoArr[1], -1)
	areaStr := strings.Join(res, "")
	area, _ := strconv.ParseFloat(areaStr, 64)

	res = valid.FindAllString(infoArr[5], -1)
	buildStr := strings.Join(res, "")
	buildYear, _ := strconv.ParseUint(buildStr, 10, 64)

	if infoLen == 7 {
		detailInfo = model.HouseDetail{
			HouseType: infoArr[0],
			Area:      area,
			Toward:    infoArr[2],
			Level:     infoArr[3],
			Floor:     infoArr[4],
			BuildYear: uint(buildYear),
			BuildType: infoArr[6],
		}
	} else if infoLen == 6 {
		detailInfo = model.HouseDetail{
			HouseType: infoArr[0],
			Area:      area,
			Toward:    infoArr[2],
			Level:     infoArr[3],
			Floor:     infoArr[4],
			//BuildYear: "",
			BuildType: infoArr[5],
		}
	} else {
		log.Println("houseInfo Parse Err:")
		log.Println(infoArr)
	}
	return detailInfo
}

func getUnitPrice(str string) uint {
	var valid = regexp.MustCompile("[0-9]")
	res := valid.FindAllStringSubmatch(str, -1)
	var priceStr string
	for _, res := range res {
		priceStr += res[0]
	}
	unitPrice, err := strconv.Atoi(priceStr)
	checkErr(err)
	return uint(unitPrice)
}

//提取小区信息
func getCommunity(element *goquery.Selection) (string, string) {
	community := element.Text()
	communityUrl, _ := element.Attr("href")
	var valid = regexp.MustCompile("[0-9]")
	res := valid.FindAllString(communityUrl, -1)
	communityId := strings.Join(res, "")
	return community, communityId
}

func checkErr(err error) {
	if err != nil {
		log.Println("checkErr")
		panic(err)
	}
}
