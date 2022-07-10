package util

import (
	"fmt"
	"strings"
)

func Exits(dict map[string]string, target string, splitPoint string, index int64) (string, bool) {

	for i, v := range dict {
		if strings.Split(i, splitPoint)[index] == target {
			return v, true
		}
	}

	return "", false
}

func Merge(dict map[string]string, t map[string]string) map[string]string {

	if t == nil {
		t = make(map[string]string)
	}

	for i, v := range dict {
		dataT := make(map[string]string)
		if len(t[i]) != 0 {
			dataT = ConvertStringToMap(t[i])
		}
		dataD := ConvertStringToMap(v)
		for dataI, dataV := range dataD {
			dataT[dataI] = dataV
		}
		newConfig := ""
		for ii, vv := range dataT {
			newConfig = newConfig + fmt.Sprintf("%v=%v\n", ii, vv)
		}
		t[i] = newConfig[:len(newConfig)-1]
	}
	return t
}

func ConvertStringToMap(s string) map[string]string {
	d := make(map[string]string)
	data := strings.Split(s, "\n")
	for _, v := range data {
		index := strings.Index(v, "=")
		if index == len(v) {
			continue
		}
		d[v[:index]] = v[index+1:]
	}
	return d
}
