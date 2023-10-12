package swa

import "time"

type CurrencyList struct {
	Success bool   `json:"success"`
	Base    string `json:"base"`
	Date    string `json:"date"`
	Rates   struct {
		AED float64 `json:"AED"`
		AFN float64 `json:"AFN"`
		ALL float64 `json:"ALL"`
		AMD float64 `json:"AMD"`
		ANG float64 `json:"ANG"`
		AOA float64 `json:"AOA"`
		ARS float64 `json:"ARS"`
		AUD float64 `json:"AUD"`
		AWG float64 `json:"AWG"`
		AZN float64 `json:"AZN"`
		BAM float64 `json:"BAM"`
		BBD float64 `json:"BBD"`
		BDT float64 `json:"BDT"`
		BGN float64 `json:"BGN"`
		BHD float64 `json:"BHD"`
		BIF float64 `json:"BIF"`
		BMD int     `json:"BMD"`
		BND float64 `json:"BND"`
		BOB float64 `json:"BOB"`
		BRL float64 `json:"BRL"`
		BSD float64 `json:"BSD"`
		BTC float64 `json:"BTC"`
		BTN float64 `json:"BTN"`
		BWP float64 `json:"BWP"`
		BYN float64 `json:"BYN"`
		BYR float64 `json:"BYR"`
		BZD float64 `json:"BZD"`
		CAD float64 `json:"CAD"`
		CDF float64 `json:"CDF"`
		CHF float64 `json:"CHF"`
		CLF float64 `json:"CLF"`
		CLP float64 `json:"CLP"`
		CNY float64 `json:"CNY"`
		COP float64 `json:"COP"`
		CRC float64 `json:"CRC"`
		CUC int     `json:"CUC"`
		CUP float64 `json:"CUP"`
		CVE float64 `json:"CVE"`
		CZK float64 `json:"CZK"`
		DJF float64 `json:"DJF"`
		DKK float64 `json:"DKK"`
		DOP float64 `json:"DOP"`
		DZD float64 `json:"DZD"`
		EGP float64 `json:"EGP"`
		ERN float64 `json:"ERN"`
		ETB float64 `json:"ETB"`
		EUR float64 `json:"EUR"`
		FJD float64 `json:"FJD"`
		FKP float64 `json:"FKP"`
		GBP float64 `json:"GBP"`
		GEL float64 `json:"GEL"`
		GGP float64 `json:"GGP"`
		GHS float64 `json:"GHS"`
		GIP float64 `json:"GIP"`
		GMD float64 `json:"GMD"`
		GNF float64 `json:"GNF"`
		GTQ float64 `json:"GTQ"`
		GYD float64 `json:"GYD"`
		HKD float64 `json:"HKD"`
		HNL float64 `json:"HNL"`
		HRK float64 `json:"HRK"`
		HTG float64 `json:"HTG"`
		HUF float64 `json:"HUF"`
		IDR float64 `json:"IDR"`
		ILS float64 `json:"ILS"`
		IMP float64 `json:"IMP"`
		INR float64 `json:"INR"`
		IQD float64 `json:"IQD"`
		IRR float64 `json:"IRR"`
		ISK float64 `json:"ISK"`
		JEP float64 `json:"JEP"`
		JMD float64 `json:"JMD"`
		JOD float64 `json:"JOD"`
		JPY float64 `json:"JPY"`
		KES float64 `json:"KES"`
		KGS float64 `json:"KGS"`
		KHR float64 `json:"KHR"`
		KMF float64 `json:"KMF"`
		KPW float64 `json:"KPW"`
		KRW float64 `json:"KRW"`
		KWD float64 `json:"KWD"`
		KYD float64 `json:"KYD"`
		KZT float64 `json:"KZT"`
		LAK float64 `json:"LAK"`
		LBP float64 `json:"LBP"`
		LKR float64 `json:"LKR"`
		LRD float64 `json:"LRD"`
		LSL float64 `json:"LSL"`
		LTL float64 `json:"LTL"`
		LVL float64 `json:"LVL"`
		LYD float64 `json:"LYD"`
		MAD float64 `json:"MAD"`
		MDL float64 `json:"MDL"`
		MGA float64 `json:"MGA"`
		MKD float64 `json:"MKD"`
		MMK float64 `json:"MMK"`
		MNT float64 `json:"MNT"`
		MOP float64 `json:"MOP"`
		MRO float64 `json:"MRO"`
		MUR float64 `json:"MUR"`
		MVR float64 `json:"MVR"`
		MWK float64 `json:"MWK"`
		MXN float64 `json:"MXN"`
		MYR float64 `json:"MYR"`
		MZN float64 `json:"MZN"`
		NAD float64 `json:"NAD"`
		NGN float64 `json:"NGN"`
		NIO float64 `json:"NIO"`
		NOK float64 `json:"NOK"`
		NPR float64 `json:"NPR"`
		NZD float64 `json:"NZD"`
		OMR float64 `json:"OMR"`
		PAB float64 `json:"PAB"`
		PEN float64 `json:"PEN"`
		PGK float64 `json:"PGK"`
		PHP float64 `json:"PHP"`
		PKR float64 `json:"PKR"`
		PLN float64 `json:"PLN"`
		PYG float64 `json:"PYG"`
		QAR float64 `json:"QAR"`
		RON float64 `json:"RON"`
		RSD float64 `json:"RSD"`
		RUB float64 `json:"RUB"`
		RWF float64 `json:"RWF"`
		SAR float64 `json:"SAR"`
		SBD float64 `json:"SBD"`
		SCR float64 `json:"SCR"`
		SDG float64 `json:"SDG"`
		SEK float64 `json:"SEK"`
		SGD float64 `json:"SGD"`
		SHP float64 `json:"SHP"`
		SLE float64 `json:"SLE"`
		SLL float64 `json:"SLL"`
		SOS float64 `json:"SOS"`
		SSP float64 `json:"SSP"`
		SRD float64 `json:"SRD"`
		STD float64 `json:"STD"`
		SYP float64 `json:"SYP"`
		SZL float64 `json:"SZL"`
		THB float64 `json:"THB"`
		TJS float64 `json:"TJS"`
		TMT float64 `json:"TMT"`
		TND float64 `json:"TND"`
		TOP float64 `json:"TOP"`
		TRY float64 `json:"TRY"`
		TTD float64 `json:"TTD"`
		TWD float64 `json:"TWD"`
		TZS float64 `json:"TZS"`
		UAH float64 `json:"UAH"`
		UGX float64 `json:"UGX"`
		USD int     `json:"USD"`
		UYU float64 `json:"UYU"`
		UZS float64 `json:"UZS"`
		VEF float64 `json:"VEF"`
		VES float64 `json:"VES"`
		VND float64 `json:"VND"`
		VUV float64 `json:"VUV"`
		WST float64 `json:"WST"`
		XAF float64 `json:"XAF"`
		XAG float64 `json:"XAG"`
		XAU float64 `json:"XAU"`
		XCD float64 `json:"XCD"`
		XDR float64 `json:"XDR"`
		XOF float64 `json:"XOF"`
		XPF float64 `json:"XPF"`
		YER float64 `json:"YER"`
		ZAR float64 `json:"ZAR"`
		ZMK float64 `json:"ZMK"`
		ZMW float64 `json:"ZMW"`
		ZWL float64 `json:"ZWL"`
	} `json:"rates"`
}

