package main

import (
	"awesomeProject/docs"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type User struct {
	Uid               int            `json:"uid"`
	Name              string         `json:"name"`
	Username          string         `json:"username"`
	Pubkey            string         `json:"pubkey"`
	Bio               string         `json:"bio"`
	Avatar            string         `json:"avatar"`
	Background        string         `json:"background"`
	Email             string         `json:"email"`
	Facebook          string         `json:"facebook"`
	Instagram         string         `json:"instagram"`
	Twitter           string         `json:"twitter"`
	Youtube           string         `json:"youtube"`
	CountryCode       string         `json:"country_code"`
	Language          string         `json:"language"`
	BlockchainAddress string         `json:"blockchain_address"`
	Vip               int            `json:"vip"`
	TypeUser          int            `json:"type_user"`
	CanMint           bool           `json:"can_mint"`
	Verify            bool           `json:"verify"`
	IsPartner         bool           `json:"is_partner"`
	App               string         `json:"app"`
	TotalFollowing    int            `json:"total_following"`
	TotalFollowers    int            `json:"total_followers"`
	ExtendData        map[string]int `json:"extend_data"`
	Block             bool           `json:"block"`
	Deleted           bool           `json:"deleted"`
	CreateTime        int            `json:"create_time"`
	UpdateTime        int            `json:"update_time"`
}

type Item struct {
	Id              int            `json:"id"`
	ImageUrl        string         `json:"image_url"`
	Url             string         `json:"url"`
	Name            string         `json:"name"`
	Status          int            `json:"status"`
	CollectionId    int            `json:"collection_id"`
	CreateTime      int            `json:"create_time"`
	UpdateTime      int            `json:"update_time"`
	ArrOwnerAddress map[string]int `json:"arr_owner_address"`
	ArrOwnerUid     map[string]int `json:"arr_owner_uid"`
	OwnerAddress    string         `json:"owner_address"`
	OwnerUid        int            `json:"owner_uid"`
	CreatorAddress  string         `json:"creator_address"`
	CreatorUid      int            `json:"creator_uid"`
	ExternalLink    string         `json:"external_link"`
	EventIds        []int          `json:"event_ids"`
	TotalView       int            `json:"total_view"`
	TotalLike       int            `json:"total_like"`
	Edition         int            `json:"edition"`
	MetaData        string         `json:"meta_data"`
	MediaType       int            `json:"media_type"`
	CategoryType    int            `json:"category_type"`
	ProductNo       string         `json:"product_no"`
	TotalEdition    int            `json:"total_edition"`
	Uri             string         `json:"uri"`
	ExtendData      struct {
		Category   string `json:"category"`
		Collection string `json:"collection"`
		Create     string `json:"create"`
		Like       string `json:"like"`
		Owner      string `json:"owner"`
		TotalOffer string `json:"total_offer"`
	} `json:"extend_data"`
	IsSensitive     bool   `json:"is_sensitive"`
	Hidden          bool   `json:"hidden"`
	Price           string `json:"price"`
	TokenId         string `json:"token_id"`
	ContractAddress string `json:"contract_address"`
	Chain           string `json:"chain"`
	CurrencyAddress struct {
		Name    string `json:"name"`
		Symbol  string `json:"symbol"`
		Address string `json:"address"`
		Chain   string `json:"chain"`
		Decimal int    `json:"decimal"`
		Img     string `json:"img"`
		Status  bool   `json:"status"`
	} `json:"currency_address"`
	TypeNft           int    `json:"type_nft"`
	TotalClaim        int    `json:"total_claim"`
	TotalLimit        int    `json:"total_limit"`
	StartAt           int    `json:"start_at"`
	EndAt             int    `json:"end_at"`
	UnlockableContent string `json:"unlockable_content"`
	IsLootBox         bool   `json:"is_loot_box"`
}

var items = []Item{}
var users = []User{}

var ctx = context.Background()
var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func main() {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	r.POST("/item", newItems)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/item", getItems)
	r.DELETE("/item/:id", deleteItems)
	r.PUT("/item/:id", editItems)
	r.POST("/user", addUser)
	r.GET("/item/:uid", GetItemByUserID)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags Item
// @Accept json
// @Produce json
// @Param pos query int64 false "position"
// @Param count query int64 false "count"
// @Success 200 {object} Item
// @Router /item [get]
func getItems(c *gin.Context) {
	var items []Item

	pos, _ := c.GetQuery("pos")
	cou, _ := c.GetQuery("count")
	position, _ := strconv.Atoi(pos)
	count, _ := strconv.Atoi(cou)
	val, _ := client.ZRange(ctx, "item", int64(position), int64(position+count)).Result()
	for _, a := range val {
		var item Item
		json.Unmarshal([]byte(a), &item)
		items = append(items, item)
	}

	c.IndentedJSON(http.StatusOK, items)

	/*		items[i] = item*/
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags Item
// @Accept json
// @Param NewItem body Item true "Thông tin item mới "
// @Produce json
// @Success 200 {object} Item
// @Router /item [post]
func newItems(c *gin.Context) {
	var NewItem Item
	if err := c.BindJSON(&NewItem); err != nil {
		return
	}
	val, err := client.ZRangeByScore(ctx, "item", &redis.ZRangeBy{
		Min:    strconv.Itoa(NewItem.Id),
		Max:    strconv.Itoa(NewItem.Id),
		Offset: 0,
		Count:  2,
	}).Result()
	if err != nil {
		logrus.Errorln(err)
		return
	}
	if len(val) != 0 {
		logrus.Errorln("Đã tồn tại id này")
		return
	}

	bData, _ := json.Marshal(NewItem)
	_, err = client.ZAdd(ctx, "item", redis.Z{Score: float64(NewItem.Id), Member: bData}).Result()
	if err != nil {
		logrus.Errorln(err)
	}
	c.IndentedJSON(http.StatusCreated, items)

}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags Item
// @Accept json
// @Param id path int true "Id cần xóa"
// @Produce json
// @Success 200 {object} Item
// @Router /item/{id} [delete]
func deleteItems(c *gin.Context) {
	id := c.Param("id")
	val, err := client.ZRangeByScore(ctx, "item", &redis.ZRangeBy{
		Min:    id,
		Max:    id,
		Offset: 0,
		Count:  2,
	}).Result()
	if err != nil {
		logrus.Errorln(err)
		return
	}
	if len(val) != 0 {
		_, err = client.ZRem(ctx, "item", val).Result()
		if err != nil {
			log.Fatalf("Error delete %s", val)
		}
		c.IndentedJSON(http.StatusOK, items)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item does not exist"})
	return
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags Item
// @Accept json
// @Param id path int true "ID cần sửa"
// @Param ItemEdit body Item true "Thông tin cần sửa"
// @Produce json
// @Success 200 {object} Item
// @Router /item/{id} [put]
func editItems(c *gin.Context) {
	id := c.Param("id")
	var ItemEdit Item
	if err := c.BindJSON(&ItemEdit); err != nil {
		return
	}
	val, err := client.ZRangeByScore(ctx, "item", &redis.ZRangeBy{
		Min:    id,
		Max:    id,
		Offset: 0,
		Count:  2,
	}).Result()
	logrus.Errorln(id)
	logrus.Errorln(val)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	if len(val) != 0 {
		_, err := client.ZRem(ctx, "item", val).Result()
		if err != nil {
			logrus.Errorln("Error delete %s")
			return
		}
		bData, _ := json.Marshal(ItemEdit)
		_, err = client.ZAdd(ctx, "item", redis.Z{Score: float64(ItemEdit.Id), Member: bData}).Result()
		if err != nil {
			logrus.Errorln("Error adding ", err)
			return
		}
		c.IndentedJSON(http.StatusOK, items)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item does not exist"})
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags User
// @Accept json
// @Param NewUser body User true "Thông tin user mới "
// @Produce json
// @Success 200 {object} User
// @Router /user [post]
func addUser(c *gin.Context) {
	var NewUser User
	if err := c.BindJSON(&NewUser); err != nil {
		return
	}
	val, err := client.ZRangeByScore(ctx, "user", &redis.ZRangeBy{
		Min:    strconv.Itoa(NewUser.Uid),
		Max:    strconv.Itoa(NewUser.Uid),
		Offset: 0,
		Count:  2,
	}).Result()
	if err != nil {
		logrus.Errorln(err)
		return
	}
	if len(val) != 0 {
		logrus.Errorln("Đã tồn tại id này")
		return
	}

	bData, _ := json.Marshal(NewUser)
	_, err = client.ZAdd(ctx, "user", redis.Z{Score: float64(NewUser.Uid), Member: bData}).Result()
	if err != nil {
		logrus.Errorln(err)
	}
	c.IndentedJSON(http.StatusCreated, items)

}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags Item
// @Accept json
// @Param uid path int true "ID của người dùng cần tìm kiếm "
// @Param pos query int64 false "position"
// @Param count query int64 false "count"
// @Produce json
// @Success 200 {object} Item
// @Router /item/{uid} [get]
func GetItemByUserID(c *gin.Context) {
	var items []Item
	uid, _ := strconv.Atoi(c.Param("uid"))
	val, _ := client.ZRange(ctx, "item", 0, -1).Result()
	for _, a := range val {
		var item Item
		json.Unmarshal([]byte(a), &item)
		items = append(items, item)
	}
	for _, a := range items {
		if a.OwnerUid == uid {
			dData, _ := json.Marshal(a)
			log.Println("main.go:346")
			_, err := client.ZAdd(ctx, strconv.Itoa(uid)+"_itembyUID", redis.Z{Score: float64(a.Id), Member: dData}).Result()
			if err != nil {
				logrus.Errorln(err)
			}
		} else {
			logrus.Println("UID này không sở hữu item nào")
		}
	}
	pos, _ := c.GetQuery("pos")
	cou, _ := c.GetQuery("count")
	position, _ := strconv.Atoi(pos)
	count, _ := strconv.Atoi(cou)
	val, _ = client.ZRange(ctx, strconv.Itoa(uid)+"_itembyUID", int64(position), int64(position+count)).Result()
	if len(val) != 0 {
		var itembyID []Item
		for _, a := range val {
			var item Item
			json.Unmarshal([]byte(a), &item)
			itembyID = append(itembyID, item)
		}
		c.IndentedJSON(http.StatusOK, itembyID)
	} else {
		logrus.Println("UID này không sở hữu item nào")
	}

}
