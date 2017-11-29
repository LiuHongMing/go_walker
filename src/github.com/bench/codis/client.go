package main

import (
	"net"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

// 可见性规则：开头大写公有访问，小写私有访问

type Protocol string

type ClientProxy struct {

	Protocol

	host, port string

}

func (c ClientProxy) getProxyAddress() string {
	return net.JoinHostPort(c.host, c.port)
}

func main() {

	// 创建值对象
	client := ClientProxy{"tcp","175.63.101.125", "19000"}

	// net
	hostAndPort := client.getProxyAddress()
	conn, err := redis.Dial(string(client.Protocol), hostAndPort)

	if err == nil {

		// info
		reply, err := conn.Do("info")

		if err == nil {
			res, err := redis.String(reply, err)
			if err == nil {
				msg := "res=" + res
				fmt.Println(msg)
			}
		} else {
			fmt.Print(err)
		}

		// select
		reply, err = conn.Do("select", "1")

		if err == nil {
			res, err := redis.String(reply, err)
			if err == nil {
				msg := "It's OK, res=" + res
				fmt.Println(msg)
			}
		} else {
			fmt.Print(err)
		}

		// set
		reply, err = conn.Do("set", "hello2", "world")

		if err == nil {
			res, err := redis.String(reply, err)
			if err == nil {
				msg := "It's OK, res=" + res
				fmt.Print(msg)
			}
		} else {
			fmt.Println(err)
		}

	}


}