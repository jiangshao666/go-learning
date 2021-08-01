package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConf struct {
	Host     string `ini:"host"`
	Port     uint16 `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RedisConf struct {
	Host     string `ini:"host"`
	Master   bool   `ini:"master"`
	Port     uint16 `ini:"port"`
	Database uint16 `ini:"database"`
}

type iniLoader struct {
	confMap map[string]map[string]interface{}
}

func (loader *iniLoader) loadIni(confPath string) bool {
	loader.confMap = make(map[string]map[string]interface{})

	bytes, err := ioutil.ReadFile(confPath)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return false
	}
	lines := strings.Split(string(bytes), "\r\n")
	var sectionName string
	for idx, line := range lines {
		lineStr := strings.TrimSpace(line)
		if strings.HasPrefix(lineStr, ";") || strings.HasPrefix(lineStr, "#") || len(lineStr) == 0 {
			continue
		}
		if strings.HasPrefix(lineStr, "[") { //模块
			if lineStr[0] != '[' || lineStr[len(lineStr)-1] != ']' {
				fmt.Printf("line %d syntax err\n", idx+1)
				return false
			}
			// 取出[]中间的内容，判断是否为空
			sectionName = strings.TrimSpace(lineStr[1 : len(lineStr)-1])
			if len(sectionName) == 0 {
				fmt.Printf("line %d syntax err\n", idx+1)
				return false
			}
			sectionKV := make(map[string]interface{}, 16)
			loader.confMap[sectionName] = sectionKV
		} else { //键值对
			kv := strings.Split(lineStr, "=")
			if len(kv) != 2 || len(kv[0]) == 0 {
				fmt.Printf("line %d syntax err\n", idx+1)
				return false
			}
			loader.confMap[sectionName][kv[0]] = kv[1]
		}
	}
	return true
}

func (loader *iniLoader) mapTo(sectionName string, typeStruct interface{}) error {
	t := reflect.TypeOf(typeStruct)
	fmt.Println(t.Kind())
	if t.Kind() != reflect.Ptr {
		err := errors.New("mapto type must be pointer")
		return err
	}
	if t.Elem().Kind() != reflect.Struct {
		err := errors.New("mapto type must be struct pointer")
		return err
	}
	sectionKV, ok := loader.confMap[sectionName]
	if !ok {
		err := fmt.Errorf("section of %s not exist", sectionName)
		return err
	}

	v := reflect.ValueOf(typeStruct)

	for i := 0; i < t.Elem().NumField(); i++ {
		elemField := t.Elem().Field(i)
		fieldName := elemField.Name
		keyName := elemField.Tag.Get("ini")
		confVal := sectionKV[keyName]
		confValStr := confVal.(string)

		//fmt.Println(fieldName, keyName, confValStr, elemField.Type.Kind())
		switch elemField.Type.Kind() {
		case reflect.String:
			v.Elem().FieldByName(fieldName).SetString(confValStr)
		case reflect.Bool:
			val, err := strconv.ParseBool(confValStr)
			if err != nil {
				fmt.Printf("section parse value %s failed", confValStr)
				return err
			}
			v.Elem().FieldByName(fieldName).SetBool(val)
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint,
			reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val, err := strconv.ParseUint(confValStr, 10, 64)
			if err != nil {
				fmt.Printf("section parse value %s failed", confValStr)
				return err
			}
			v.Elem().FieldByName(fieldName).SetUint(val)
		}
	}
	return nil
}

func main() {
	var loader iniLoader
	ok := loader.loadIni("./conf.ini")
	if !ok {
		fmt.Println("loadIni failed")
		return
	}
	fmt.Println(loader.confMap)
	var mysql MysqlConf
	err := loader.mapTo("mysql", &mysql)
	if err != nil {
		fmt.Printf("mapTo failed , err:%v", err)
	}
	fmt.Printf("mysql: %#v \n", mysql)
	var redis RedisConf
	err = loader.mapTo("redis", &redis)
	if err != nil {
		fmt.Printf("mapTo failed , err:%v", err)
	}
	fmt.Printf("redis: %#v \n", redis)
}
