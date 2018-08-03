package configmap

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Name of the struct tag used in examples
const tagName = "k8s_configmap"

// Unmarshal a ConfigMap to a Go struct.
func Unmarshal(clientset *kubernetes.Clientset, namespace, name string, obj interface{}) error {
	cfg, err := clientset.CoreV1().ConfigMaps(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return errors.Wrap(err, "failed to lookup ConfigMap")
	}

	return unmarshal(cfg.Data, obj)
}

// Helper function for unmarshalling data to a Go struct.
func unmarshal(data map[string]string, obj interface{}) error {
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
		tag := val.Type().Field(i).Tag.Get(tagName)

		if d, ok := data[tag]; ok {
			switch val.Field(i).Kind() {
			case reflect.String:
				val.Field(i).SetString(d)

			case reflect.Int64:
				v, err := strconv.ParseInt(d, 10, 64)
				if err != nil {
					return fmt.Errorf("failed to unmarshal %s with value %s to type int: %s", tag, d, err)
				}
				val.Field(i).SetInt(v)

			case reflect.Bool:
				v, err := strconv.ParseBool(d)
				if err != nil {
					return fmt.Errorf("failed to unmarshal %s with value %s to type bool: %s", tag, d, err)
				}
				val.Field(i).SetBool(v)
			}
		}
	}

	return nil
}
