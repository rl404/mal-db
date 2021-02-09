package utils

import "strings"

// ArrayFilter to remove empty string from slice.
func ArrayFilter(s []string) []string {
	if s == nil {
		return nil
	}

	r := []string{}
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// InArrayStr to get if value is in array string.
func InArrayStr(arr []string, v string) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}

// InArrayInt to check if int value is in array int.
func InArrayInt(array []int, value int) bool {
	for _, a := range array {
		if a == value {
			return true
		}
	}
	return false
}

// GetMapKey to get key from map.
func GetMapKey(m map[int]string, value string) (key int, found bool) {
	for i, v := range m {
		if strings.EqualFold(v, value) {
			return i, true
		}
	}
	return 0, false
}

// UniqueInt to remove duplicate in slice.
func UniqueInt(list []int) []int {
	if list == nil {
		return nil
	}

	unique := []int{}
	keys := make(map[int]bool)
	for _, v := range list {
		if _, value := keys[v]; !value {
			keys[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}
