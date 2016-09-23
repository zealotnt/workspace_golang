package main

import (
	"github.com/fatih/color"
	"github.com/fatih/structs"

	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Amount struct {
	Subtotal uint `json:"subtotal"`
	Shipping uint `json:"shipping"`
	Total    uint `json:"total"`
}

type OrderInfo struct {
	OrderCode       string `json:"id" sql:"order_code"`
	CustomerName    string `json:"name" sql:"customer_name"`
	CustomerAddress string `json:"address" sql:"customer_address"`
	CustomerPhone   string `json:"phone" sql:"customer_phone"`
	CustomerEmail   string `json:"email" sql:"customer_email"`
	CustomerNote    string `json:"note" sql:"customer_note"`
}

type OrderItem struct {
	Id                       uint    `json:"id" sql:"id"`
	OrderId                  uint    `json:"-" sql:"REFERENCES Orders(id)"`
	ProductId                uint    `json:"product_id" sql:"REFERENCES products(id)"`
	ProductName              string  `json:"name" sql:"-"`
	ProductImageThumbnailUrl *string `json:"image_thumbnail_url" sql:"-"`
	ProductPrice             int     `json:"price" sql:"-"`
	Quantity                 int     `json:"quantity"`
}

type Order struct {
	Id          uint   `json:"-"`
	AccessToken string `json:"access_token,omitempty"`
	Status      string `json:"-"`

	Items []OrderItem `json:"items" sql:"order_items"`

	OrderInfo `json:"customer"`

	Amounts Amount `json:"amounts" sql:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    int     `json:"price"`
	Provider string  `json:"provider"`
	Rating   float32 `json:"rating"`
	Status   string  `json:"status"`
	Detail   string  `json:"detail"`

	ImageData          []byte    `json:"-" sql:"-"`
	ImageThumbnailData []byte    `json:"-" sql:"-"`
	ImageDetailData    []byte    `json:"-" sql:"-"`
	Image              string    `json:"-"`
	NewImage           string    `json:"-" sql:"-"`
	ImageUpdatedAt     time.Time `json:"-"`
	ImageUrl           *string   `json:"image_url" sql:"-"`
	ImageThumbnailUrl  *string   `json:"image_thumbnail_url" sql:"-"`
	ImageDetailUrl     *string   `json:"image_detail_url" sql:"-"`

	AvatarData          []byte    `json:"-" sql:"-"`
	AvatarThumbnailData []byte    `json:"-" sql:"-"`
	AvatarDetailData    []byte    `json:"-" sql:"-"`
	Avatar              string    `json:"-"`
	NewAvatar           string    `json:"-" sql:"-"`
	AvatarUpdatedAt     time.Time `json:"-"`
	AvatarUrl           *string   `json:"avatar_url" sql:"-"`
	AvatarThumbnailUrl  *string   `json:"avatar_thumbnail_url" sql:"-"`
	AvatarDetailUrl     *string   `json:"avatar_detail_url" sql:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vertex struct {
	X int
	Y int
	Z int
}

var PR = fmt.Println
var PR_ERR = color.New(color.FgRed).PrintfFunc()
var PR_NOTI = color.New(color.FgYellow).PrintfFunc()

func simple_test() {
	v := Vertex{1, 2, 3}

	if val, err := GetField(&v, "D"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}

func my_order_test(field string) {
	PR_NOTI("\r\n\tTest with Order struct, with field = **" + field + "**\r\n")

	order := Order{
		Id:          1,
		AccessToken: "this_is_a_access_token",
		Status:      "cart",

		OrderInfo: OrderInfo{
			OrderCode:       "this_is_order_code",
			CustomerName:    "T V A",
			CustomerAddress: "HCMUT",
			CustomerPhone:   "123456",
			CustomerEmail:   "tva@abc.com",
			CustomerNote:    "Some note",
		},

		Amounts: Amount{
			Subtotal: 100,
			Shipping: 10,
			Total:    110,
		},

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	ret, err := FieldSelection(order, field)
	if err != nil {
		PR("error: " + err.Error())
		return
	}

	json_out, _ := json.MarshalIndent(ret, "", "    ")
	fmt.Printf("Json output:\r\n%s\r\n", json_out)
}

func my_product_test() {
	PR_NOTI("\r\n\tTest with Product struct:\r\n")

	p := Product{
		Id:       1,
		Name:     "XBox",
		Price:    12000,
		Provider: "Sony",
		Rating:   4.5,
	}

	ret, err := FieldSelection(p, "Id,name,price,provider,rating,status,Detail")
	if err != nil {
		PR("error: " + err.Error())
		return
	}

	json_out, _ := json.Marshal(ret)
	fmt.Printf("Json output:\r\n%s\r\n", json_out)
}

func main() {
	// simple_test()
	// my_product_test()
	// my_order_test("Id,access_token")
	// my_order_test("access_token,amounts,created_at,updated_at,OrderInfo")
	my_order_test("")
}

func ValidateStringField(field string) ([]string, error) {
	// Remove space if any
	field = strings.Replace(field, " ", "", -1)

	// Split the field by comma
	fields := strings.Split(field, ",")

	return fields, nil
}

func FieldSelection(v interface{}, field string) (map[string]interface{}, error) {
	var fields []string
	var err error
	map_out := make(map[string]interface{})

	// If fields is empty, just return the whole struct
	if field == "" {
		s := structs.New(v)
		s.TagName = "json"
		return s.Map(), nil
	}

	// Check the input field, is it in the right format
	if fields, err = ValidateStringField(field); err != nil {
		return nil, err
	}

	// Loop through the field
	for _, field := range fields {
		field_name, _ := GetFieldNameFromJson(v, field)

		// Check if the field in the struct
		field_value, err := GetField(v, field_name)
		if err != nil {
			return nil, err
		}

		// if json tag specify the field's name, use it
		field_name, err = GetFieldJsonName(v, field_name, field_value)
		if err != nil {
			return nil, err
		}

		// add it to the map output
		map_out[field_name] = field_value
	}

	return map_out, nil
}

func GetFieldJsonName(v interface{}, field_name string, field_value interface{}) (string, error) {
	s := structs.New(v)
	f := s.Field(field_name)

	// Get the value of field's json tag value
	json_tag := f.Tag("json")

	// Split the tag value by comma
	json_fields := strings.Split(json_tag, ",")

	// If there is no value in json tag -> len return is 0, use FieldName instead
	if len(json_fields) == 0 {
		return field_name, nil
	}

	// If user wants to select unexported json field, returns error
	if json_fields[0] == "-" {
		return "", errors.New(field_name + " is not exported to struct " + s.Name() + "'s json output")
	}

	// If the tags is ",omitempty"
	// -> the slice return is : len = 2, elem_1 = "", elem_2 = "omitempty"
	// If the tag has value ex, "access_token,omitempty"
	// -> elem_1 = "access_token"
	if json_fields[0] != "" {
		return json_fields[0], nil
	}

	// The json tag hasn't specified the field's return name, use the field name
	return field_name, nil
}

func GetFieldNameFromJson(v interface{}, json_name string) (string, error) {
	s := structs.New(v)
	fs := s.Fields()
	for _, f := range fs {
		json_tag := f.Tag("json")
		json_fields := strings.Split(json_tag, ",")
		if len(json_fields) == 0 {
			continue
		}
		if json_fields[0] == json_name {
			return f.Name(), nil
		}
	}
	return json_name, nil
}

func GetField(v interface{}, field string) (interface{}, error) {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	if !f.IsValid() {
		panic(field + " is not part of " + r.Type().String())
		return nil, errors.New(field + " is not part of " + r.Type().String())
	}
	return f.Interface(), nil
}
