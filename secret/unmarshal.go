package secret

import (
	"encoding/base64"
	"fmt"
	"reflect"

	corev1 "k8s.io/api/core/v1"
)

// Unmarshal a Secret to a struct.
func Unmarshal(s *corev1.Secret, obj interface{}) error {
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
			val = elm
		}
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get(Tag)

		if d, ok := s.Data[tag]; ok {
			if val.Field(i).Kind() == reflect.String {
				dec, err := base64.StdEncoding.DecodeString(string(d))
				if err != nil {
					return fmt.Errorf("failed to unmarshal %s with value %s to type string: %s", tag, d, err)
				}
				val.Field(i).SetString(string(dec))
			}
		}
	}

	return nil
}
