package Server

import "servel/Container"

func GetApps(environment *map[string]Container.Environment) {
	(*environment)["localhost"] = Container.Environment{Exec: "node", Path: "/Users/*/node/srvles", Port: "3000", Index: "index.js"}
}
