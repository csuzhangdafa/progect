package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// shoppcatmanage实现购物车管理
type ShoppingCarManage struct {
	cart   map[string]int
	keys   []string //用于保存插入顺序的键
	mutext sync.Mutex
}

var once sync.Once
var instance *ShoppingCarManage

// 获取购物车实例(申明一个实例)
func getInstance() *ShoppingCarManage {
	once.Do(func() {
		instance = &ShoppingCarManage{
			cart: make(map[string]int),
		}
	})
	return instance
}

// 添加商品到购物车检查itemName是否存在于scm.cart这个map中，
// 如果不存在，则执行if语句块中的代码
func (scm *ShoppingCarManage) addToCart(itemName string, quantity int) {
	scm.mutext.Lock()
	defer scm.mutext.Unlock()
	if _, exits := scm.cart[itemName]; !exits {
		scm.keys = append(scm.keys, itemName)
	}
	scm.cart[itemName] += quantity
}

// viewCart查看购物车并且安按照顺序输出
func (scm *ShoppingCarManage) viewCart() {
	scm.mutext.Lock()
	defer scm.mutext.Unlock()

	for _, itemName := range scm.keys {
		quantity := scm.cart[itemName]
		fmt.Printf("%s %d\n", itemName, quantity)
	}
}

func main() {
	cart := getInstance()
	scanner := bufio.NewScanner(os.Stdin)
	/*创建了一个新的Scanner对象，用于从标准输入（通常是键盘输入）读取并解析文本。
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("Read:", line)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from standard input:", err)
		}
	}
	*/

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		parts := strings.Fields(input)
		itemName := parts[0]
		quantity := 0
		if len(parts) > 1 {
			fmt.Scanf(parts[1], "%d", &quantity)
		}

		//获取购物车并添加实例
		cart.addToCart(itemName, quantity)
	}
	//输出购物车内容
	cart.viewCart()
}