type ExchangeRate struct {
	BaseRate   int     `json:"baseRate"`
	ChangeRate float64 `json:"changeRate"`
	Change     string  `json:"change"`
	Base       string  `json:"base"`
	Symbol     string  `json:"symbol"`
}

type RandomProfile struct {
	Status string `json:"status"`
	Data   []struct {
		Steamid                  string      `json:"steamid"`
		Communityvisibilitystate *int        `json:"communityvisibilitystate"`
		Profilestate             *int        `json:"profilestate"`
		Personaname              string      `json:"personaname"`
		Commentpermission        *int        `json:"commentpermission"`
		Profileurl               string      `json:"profileurl"`
		Avatar                   string      `json:"avatar"`
		Avatarmedium             string      `json:"avatarmedium"`
		Avatarfull               string      `json:"avatarfull"`
		Avatarhash               string      `json:"avatarhash"`
		Lastlogoff               interface{} `json:"lastlogoff"`
		Personastate             interface{} `json:"personastate"`
		Realname                 *string     `json:"realname"`
		Primaryclanid            *string     `json:"primaryclanid"`
		Timecreated              *int        `json:"timecreated"`
		Locstatecode             interface{} `json:"locstatecode"`
		Loccityid                interface{} `json:"loccityid"`
		Loccountrycode           interface{} `json:"loccountrycode"`
		Accountname              string      `json:"accountname"`
		Inventoryworth           struct {
			Worth     float64   `json:"worth"`
			Size      int       `json:"size"`
			CreatedAt time.Time `json:"createdAt"`
		} `json:"inventoryworth,omitempty"`
	} `json:"data"`
}

