package main

import (
	"fmt"
	"reflect"
)

type BackUpKeyList struct {
	T1        BackUpKeyListItem `json:"T1"`
	T2        BackUpKeyListItem `json:"T2"`
	T3        BackUpKeyListItem `json:"T3"`
	T3V2      BackUpKeyListItem `json:"T3V2"`
	SIT       BackUpKeyListItem `json:"SIT"`
	STAGING   BackUpKeyListItem `json:"STAGING"`
	STAGINGV2 BackUpKeyListItem `json:"STAGINGV2"`
	RC        BackUpKeyListItem `json:"RC"`
	BETA      BackUpKeyListItem `json:"BETA"`
	PROD      BackUpKeyListItem `json:"PROD"`
	PRODV2    BackUpKeyListItem `json:"PRODV2"`
}

// LastVersion 也在 BUVListTargetKey 中
// BUVListTargetKey 应该是 backup-{date}
type BackUpKeyListItem struct {
	BUVListTargetKey     []string `json:"versionList"`
	LastVersionTargetKey string   `json:"lastVersion"`
}

func main() {

	backupItem := &BackUpKeyListItem{
		BUVListTargetKey:     []string{"go", "python", "java"},
		LastVersionTargetKey: "go",
	}

	backup := &BackUpKeyList{T3: *backupItem}
	fmt.Printf("backup=%v\n", backup) // prints 1

	immutable := reflect.ValueOf(backup)
	immutable = immutable.Elem()

	val := immutable.FieldByName("T3")
	fmt.Printf("backup=%v\n", val) // prints 1

	// TODO 转类型
	a := val.Interface().(BackUpKeyListItem)
	// strArr := val.FieldByName("BUVListTargetKey")
	fmt.Printf("backup=%v\n", a.LastVersionTargetKey) // prints 1
	// fmt.Printf("backup=%v\n", strArr.FieldByIndex(0)) // prints 1
}
