package main

import "fmt"

func main() {
    // 定义一个 5x5 的特殊矩阵
    mat := [][]int{
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0},
    }

    // 设置矩阵中的非 0 元素
    mat[1][2] = 2
    mat[3][4] = 4

    // 输出原始矩阵
    fmt.Println("Original Matrix:")
    for i := 0; i < len(mat); i++ {
        fmt.Println(mat[i])
    }
    // Output:
    // Original Matrix:
    // [0 0 0 0 0]
    // [0 0 2 0 0]
    // [0 0 0 0 0]
    // [0 0 0 0 4]
    // [0 0 0 0 0]

    // 计算矩阵中非 0 元素的个数
    var count int
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[i]); j++ {
            if mat[i][j] != 0 {
                count++
            }
        }
    }

    // 创建稀疏矩阵的存储结构
    sparseMat := make([][]int, count+1)
    sparseMat[0] = []int{len(mat), len(mat[0]), count}

    // 将非 0 元素存储到稀疏矩阵中
    var k int
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[i]); j++ {
            if mat[i][j] != 0 {
                k++
                sparseMat[k] = []int{i, j, mat[i][j]}
            }
        }
    }

    // 输出稀疏矩阵
    fmt.Println("Sparse Matrix:")
    for i := 0; i < len(sparseMat); i++ {
        fmt.Println(sparseMat[i])
    }
    // Output:
    // Sparse Matrix:
    // [5 5 2]
    // [1 2 2]
    // [3 4 4]
}