type HighValueProfile struct {
	Status string `json:"status"`
	Data   []struct {
		Accountname interface{} `json:"accountname"`
		Username    string      `json:"username"`
		Steamid     string      `json:"steamid"`
		Worth       string      `json:"worth"`
		Size        int         `json:"size"`
		Created     time.Time   `json:"created"`
		Avatarhash  string      `json:"avatarhash"`
		Country     string      `json:"country"`
	} `json:"data"`
}

// will be []
type GameItem struct {
	Id                        string      `json:"id"`
	Markethashname            string      `json:"markethashname"`
	Marketname                string      `json:"marketname"`
	Slug                      string      `json:"slug"`
	Color                     string      `json:"color"`
	Bordercolor               string      `json:"bordercolor"`
	Pricelatest               float64     `json:"pricelatest"`
	Pricemedian               float64     `json:"pricemedian"`
	Priceavg                  float64     `json:"priceavg"`
	Pricereal                 float64     `json:"pricereal"`
	Pricemin                  float64     `json:"pricemin"`
	Pricemax                  float64     `json:"pricemax"`
	Pricesafe                 float64     `json:"pricesafe"`
	Pricesafe24H              float64     `json:"pricesafe24h"`
	Pricesafe7D               float64     `json:"pricesafe7d"`
	Pricesafe30D              float64     `json:"pricesafe30d"`
	Pricesafe90D              float64     `json:"pricesafe90d"`
	Winlosspercentage         string      `json:"winlosspercentage"`
	Unstable                  int         `json:"unstable"`
	Unstablereason            interface{} `json:"unstablereason"`
	Dailysoldvolume           int         `json:"dailysoldvolume"`
	Sold24H                   int         `json:"sold24h"`
	Sold7D                    int         `json:"sold7d"`
	Sold30D                   int         `json:"sold30d"`
	Sold90D                   int         `json:"sold90d"`
	Isstattrack               int         `json:"isstattrack"`
	Iscase                    int         `json:"iscase"`
	Iskey                     int         `json:"iskey"`
	Isgraffiti                int         `json:"isgraffiti"`
	Issticker                 int         `json:"issticker"`
	Wear                      string      `json:"wear"`
	Itemgroup                 string      `json:"itemgroup"`
	Type                      string      `json:"type"`
	Skin                      string      `json:"skin"`
	Isstar                    int         `json:"isstar"`
	Rarity                    interface{} `json:"rarity"`
	Quality                   string      `json:"quality"`
	Markettradablerestriction interface{} `json:"markettradablerestriction"`
	Priceupdatedat            string      `json:"priceupdatedat"`
	Updatedat                 string      `json:"updatedat"`
	Itemimage                 string      `json:"itemimage"`
}

