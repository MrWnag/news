package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
)

type RedisGit struct {
	beego.Controller
}

func(this*RedisGit)ShowRedis(){
	//链接数据库
	conn ,err:= redis.Dial("tcp",":6379")
	if err != nil{
		beego.Error("redis链接错误",err)
		return
	}

	//操作函数
	resp,err :=conn.Do("mget","kk","s6","ll")

	//回复助手函数  类型转换
	result,_:=redis.Values(resp,err)

	var v1,v2 string
	var v3 int
	redis.Scan(result,&v1,&v2,&v3)

	beego.Info(v1,v2, v3)




	//conn.Send("set","kk","value")
	//conn.Flush()
	//conn.Receive()
}