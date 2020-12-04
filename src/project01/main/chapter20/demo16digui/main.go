package main

import "fmt"

func print(myMap *[8][7]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf("%d ", myMap[i][j])
		}
		fmt.Println()
	}
}

//SetWay 编写函数，完成找路
//myMap地图，要保证地图是同一份，因此使用的是引用
//i,j表示对地图的哪个点进行探测
func SetWay(myMap *[8][7]int, i, j int) bool {
	//分析出什么情况下，找出通路
	//myMap[6][5] == 2
	if myMap[6][5] == 2 {
		return true
	} else {
		//说明要继续找
		if myMap[i][j] == 0 { //如果这个点是可以探测的
			//假设myMap[i][j]这个点是可以通的，但是需要探测 下右上左 [上下左右不行]
			myMap[i][j] = 2
			if SetWay(myMap, i+1, j) { //下
				return true
			} else if SetWay(myMap, i, j+1) { //右
				return true
			} else if SetWay(myMap, i-1, j) { //上
				return true
			} else if SetWay(myMap, i, j-1) { //左
				return true
			} else {
				//说明是一个死路
				myMap[i][j] = 3
				return false
			}
		} else { //说明这个点是一堵墙
			return false
		}
	}
}

//迷宫地图回溯
// ==========
// =+       =
// ===      =
// =        =
// =        =
// =       E=
// =========
func main() {
	//先创建一个二维数组，模拟迷宫
	//规则
	//1. 如果元素的值为1，表示墙
	//2. 如果元素为0，则代表没有探过
	//3. 如果元素值为2，则表示一个通路
	//4. 如果元素的值为3，是走过的点，但是走不通
	var myMap [8][7]int

	//先把地图的最上和最下设置为1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	//再把地图的最左和最右设置为1
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}

	//设置3,1和3,2的墙
	myMap[3][1] = 1
	myMap[3][2] = 1

	//输出地图
	print(&myMap)

	SetWay(&myMap, 1, 1)
	fmt.Println("探测完毕...")
	//输出地图
	print(&myMap)

	fmt.Println("exit success")
}
