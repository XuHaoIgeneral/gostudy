package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {

	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()

	conn,err:=redis.Dial("tcp","127.0.0.1:6379")
	if err!=nil{
		fmt.Println("Connect to redis error",err)
		return
	}

	_,err=conn.Do("SET","test","python")

	if err!=nil{
		fmt.Println("redis set failed",err)
	}

	username,err:=redis.String(conn.Do("GET","makey"))
	if err!=nil{
		fmt.Println("redis get fauled",err)
	}else {
		fmt.Println("Get mykey",username)
	}
	defer conn.Close()
}
