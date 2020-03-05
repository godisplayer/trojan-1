package util

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	// RED 红色
	RED = "\033[31m"
	// GREEN 绿色
	GREEN = "\033[32m"
	// YELLOW 黄色
	YELLOW = "\033[33m"
	// BLUE 蓝色
	BLUE = "\033[34m"
	// FUCHSIA 紫红色
	FUCHSIA = "\033[35m"
	// CYAN 青色
	CYAN = "\033[36m"
	// WHITE 白色
	WHITE = "\033[37m"
	// RESET 重置颜色
	RESET = "\033[0m"
)

// IsNumeric is_numeric()
// Numeric strings consist of optional sign, any number of digits, optional decimal part and optional exponential part.
// Thus +0123.45e6 is a valid numeric value.
func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		// Trim any whitespace
		str = strings.TrimSpace(str)
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// hex
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9, Point, Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}

	return false
}

// RandString 随机字符串
func RandString(length int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getChar(str string) string {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()
	fmt.Print(str)
	char, _, _ := keyboard.GetKey()
	fmt.Printf("%c\n", char)
	if char == 0 {
		return ""
	} else {
		return string(char)
	}
}

// LoopInput 循环输入选择, 或者直接回车退出
func LoopInput(tip string, choices interface{}, print bool) int {
	reflectValue := reflect.ValueOf(choices)
	if reflectValue.Kind() != reflect.Slice && reflectValue.Kind() != reflect.Array {
		fmt.Println("only support slice or array type!")
		return -1
	}
	length := reflectValue.Len()
	if print && reflectValue.Type().String() == "[]string" {
		for i := 0; i < length; i++ {
			fmt.Printf("%d.%s\n\n", i+1, reflectValue.Index(i).Interface())
		}
	}
	for {
		inputString := ""
		if length < 10 {
			inputString = getChar(tip)
		} else {
			fmt.Print(tip)
			_, _ = fmt.Scanln(&inputString)
		}
		if inputString == "" {
			return -1
		} else if !IsNumeric(inputString) {
			fmt.Println("输入有误,请重新输入")
			continue
		}
		number, _ := strconv.Atoi(inputString)
		if number <= length && number > 0 {
			return number
		} else {
			fmt.Println("输入数字越界,请重新输入")
		}
	}
}

// Input 读取终端用户输入
func Input(tip string, defaultValue string) string {
	input := ""
	fmt.Print(tip)
	_, _ = fmt.Scanln(&input)
	if input == "" && defaultValue != "" {
		input = defaultValue
	}
	return input
}

// Red
func Red(str string) string {
	return RED + str + RESET
}

// Green
func Green(str string) string {
	return GREEN + str + RESET
}

// Yellow
func Yellow(str string) string {
	return YELLOW + str + RESET
}

// Blue
func Blue(str string) string {
	return BLUE + str + RESET
}

// Fuchsia
func Fuchsia(str string) string {
	return FUCHSIA + str + RESET
}

// Cyan
func Cyan(str string) string {
	return CYAN + str + RESET
}

// White
func White(str string) string {
	return WHITE + str + RESET
}
