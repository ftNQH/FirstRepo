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
	Uid               int    `json:"uid"`
	Name              string `json:"name"`
	Username          string `json:"username"`
	Pubkey            string `json:"pubkey"`
	Bio               string `json:"bio"`
	Avatar            string `json:"avatar"`
	Background        string `json:"background"`
	Email             string `json:"email"`
	Facebook          string `json:"facebook"`
	Instagram         string `json:"instagram"`
	Twitter           string `json:"twitter"`
	Youtube           string `json:"youtube"`
	CountryCode       string `json:"country_code"`
	Language          string `json:"language"`
	BlockchainAddress string `json:"blockchain_address"`
	Vip               int    `json:"vip"`
	TypeUser          int    `json:"type_user"`
	CanMint           bool   `json:"can_mint"`
	Verify            bool   `json:"verify"`
	IsPartner         bool   `json:"is_partner"`
	App               string `json:"app"`
	TotalFollowing    int    `json:"total_following"`
	TotalFollowers    int    `json:"total_followers"`
	ExtendData        struct {
	} `json:"extend_data"`
	Block      bool `json:"block"`
	Deleted    bool `json:"deleted"`
	CreateTime int  `json:"create_time"`
	UpdateTime int  `json:"update_time"`
}

type Item struct {
	Id              int    `json:"id"`
	ImageUrl        string `json:"image_url"`
	Url             string `json:"url"`
	Name            string `json:"name"`
	Status          int    `json:"status"`
	CollectionId    int    `json:"collection_id"`
	CreateTime      int    `json:"create_time"`
	UpdateTime      int    `json:"update_time"`
	ArrOwnerAddress struct {
		Xce85A994E411901Fc9Bda228D07F7B66F8320Ecd int `json:"0xce85a994e411901fc9bda228d07f7b66f8320ecd"`
	} `json:"arr_owner_address"`
	ArrOwnerUid struct {
		Field1 int `json:"221139"`
	} `json:"arr_owner_uid"`
	OwnerAddress   string `json:"owner_address"`
	OwnerUid       int    `json:"owner_uid"`
	CreatorAddress string `json:"creator_address"`
	CreatorUid     int    `json:"creator_uid"`
	ExternalLink   string `json:"external_link"`
	EventIds       []int  `json:"event_ids"`
	TotalView      int    `json:"total_view"`
	TotalLike      int    `json:"total_like"`
	Edition        int    `json:"edition"`
	MetaData       string `json:"meta_data"`
	MediaType      int    `json:"media_type"`
	CategoryType   int    `json:"category_type"`
	ProductNo      string `json:"product_no"`
	TotalEdition   int    `json:"total_edition"`
	Uri            string `json:"uri"`
	ExtendData     struct {
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

var items = []Item{{
	Id:           1290,
	ImageUrl:     "https://mediacloud.mobilelab.vn/2022-12-06/17_48_54-48757153-d1b1-4f0b-9b45-a9fb19fd9e23.jpg",
	Url:          "",
	Name:         " Trong suốt giai đoạn này, miền Nam Việt Nam đã phát hành nhiều đợt tiền giấy có chất lượng cao nhất thông qua Ngân hàng Quốc gia Việt Nam (NHQG), bao gồm các chủ đề về anh hùng dân tộc, tòa nhà NHQG và động vật. Bộ sưu tập tiền giấy đầy đủ trong giai đoạn này (cả miền Bắc và miền Nam Việt Nam) sẽ được trình bày trong các Phụ lục B và C. Phần dưới đây nêu lên những chủ đề chính trong mẫu thiết kế tiền giấy ở mỗi miền.   \n\nThere were many significant events that happened during this period. Ngo Dinh Diem and his brother Ngo Dinh Nu, who were backed by the Americans, were assassinated on 1 November 1963. A new administration headed by General Duong Van Minh continued to run South Vietnam with the backing of the American.  \nOn 2 September 1969, Ho Chi Minh passed away. However, the struggle for reunification continued. During this period, the National Liberation Front (NLF) began its series of influences to liberate the South from the Americans. Unification was finally achieved in 1975. \n Interestingly, during this period, South Vietnam through its National Bank of Vietnam (NBV) issued some of the most attractive series of notes, including the Heroes series, the National Bank Building series and the Animal series. The full range of currency notes issued (both South and North Vietnam) for this period are shown in the Annexes B and C. This section below highlights some of the main focus in the design of currency notes for each region. \n\nMặt trước :Quốc huy (bên phải) và chân dung Hồ Chí Minh bên trái\nObverse: National emblem (on the right) and portrait of Ho Chi Minh on the left\nMặt sau: Máy móc trang trại\nReverse: Machineries on farm \nPhát hành/ Issued in: 1969\n",
	Status:       8,
	CollectionId: 263,
	CreateTime:   1670323964,
	UpdateTime:   1670378689,
	ArrOwnerAddress: struct {
		Xce85A994E411901Fc9Bda228D07F7B66F8320Ecd int `json:"0xce85a994e411901fc9bda228d07f7b66f8320ecd"`
	}{100},
	ArrOwnerUid: struct {
		Field1 int `json:"221139"`
	}{100},
	OwnerAddress:   "0xce85a994e411901fc9bda228d07f7b66f8320ecd",
	OwnerUid:       221139,
	CreatorAddress: "0x3343d638fe2f2f8b5878577f3bc7b3970224ec9b",
	CreatorUid:     221168,
	ExternalLink:   "",
	EventIds:       nil,
	TotalView:      1,
	TotalLike:      0,
	Edition:        0,
	MetaData:       "{\"properties\":[{\"type\":\"Series Allocation\",\"name\":\"North Vietnam Emblem 1965 - 1975\"},{\"type\":\"Color\",\"name\":\"Green background with red images of Ho Chi Minh and Emblem\"}],\"sensitive\":false}",
	MediaType:      0,
	CategoryType:   5,
	ProductNo:      "",
	TotalEdition:   100,
	Uri:            "https://ipfs.trustkeys.network/ipfs/QmVcNo3KiKzQzj2hv98NTnYZHVT43HLsNfLAZQgpptgRgG/metadata.json",
	ExtendData: struct {
		Category   string `json:"category"`
		Collection string `json:"collection"`
		Create     string `json:"create"`
		Like       string `json:"like"`
		Owner      string `json:"owner"`
		TotalOffer string `json:"total_offer"`
	}{
		Category:   "{\"id\":5,\"name\":{\"en\":\"Collectibles\",\"ja\":\"Collectibles\",\"ko\":\"Collectibles\",\"th\":\"Collectibles\",\"vi\":\"Sưu tập\"},\"img\":\"https://mediacloud.mobilelab.vn/2022-10-06/15_37_43-ff506ae1-75b9-430a-8fd5-223e06d9e431.png\",\"background\":\"https://mediacloud.mobilelab.vn/2022-11-01/17_42_16-7806c98e-5945-4028-99b7-89d286574d31.jpg\",\"list_user\":{},\"id_main_category\":0,\"is_sub_category\":false,\"is_limit\":false,\"extend_data\":{}}",
		Collection: "{\"id\":263,\"name\":\"Vietnam Currency Notes: The Untold Stories\",\"image\":\"https://mediacloud.mobilelab.vn/2022-10-24/17_38_09-77745b4f-aada-4a49-9fc5-a4719b1bb89a.jpg\",\"avatar\":\"https://mediacloud.mobilelab.vn/2022-10-27/17_30_45-1d0502f5-3d4c-44b2-b5ad-54bf50b616c0.jpg\",\"description\":\"“Vietnam Currency Notes: The Untold Stories” is the FIRST BIGGEST \\u0026 TRUSTED NFT COLLECTION OF VIETNAM CURRENCY NOTES IN VIETNAM AND OVER THE WORLD, issued by Mr Thng Tien Tat - a famous Singaporean  collector.\",\"create_time\":1666604043,\"update_time\":1669881350,\"creator_id\":52245,\"creator_address\":\"0xAA3c0D54f8aa114E47E5698f2c1041D9250DBA79\",\"main_collection\":false,\"floor_price\":\"0\",\"volume_traded\":\"0\",\"extend_data\":{},\"contract_address\":\"\"}",
		Create:     "{\"uid\":221168,\"name\":\"Admin\",\"username\":\"Admin\",\"pubkey\":\"022be6845f1ff06b661192fa6de84612bf480cfa827fc5324aa67bbef091eb29eb\",\"bio\":\"Official Admin Account of TheOnly.Biz NFT Marketplace\",\"avatar\":\"https://photocloud.mobilelab.vn/2022-10-27/ded67944-2a4d-4c55-bfa9-e478ab5f176f.jpg\",\"email\":\"\",\"background\":\"https://mediacloud.mobilelab.vn/2022-11-01/17_52_41-32c3ad7a-a016-444e-8250-87a93392571a.jpg\",\"blockchain_address\":\"0x3343d638fe2f2f8b5878577f3bc7b3970224ec9b\",\"type_user\":0,\"can_mint\":false,\"total_following\":0,\"total_followers\":3,\"extend_data\":{},\"create_time\":1666836963,\"is_partner\":true}",
		Like:       "false",
		Owner:      "{\"uid\":221139,\"name\":\"Thng Tien Tat\",\"username\":\"tientat\",\"pubkey\":\"02c18becaa641b7ea88d410c84394c349047e2f1f0eb639dc96bdad001d552b093\",\"bio\":\"\\nFull name: Thng Tien Tat\\nDOB : 8th April 1968\\n- Tien Tat is currently working for a Private Equity firm in Singapore after moving back from Vietnam in 2016. Prior to working back to Singapore, he was the Country Manager for UOB in Vietnam. He has more than 20 years of experience in banking with the last 14 years developing banking business for UOB in the Mekong region, in particular Vietnam and Myanmar. Tien Tat used to be based in Ho Chi Minh City, Vietnam and was there since 2001.\\n- The former President of the Singapore Business Group in Vietnam, Tien Tat is active in the business community in Mekong Region. In 2010, he initiated an Interest Free Loan programme for more than 600 university students in Vietnam under the UOB Interest Free Loan Programme. In 2014, he also set up UOB University Scholarships Programme that benefitted 45 students in Myanmar.  \\n- Tien Tat is the author of book entitled “Vietnam Currency Notes : The Untold Stories” which\\nwas published in 2015.\\n- Tien Tat graduated from Nanyang Technological University, Singapore with a Degree in Banking and Insurance. He speaks English and Mandarin.\",\"avatar\":\"https://mediacloud.mobilelab.vn/2022-10-26/11_06_03-4c0148c0-d964-4170-8241-129ee56174db.png\",\"email\":\"\",\"background\":\"https://mediacloud.mobilelab.vn/2022-10-25/12_05_51-6dd5fe31-2ff6-407f-b675-33c3fdad2111.jpg\",\"blockchain_address\":\"0xce85a994e411901fc9bda228d07f7b66f8320ecd\",\"type_user\":0,\"can_mint\":false,\"total_following\":0,\"total_followers\":0,\"extend_data\":{},\"create_time\":1666597593,\"is_partner\":true}",
		TotalOffer: "0",
	},
	IsSensitive:     false,
	Hidden:          false,
	Price:           "000000000000000000000000000000000000000000000000000000000200000000000000000000",
	TokenId:         "53",
	ContractAddress: "0x90912bfaC6D123e299fE7F075635432787532788",
	Chain:           "bsc",
	CurrencyAddress: struct {
		Name    string `json:"name"`
		Symbol  string `json:"symbol"`
		Address string `json:"address"`
		Chain   string `json:"chain"`
		Decimal int    `json:"decimal"`
		Img     string `json:"img"`
		Status  bool   `json:"status"`
	}{Name: "BUSD Token",
		Symbol:  "BUSD",
		Address: "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
		Chain:   "bsc",
		Decimal: 18,
		Img:     "https://mediacloud.mobilelab.vn/2022-07-28/11_52_49-b9cfc474-a74a-4dde-bbca-6f7cf8046a17.png",
		Status:  false},
	TypeNft:           1,
	TotalClaim:        0,
	TotalLimit:        0,
	StartAt:           0,
	EndAt:             0,
	UnlockableContent: "{}",
	IsLootBox:         false,
},
}
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
	r.DELETE("/item/", deleteItems)
	r.PUT("/item/", editItems)
	r.POST("/user", addUser)
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
	pos, _ := c.GetQuery("pos")
	cou, _ := c.GetQuery("count")
	position, _ := strconv.Atoi(pos)
	count, _ := strconv.Atoi(cou)
	val, _ := client.ZRange(ctx, "item", int64(position), int64(position+count)).Result()
	c.IndentedJSON(http.StatusOK, val)
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
// @Param id query int true "Id cần xóa  "
// @Produce json
// @Success 200 {object} Item
// @Router /item/ [delete]
func deleteItems(c *gin.Context) {
	id, _ := c.GetQuery("id")
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
		c.IndentedJSON(http.StatusNoContent, items)
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
// @Param ID query int true "ID cần sửa"
// @Param ItemEdit body Item true "Thông tin cần sửa"
// @Produce json
// @Success 200 {object} Item
// @Router /item/ [put]
func editItems(c *gin.Context) {
	id, _ := c.GetQuery("ID")
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
// @Tags Item
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
