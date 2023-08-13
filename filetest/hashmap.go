package main

import (
	"fmt"
	"hash/crc32"
	"runtime/debug"
	"sort"
	"sync"
	"unsafe"

	v2 "sigs.k8s.io/controller-tools/pkg/version"
)

type rindex []uint32 //hash环索引

type ring struct {
	rmap      map[uint32]string //环结构体
	rindexarr rindex            //索引
	lock      *sync.RWMutex     //线程安全
}

// 比大小
func (this rindex) Less(i, j int) bool {
	return this[i] < this[j]
}

// 长度
func (this rindex) Len() int {
	return len(this)
}

// 交换
func (this rindex) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *ring) Addnode(nodename string) {
	this.lock.Lock()
	defer this.lock.Unlock()

	index := crc32.ChecksumIEEE([]byte(nodename)) //sha256
	if _, ok := this.rmap[index]; ok {
		return //返回
	}
	this.rmap[index] = nodename                    //赋值
	this.rindexarr = append(this.rindexarr, index) //加载索引
	sort.Sort(this.rindexarr)                      //排序

}
func (this *ring) Removenode(nodename string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	index := crc32.ChecksumIEEE([]byte(nodename)) //sha256
	if _, ok := this.rmap[index]; !ok {
		return //返回
	}
	delete(this.rmap, index) //删除map内置数据
	this.rindexarr = rindex{}
	for k := range this.rmap {
		this.rindexarr = append(this.rindexarr, k) //插入数据
	}
	sort.Sort(this.rindexarr) //排序
}
func (this *ring) Getnode(nodename string) string {
	this.lock.RLock() //其他线程可以读取，不可以修改
	defer this.lock.RUnlock()
	//查找
	hash := crc32.ChecksumIEEE([]byte(nodename))

	i := sort.Search(len(this.rindexarr), func(i int) bool {
		return this.rindexarr[i] == hash
	})
	if i < 0 || i > len(this.rindexarr)-1 {
		return ""
	}
	node := this.rmap[this.rindexarr[i]] //取得节点
	return node

}

var Version string

func main() {

	var b bool
	fmt.Println(unsafe.Sizeof(b)) // 输出 1

	info, ok := debug.ReadBuildInfo()
	if !ok {
		// binary has not been built with module support
		fmt.Println("no ok")
	}
	v2.Print()
	fmt.Println("lhb2", info.Main.Version)
	filelist := []string{"123", "456", "789"}
	hashmap := &ring{map[uint32]string{}, rindex{}, new(sync.RWMutex)}
	fmt.Println(filelist, hashmap)
	for _, v := range filelist {
		index := crc32.ChecksumIEEE([]byte(v))               //循环索引
		hashmap.rmap[index] = v                              //设定索引
		hashmap.rindexarr = append(hashmap.rindexarr, index) //插入

	}
	//处理索引数组
	sort.Sort(hashmap.rindexarr)
	fmt.Println(hashmap)

	//fmt.Println(hashmap.Getnode("xadasd"))

	hashmap.Addnode("xyz123")
	fmt.Println(hashmap)

	hashmap.Removenode("xyz123")
	fmt.Println(hashmap)

	mystr := hashmap.Getnode("789123")
	if mystr == "" {
		fmt.Println("找不到")
	} else {
		fmt.Println("找到", mystr)
	}
}
