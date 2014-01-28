gapp
====

Golang thread-safe configuration with notification and cli support.

Usage
-----

```go
package main

import (
  "fmt"
  . "github.com/pack/gapp"
  "reflect"
)

func main() {
  // Config.Add(<name>, <alias>, <description>, <data>, <data_type>, <cli_param>, <required>)
  Config.Add("http_port", "p", "Http web port", 8080, reflect.Int, false, false)
  Config.Add("http_host", "H", "Http hostname", "localhost", reflect.String, false, false)
  Config.Add("redis_port", "r", "Redis port", 6379, reflect.Int, false, false)
  Config.Add("redis_host", "R", "Redis hostname", "localhost", reflect.String, false, false)

  http_entry, _ := Config.Get_entry("p")
  http_desc := http_entry.Description
  http_port, _ := Config.Get("http_port")
  fmt.Println(fmt.Sprintf("%s: %v", http_desc, http_port))

  red_entry, _ := Config.Get_entry("R")
  red_desc := red_entry.Description
  red_port, _ := Config.Get("R")
  fmt.Println(fmt.Sprintf("%s: %s", red_desc, red_port))
}

```

TODO
----

* Add support for application commands
* Add support for cli flags
* Properly handle the case of removing or renaming config entries
