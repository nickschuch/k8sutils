K8s Utils
=========

Tools for a nicer Kubernetes development experience.

## Tools

### Unmarshal a Kubernetes ConfigMap to a Struct

```go
package main

import (
    "fmt"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"

    "github.com/previousnext/mysql-toolkit/stuff/k8s/configmap"
)

type Backend struct {
    Host string `k8s_configmap:"backend.host"`
    Port string `k8s_configmap:"backend.port"`
}

func main() {
    var b Backend

    var (
        namespace = "default"
        name      = "example"
    )

    config, err := clientcmd.BuildConfigFromFlags("", "/home/USER/.kube/config")
    if err != nil {
        panic(err)
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err)
    }

    err = configmap.Unmarshal(clientset, namespace, name, &b)
    if err != nil {
        panic(err)
    }

    fmt.Println("Host:", b.Host)
    fmt.Println("Port:", b.Port)
}
```
