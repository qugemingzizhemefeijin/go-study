package main

import (
	"fmt"
)

//QuickSort2 ...
func QuickSort2(startIndex, endIndex int, arr *[10]int) {
	//递归结束条件：startIndex >= endIndex时
	if startIndex >= endIndex {
		return
	}
	//得到基准元素
	pivot := partition(startIndex, endIndex, arr)
	//根据基准元素，分成两部分进行递归排序
	QuickSort2(startIndex, pivot-1, arr)
	QuickSort2(pivot+1, endIndex, arr)
}

//获取基准元素位置(双边循环法)
func partition(startIndex, endIndex int, arr *[10]int) int {
	//取第1个位置(也可以选择随机位置)的元素作为基准元素
	pivot := arr[startIndex]
	left := startIndex
	right := endIndex

	for left != right {
		//控制right指针比较并左移
		for left < right && arr[right] > pivot {
			right--
		}
		//控制left指针比较并右移[这里必须是小于等于，否则left就卡在第一位了]
		for left < right && arr[left] <= pivot {
			left++
		}
		//交换left和right指向的元素
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	//最后pivot和指针重合点进行交换
	arr[startIndex] = arr[left]
	arr[left] = pivot

	return left
}

//QuickSort ...
//left  表示数组左侧下标
//right 表示数组右侧下标
func QuickSort(left, right int, arr *[10]int) {
	l := left
	r := right
	//pivot 是中轴，支点
	pivot := arr[(left+right)/2]

	//for循环的目标是将比pivot小的数放到左边，大的数放到右边
	for l < r {
		//先从pivot左边拿到比pivot大的数的下标
		for arr[l] < pivot {
			l++
		}
		//再从pivot右边到到比pivot小的数的下标
		for arr[r] > pivot {
			r--
		}
		//l >= r表明本次分解任务完成，退出
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		//在arr[l]和arr[r]都等于pivot，就不需要交换了[此步骤可以省略，仅是优化]
		if arr[l] == arr[r] && arr[l] == pivot {
			r--
			l++
		}
	}

	//如果l与r相等的话，各自移动一位
	if l == r {
		l++
		r--
	}

	if left < r {
		QuickSort(left, r, arr)
	}
	if right > l {
		QuickSort(l, right, arr)
	}
}

//快速排序思想：
//通过一趟比较将要排序的数据分割成独立的两部分，其中一部分的所有数据比另外一部分的所有数据都要小
//然后按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列
func main() {
	arr := [10]int{-9, 78, 0, 23, -567, 70, 80, -74, -1, 998}
	QuickSort2(0, len(arr)-1, &arr)
	fmt.Println(arr)

	fmt.Println("exit success")
}
