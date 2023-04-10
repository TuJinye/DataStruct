package main

import "fmt"

const INF = 1000000000 // 定义一个足够大的值代表无穷大

func floydWarshall(graph [][]int, n int) [][]int {
	// 初始化距离矩阵为原图
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		copy(dist[i], graph[i])
	}

	// 逐步更新距离矩阵，k代表中间节点
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// 如果经过中间节点k可以得到更短的路径，则更新距离矩阵
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

func main() {
	// 测试数据
	graph := [][]int{
		{0, 5, INF, 10},
		{INF, 0, 3, INF},
		{INF, INF, 0, 1},
		{INF, INF, INF, 0},
	}

	// 计算所有节点之间的最短路径
	dist := floydWarshall(graph, 4)

	// 输出距离矩阵
	fmt.Println("距离矩阵：")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if dist[i][j] == INF {
				fmt.Print("INF ")
			} else {
				fmt.Printf("%d ", dist[i][j])
			}
		}
		fmt.Println()
	}
}
