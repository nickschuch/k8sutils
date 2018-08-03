K8s Utils
=========

Tools for a nicer Kubernetes development experience.

## Tools

### Unmarshal a Kubernetes ConfigMap to a Struct

```go
package main

import (
    "fmt"

    "github.com/previousnext/mysql-toolkit/stuff/k8s/configmap"
)

type Backend struct {
    Host string `k8s_configmap:"backend.host"`
    Port string `k8s_configmap:"backend.port"`
}

func main() {
    cfg := corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "example",
			Name: "conf",
		},
		Data: map[string]string{
			"backend.host": "1.1.1.1",
			"backend.port": "443",
		},
	}

    err = configmap.Unmarshal(clientset, namespace, name, &b)
    if err != nil {
        panic(err)
    }

    fmt.Println("Host:", b.Host)
    fmt.Println("Port:", b.Port)
}
```