// will be []
type SearchItem struct {
	Id              string        `json:"id"`
	Count           int           `json:"count"`
	Assetid         interface{}   `json:"assetid"`
	Classid         *string       `json:"classid"`
	Instanceid      *string       `json:"instanceid"`
	Markethashname  string        `json:"markethashname"`
	Marketname      string        `json:"marketname"`
	Hashid          string        `json:"hashid"`
	Nameid          int           `json:"nameid"`
	Color           *string       `json:"color"`
	Bordercolor     string        `json:"bordercolor"`
	Type            string        `json:"type"`
	Rarity          *string       `json:"rarity"`
	Quality         *string       `json:"quality"`
	Collectionname  interface{}   `json:"collectionname"`
	Marketable      interface{}   `json:"marketable"`
	Tradable        interface{}   `json:"tradable"`
	Pricelatest     float64       `json:"pricelatest"`
	Pricemedian     float64       `json:"pricemedian"`
	Pricesafe       float64       `json:"pricesafe"`
	Priceavg        float64       `json:"priceavg"`
	Pricemin        float64       `json:"pricemin"`
	Pricemax        float64       `json:"pricemax"`
	Pricesafe24H    float64       `json:"pricesafe24h"`
	Pricesafe7D     float64       `json:"pricesafe7d"`
	Pricesafe30D    float64       `json:"pricesafe30d"`
	Pricesafe90D    float64       `json:"pricesafe90d"`
	Pricereal       interface{}   `json:"pricereal"`
	Pricemix        float64       `json:"pricemix"`
	Sold24H         int           `json:"sold24h"`
	Sold7D          int           `json:"sold7d"`
	Sold30D         int           `json:"sold30d"`
	Sold90D         int           `json:"sold90d"`
	Slug            string        `json:"slug"`
	Image           string        `json:"image"`
	ItemImages      []string      `json:"itemImages"`
	Wear            string        `json:"wear"`
	Insertid        int           `json:"insertid"`
	Tag             []interface{} `json:"tag"`
	Createdat       time.Time     `json:"createdat"`
	Unstable        bool          `json:"unstable"`
	Unstablereason  *string       `json:"unstablereason"`
	Stattrack       bool          `json:"stattrack"`
	Dailysoldvolume *int          `json:"dailysoldvolume"`
	Firstseentime   string        `json:"firstseentime"`
	Iscase          bool          `json:"iscase"`
	Iskey           bool          `json:"iskey"`
	Isgraffiti      bool          `json:"isgraffiti"`
	Issticker       bool          `json:"issticker"`
	Itemgroup       string        `json:"itemgroup"`
	Isstar          bool          `json:"isstar"`
	Actions         []struct {
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"actions"`
	Marketactions []struct {
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"marketactions"`
	Descriptions []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
		Color string `json:"color,omitempty"`
	} `json:"descriptions"`
	Markettradablerestriction *int `json:"markettradablerestriction"`
	Tags                      []struct {
		Category              string `json:"category"`
		InternalName          string `json:"internal_name"`
		LocalizedCategoryName string `json:"localized_category_name"`
		LocalizedTagName      string `json:"localized_tag_name"`
		Color                 string `json:"color,omitempty"`
	} `json:"tags"`
	Inspectlink       *string       `json:"inspectlink"`
	Inspectlinkparsed interface{}   `json:"inspectlinkparsed"`
	Prices            []interface{} `json:"prices"`
	Winloss           interface{}   `json:"winloss"`
}

