package utils

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/741369/go_utils/log"

	"github.com/bwmarrin/snowflake"
)

func InArray(str string, data []string) bool {
	for _, v := range data {
		if str == v {
			return true
		}
	}
	return false
}

// 获取语言
func GetLanguage(language string) string {
	languageLower := strings.ToLower(language)

	// 繁体中文返回英文，其中繁体中文对应的字符串为：zh-hk,zh-tw,zh-mo,zh-Hant
	reg := regexp.MustCompile(`^zh-hk|^zh-tw|^zh-mo|^zh-hant.*`)
	if reg.MatchString(languageLower) {
		return "en"
	}

	splitLanguage := strings.Split(languageLower, ",")
	regEn := regexp.MustCompile("^en-.*")
	if regEn.MatchString(splitLanguage[0]) {
		return "en"
	}
	regZh := regexp.MustCompile("^zh-.*")
	if regZh.MatchString(splitLanguage[0]) {
		return "zh"
	}
	regTh := regexp.MustCompile("^th-.*")
	if regTh.MatchString(splitLanguage[0]) {
		return "th"
	}
	return "zh"
}

//base64 生成编码
func Base64Encode(content string) string {
	input := []byte(content)
	return base64.StdEncoding.EncodeToString(input)
}

/**
JsonEncode
*/
func JsonEncode(data interface{}) string {
	encode, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(encode)
}

// 生成随机码
func GenerateCode(randomLength, randomNum int) []string {
	privilegeArr := make([]string, randomNum)
	privilegeTmp := make([]rune, randomLength)
	//var letters = []rune("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789")
	var letters = []rune("AaBbCcDdEeFfGgHhJjKkMmNnPpQqRrSsTtUuVvWwXxYyZz23456789")
	randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < randomNum; i++ {
		for j := 0; j < randomLength; j++ {
			privilegeTmp[j] = letters[randSeed.Intn(len(letters))]
		}
		privilegeArr[i] = string(privilegeTmp)
	}
	return privilegeArr
}

// 生成雪花 id
func GetSnowId() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Infof(nil, "generate snowflake id error, err = %s", err)
		return 0
	}
	return node.Generate().Int64()
}

// 生成雪花 id
func GetSnowIdString() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Infof(nil, "generate snowflake id error, err = %s", err)
		return "0"
	}
	return node.Generate().String()
}

// 分页查询数据, return offset, limit
func GetPagePageSize(param map[string]string) (int, int) {
	page, pageSize := 1, 20
	if param["page"] != "" {
		pageTmp, err := strconv.Atoi(param["page"])
		if err == nil {
			page = pageTmp
		}
	}
	if page <= 0 {
		page = 1
	}

	if param["page_size"] != "" {
		pageSizeTmp, err := strconv.Atoi(param["page_size"])
		if err == nil {
			pageSize = pageSizeTmp
		}
	}

	// 设置一页最大20条
	if pageSize < 0 || pageSize > 20 {
		pageSize = 20
	}

	return (page - 1) * pageSize, pageSize
}
