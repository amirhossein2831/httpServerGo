package main

import "github.com/amirhossein2831/httpServerGo/src/Jobs"

func main() {
	sm := Jobs.NewSimpleMessageJob()
	sm.Publish()
	sm.Consume()
}
