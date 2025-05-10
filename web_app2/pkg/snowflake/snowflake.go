package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	_ "os"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineId int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}

	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineId)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}

func main() {
	if err := Init("2020-07-01", 1); err != nil {
		fmt.Println("init err:", err)
		return
	}
	id := GenID()
	fmt.Println(id)
}
