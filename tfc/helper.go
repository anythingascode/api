package tfc

import "github.com/go-test/deep"

func toMap(in interface{}) map[string]interface{} {
	if in != nil {
		return in.(map[string]interface{})
	}
	return nil
}

func isEqual(key string, before, after map[string]interface{}) bool {
	if _, ok := before[key]; ok && deep.Equal(before[key], after[key]) != nil && before[key] != nil && after[key] != nil {
		return false
	}
	return true
}
