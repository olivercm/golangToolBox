package main

import (
	"fmt"
	"github.com/golang/groupcache/consistenthash"
)

//reference https://mp.weixin.qq.com/s/Wya4QBECtJLf3P8ALUPFuw
func main() {
	// 构造一个 consistenthash 对象，每个节点在 Hash 环上都一共有三个虚拟节点。
	hash := consistenthash.New(3, nil)

	// 添加节点
	hash.Add(
		"127.0.0.1:8080",
		"127.0.0.1:8081",
		"127.0.0.1:8082",
	)

	// 根据 key 获取其对应的节点
	node := hash.Get("cyhone.com")

	fmt.Println(node)
}