// item details
type ItemDetail struct {
	Id              string        `json:"id"`
	Assetid         interface{}   `json:"assetid"`
	Classid         interface{}   `json:"classid"`
	Instanceid      interface{}   `json:"instanceid"`
	Markethashname  string        `json:"markethashname"`
	Marketname      string        `json:"marketname"`
	Hashid          string        `json:"hashid"`
	Nameid          int           `json:"nameid"`
	Color           string        `json:"color"`
	Bordercolor     string        `json:"bordercolor"`
	Type            string        `json:"type"`
	Rarity          string        `json:"rarity"`
	Quality         string        `json:"quality"`
	Marketable      interface{}   `json:"marketable"`
	Tradable        interface{}   `json:"tradable"`
	Pricelatest     float64       `json:"pricelatest"`
	Pricemedian     float64       `json:"pricemedian"`
	Pricesafe       float64       `json:"pricesafe"`
	Priceavg        float64       `json:"priceavg"`
	Pricemin        float64       `json:"pricemin"`
	Pricemax        float64       `json:"pricemax"`
	Pricesafe24H    float64       `json:"pricesafe24h"`
	Pricesafe7D     float64       `json:"pricesafe7d"`
	Pricesafe30D    float64       `json:"pricesafe30d"`
	Pricesafe90D    float64       `json:"pricesafe90d"`
	Pricereal       int           `json:"pricereal"`
	Pricemix        float64       `json:"pricemix"`
	Sold24H         int           `json:"sold24h"`
	Sold7D          int           `json:"sold7d"`
	Sold30D         int           `json:"sold30d"`
	Sold90D         int           `json:"sold90d"`
	Slug            string        `json:"slug"`
	Image           string        `json:"image"`
	ItemImages      []string      `json:"itemImages"`
	Wear            string        `json:"wear"`
	Insertid        int           `json:"insertid"`
	Tag             []interface{} `json:"tag"`
	Createdat       time.Time     `json:"createdat"`
	Unstable        bool          `json:"unstable"`
	Unstablereason  string        `json:"unstablereason"`
	Stattrack       bool          `json:"stattrack"`
	Dailysoldvolume int           `json:"dailysoldvolume"`
	Firstseentime   string        `json:"firstseentime"`
	Iscase          bool          `json:"iscase"`
	Iskey           bool          `json:"iskey"`
	Isgraffiti      bool          `json:"isgraffiti"`
	Issticker       bool          `json:"issticker"`
	Itemgroup       string        `json:"itemgroup"`
	Isstar          bool          `json:"isstar"`
	Actions         interface{}   `json:"actions"`
	Marketactions   interface{}   `json:"marketactions"`
	Descriptions    []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
		Color string `json:"color,omitempty"`
	} `json:"descriptions"`
	Markettradablerestriction interface{} `json:"markettradablerestriction"`
	Tags                      []struct {
		Category              string `json:"category"`
		InternalName          string `json:"internal_name"`
		LocalizedCategoryName string `json:"localized_category_name"`
		LocalizedTagName      string `json:"localized_tag_name"`
		Color                 string `json:"color,omitempty"`
	} `json:"tags"`
	Inspectlink       interface{} `json:"inspectlink"`
	Inspectlinkparsed interface{} `json:"inspectlinkparsed"`
	Prices            []struct {
		Source    string        `json:"source"`
		Name      string        `json:"name"`
		Type      string        `json:"type"`
		Tags      []interface{} `json:"tags"`
		Createdat string        `json:"createdat"`
		Price     float64       `json:"price"`
		Logo      string        `json:"logo"`
		Winloss   float64       `json:"winloss"`
		Link      string        `json:"link"`
	} `json:"prices"`
	Winloss float64 `json:"winloss"`
}

// item history will be []array of this struct
type ItemHistory struct {
	Id        string  `json:"id"`
	Price     float64 `json:"price"`
	Type      string  `json:"type"`
	Source    string  `json:"source"`
	Createdat string  `json:"createdat"`
	Status    bool    `json:"status"`
	Realprice bool    `json:"realprice"`
}

// steam inventory will be []array of this struct
type SteamInventory struct {
	Id              string        `json:"id"`
	Count           int           `json:"count"`
	Assetid         string        `json:"assetid"`
	Classid         string        `json:"classid"`
	Instanceid      string        `json:"instanceid"`
	Markethashname  string        `json:"markethashname"`
	Marketname      string        `json:"marketname"`
	Hashid          string        `json:"hashid"`
	Nameid          int           `json:"nameid"`
	Color           string        `json:"color"`
	Bordercolor     string        `json:"bordercolor"`
	Type            string        `json:"type"`
	Rarity          string        `json:"rarity"`
	Quality         string        `json:"quality"`
	Marketable      bool          `json:"marketable"`
	Tradable        bool          `json:"tradable"`
	Pricelatest     float64       `json:"pricelatest"`
	Pricemedian     float64       `json:"pricemedian"`
	Pricesafe       float64       `json:"pricesafe"`
	Priceavg        float64       `json:"priceavg"`
	Pricemin        float64       `json:"pricemin"`
	Pricemax        float64       `json:"pricemax"`
	Pricesafe24H    float64       `json:"pricesafe24h"`
	Pricesafe7D     float64       `json:"pricesafe7d"`
	Pricesafe30D    float64       `json:"pricesafe30d"`
	Pricesafe90D    float64       `json:"pricesafe90d"`
	Pricereal       float64       `json:"pricereal"`
	Pricemix        float64       `json:"pricemix"`
	Sold24H         int           `json:"sold24h"`
	Sold7D          int           `json:"sold7d"`
	Sold30D         int           `json:"sold30d"`
	Sold90D         int           `json:"sold90d"`
	Slug            string        `json:"slug"`
	Image           string        `json:"image"`
	ItemImages      []string      `json:"itemImages"`
	Wear            *string       `json:"wear"`
	Insertid        int           `json:"insertid"`
	Tag             []interface{} `json:"tag"`
	Createdat       time.Time     `json:"createdat"`
	Unstable        bool          `json:"unstable"`
	Unstablereason  *string       `json:"unstablereason"`
	Stattrack       bool          `json:"stattrack"`
	Dailysoldvolume int           `json:"dailysoldvolume"`
	Firstseentime   string        `json:"firstseentime"`
	Iscase          bool          `json:"iscase"`
	Iskey           bool          `json:"iskey"`
	Isgraffiti      bool          `json:"isgraffiti"`
	Issticker       bool          `json:"issticker"`
	Itemgroup       string        `json:"itemgroup"`
	Isstar          bool          `json:"isstar"`
	Actions         []struct {
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"actions"`
	Marketactions []struct {
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"marketactions"`
	Descriptions []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
		Color string `json:"color,omitempty"`
	} `json:"descriptions"`
	Markettradablerestriction int `json:"markettradablerestriction"`
	Tags                      []struct {
		Category              string `json:"category"`
		InternalName          string `json:"internal_name"`
		LocalizedCategoryName string `json:"localized_category_name"`
		LocalizedTagName      string `json:"localized_tag_name"`
		Color                 string `json:"color,omitempty"`
	} `json:"tags"`
	Inspectlink       *string `json:"inspectlink"`
	Inspectlinkparsed *struct {
		A string `json:"a"`
		S string `json:"s"`
		D string `json:"d"`
	} `json:"inspectlinkparsed"`
	Prices  []interface{} `json:"prices"`
	Winloss float64       `json:"winloss"`
}

