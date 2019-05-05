package v1

import (
	cs "Gin_demo/pkg/constant"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
)

var png_arrs = [...]string{"Cake Costume Cream.png", "Elementary School Costume.png", "Dress Costume Brown.png", "pajamas-costume.png", "Winter Coat Costume White.png", "Whiteday Costume Red.png", "Sarori Costume.png", "Winter Fairy Costume Pink.png", "Sakura Fairy Costume.png", "Santa 2018 Costume Red.png", "Santa 2018 Costume Green.png", "Sinsiroad Shop Costume Senior.png", "School 2019 Costume Black.png", "Magical Girl Costume Pink.png", "Voice Story Costume.png", "Star Witch Costume.png", "Ribbon Dress Costume Yellow.png", "New2015 Costume Pajamas.png", "Witch Costume.png", "Cake Costume Choco.png", "Forest Witch Costume Green.png", "Qipao Costume Red.png", "Frill Blouse Costume Green.png", "Shaman Costume Black.png", "Overalls Costume White.png", "Furisode Costume.png", "Sorceress Costume.png", "Literature Girl Costume Brown.png", "Marine Costume White.png", "Sinsiroad Costume.png", "Winter Fairy Costume Black.png", "Summer Dress Costume White.png", "School 2017 Costume Yellow.png", "School 2017 Costume Gray.png", "Pajamas Costume Pink.png", "Valentine Costume Pink.png", "Overalls Costume.png", "Swimsuit 2017 Costume Red.png", "Summer Uniform Costume Blue.png", "Turtleneck Costume Red.png", "Winter Coat Costume Pink.png", "Ribbon Dress Costume Red.png", "Succubus Costume Red.png", "Priest Costume Junior.png", "School 2019 Costume Pink.png", "Sukumizu Costume.png", "Star Witch Costume Brown.png", "Goddess Costume Pink.png", "Frill Bikini Costume Purple.png", "Marine Costume Navy.png", "Fall Dress Costume Beige.png", "Hanbok Costume Pink.png", "Hanbok Costume Red.png", "Bunny Girl Costume Red.png", "Forest Witch Costume Brown.png", "Healer Costume.png", "Vampire Costume Real.png", "Sailor Costume Black.png", "Turtleneck Costume.png", "Akiba Idol Costume.png", "Sakura Fairy Costume Real.png", "Santa Costume Green.png", "Sinsiroad Shop Costume Junior.png", "SFC Uniform Costume Red.png", "Dress Costume.png", "Sporty Hood Costume Blue.png", "Winter Coat 2017 Costume White.png", "Literature Girl Costume Navy.png", "Lolita Costume Skyblue.png", "Sporty Hood Costume Black.png", "Kids Costume Navy.png", "Magical Girl Costume Purple.png", "Winter Costume.png", "Maid Costume.png", "Vampire Costume.png", "Fall Dress Costume Brown.png", "Tirami1 Costume.png", "Succubus Costume Black.png", "Sukumizu Costume White.png", "Sakura Costume.png", "Maid Costume Red.png", "Valentine Costume Brown.png", "Hanbok Costume Yellow.png", "Witch Costume White.png", "Hanbok Costume Skyblue.png", "Winter Costume White.png", "Santa Costume.png", "Party Dress Costume Purple.png", "Frill Bikini Costume Green.png", "Whiteday Costume Purple.png", "Party Dress Costume Brown.png", "Qipao Costume Pink.png", "Hanbok Costume.png", "Summer Dress Costume Blue.png", "Winter Coat 2017 Costume Brown.png", "Sailor Costume.png", "Kids Costume.png", "School Costume Red.png", "Nightsky Costume.png", "Elementary School Costume Navy.png", "default-costume.png", "Sakura Costume Navy.png", "Swimsuit 2017 Costume Navy.png", "Animal Costume Racoon.png", "SFC Uniform Costume Yellow.png", "New2015 Costume.png", "Frill Blouse Costume Red.png", "Shaman Costume Blue.png", "school-costume.png", "Summer Uniform Costume Red.png", "Animal Costume.png", "Night Witch Costume Black.png", "Lolita Costume Red.png", "Goddess Costume White.png", "Bunny Girl Costume.png", "Night Witch Costume Gray.png", "Witch Costume Special.png", "Priest Costume Senior.png", "Halloween Costume.png"}

// @Summary 查询换装数据
// @Produce  json
// @Param id path int true "ID"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/model/ [get]
func GetModelJson(c *gin.Context) {
	// data, err := ioutil.ReadFile("/Users/wushaoqiang/Downloads/live2DDDD/Pio/model.json")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// jsonData := string(data)
	//获取cookie
	// cookie, err := c.Cookie("pio")
	// fmt.Println("cookie为:", cookie)
	// if err != nil || cookie == "" {
	// 	cookie = "0"
	// } else {
	// 	s, _ := strconv.Atoi(cookie)
	// 	cookie = string(s + 1)
	// }
	// c.SetCookie("pio", cookie, 3600, "/", "localhost", false, true)
	i := rand.Intn(120)
	pngName := "textures/" + png_arrs[i]
	r := "/Users/wushaoqiang/go/src/Gin_demo/static/live2D/model/Pio/model.json"
	err := HandleJson(r, pngName)
	if err != nil {
		c.JSON(cs.SUCCESS, gin.H{
			"msg": "SUCCESS",
		})
	}
}

func HandleJson(jsonFile string, png string) error {
	// Read json buffer from jsonFile
	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	// We have known the outer json object is a map, so we define result as map.
	// otherwise, result could be defined as slice if outer is an array
	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return err
	}

	// handle peers
	nodes := result["textures"].([]interface{})
	nodes[0] = png

	// Convert golang object back to byte
	byteValue, err = json.Marshal(result)
	if err != nil {
		return err
	}
	// Write back to file
	f, err := os.OpenFile(jsonFile, os.O_WRONLY|os.O_TRUNC, 0600)

	defer f.Close()

	_, err1 := f.WriteString(string(byteValue))

	return err1
}
