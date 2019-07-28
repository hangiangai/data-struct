package main

import (
	"7-26-restful/model"
	"7-26-restful/restful"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
)

const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "hangiangai"
)

func main() {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
	}

	app := iris.New()

	r := restful.New()

	r.Init(db, app)

	r.Register(model.Student{}, "student")
	r.Register(model.Teacher{}, "teacher")

	app.Run(iris.Addr(":8000"))

	// t := reflect.TypeOf(3)
	// fmt.Println(t)
	// fmt.Println(t.String())

	// v := reflect.ValueOf(3)
	// fmt.Println(v)
	// fmt.Printf("%v\n", v)
	// fmt.Println(v.String())

	// t1 := v.Type()
	// fmt.Println(t1.String())

	// v := reflect.ValueOf(3)

	// x := v.Interface()

	// i := x.(int)

	// fmt.Println(i)

	// strangelove := Movie{
	// 	Title:    "Dr. Strangelove",
	// 	Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
	// 	Year:     1964,
	// 	Color:    false,
	// 	Actor: map[string]string{
	// 		"Dr. Strangelove":            "Peter Sellers",
	// 		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
	// 		"Pres. Merkin Muffley":       "Peter Sellers",
	// 		"Gen. Buck Turgidson":        "George C. Scott",
	// 		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
	// 		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	// 	},

	// 	Oscars: []string{
	// 		"Best Actor (Nomin.)",
	// 		"Best Adapted Screenplay (Nomin.)",
	// 		"Best Director (Nomin.)",
	// 		"Best Picture (Nomin.)",
	// 	},
	// }

	// Display("strangelove", strangelove)

}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	//v.Kind()返回Type
	switch v.Kind() {
	case reflect.Invalid:
		return "Invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Func, reflect.Map, reflect.Slice, reflect.Ptr, reflect.Chan:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + "value"
	}
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

//
func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s.[%d]", path, i), v.Index(i))
		}
	case reflect.Struct: //结构体
		for i := 0; i < v.NumField(); i++ {
			filepath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(filepath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("*%s", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		// basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}

}