// steam ids
type SteamIds struct {
	Steamids struct {
		Steamid2  string `json:"steamid2"`
		Steamid3  string `json:"steamid3"`
		Steamid64 string `json:"steamid64"`
	} `json:"steamids"`
}

// steam user profile
type SteamProfile struct {
	Steamid                  string `json:"steamid"`
	Communityvisibilitystate int    `json:"communityvisibilitystate"`
	Profilestate             int    `json:"profilestate"`
	Personaname              string `json:"personaname"`
	Profileurl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	Avatarmedium             string `json:"avatarmedium"`
	Avatarfull               string `json:"avatarfull"`
	Avatarhash               string `json:"avatarhash"`
	Personastate             int    `json:"personastate"`
	Realname                 string `json:"realname"`
	Primaryclanid            string `json:"primaryclanid"`
	Timecreated              int    `json:"timecreated"`
	Personastateflags        int    `json:"personastateflags"`
	Steamids                 struct {
		Steamid2  string `json:"steamid2"`
		Steamid3  string `json:"steamid3"`
		Steamid64 string `json:"steamid64"`
	} `json:"steamids"`
	Accountname string `json:"accountname"`
}

// steam friend list will be array of this struct
type SteamFriends struct {
	Steamid                  string      `json:"steamid"`
	Communityvisibilitystate int         `json:"communityvisibilitystate"`
	Profilestate             *int        `json:"profilestate"`
	Personaname              string      `json:"personaname"`
	Commentpermission        *int        `json:"commentpermission"`
	Profileurl               string      `json:"profileurl"`
	Avatar                   string      `json:"avatar"`
	Avatarmedium             string      `json:"avatarmedium"`
	Avatarfull               string      `json:"avatarfull"`
	Avatarhash               string      `json:"avatarhash"`
	Lastlogoff               interface{} `json:"lastlogoff"`
	Personastate             int         `json:"personastate"`
	Realname                 *string     `json:"realname"`
	Primaryclanid            *string     `json:"primaryclanid"`
	Timecreated              *int        `json:"timecreated"`
	Locstatecode             *string     `json:"locstatecode"`
	Loccityid                *int        `json:"loccityid"`
	Loccountrycode           *string     `json:"loccountrycode"`
	Accountname              string      `json:"accountname"`
	Friendssince             int         `json:"friendssince"`
	Friendssinceat           time.Time   `json:"friendssinceat"`
	Lastonlinebeforeseconds  interface{} `json:"lastonlinebeforeseconds"`
	Lastonlinebeforeat       interface{} `json:"lastonlinebeforeat"`
}
