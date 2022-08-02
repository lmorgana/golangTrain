package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func randomInt(max, min int) int {
	//rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func fillMatrixRandom(matrix [][]int, max, min int) [][]int {
	for i, row := range matrix {
		for j, _ := range row {
			matrix[i][j] = randomInt(max, min)
		}
	}
	return matrix
}

func strContainsOnly(s1, chars string) bool {
	for _, valS := range s1 {
		if !strings.Contains(s1, string(valS)) {
			return false
		}
	}
	return true
}

func getIntSliceFromString(s1 string) []int {
	arr := make([]int, len(s1)-1)
	for i := 0; i < len(s1)-1; i++ {
		arr[i], _ = strconv.Atoi(string(s1[i]))
	}
	return arr
}

func fillMatrixByFile(fileName string) ([][]int, error) {
	var matrix [][]int
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	lenLine := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if lenLine == 0 {
			lenLine = len(line)
		}
		if len(line) < 1 || lenLine != len(line) {
			return nil, errors.New("ReadFile: not valid len of string")
		}
		if !strContainsOnly(line, "01\n") {
			return nil, errors.New("ReadFile: unknown char")
		}
		matrix = append(matrix, getIntSliceFromString(line))
	}
	return matrix, nil
}

func printMatrix(matrix [][]int) {
	for _, str := range matrix {
		fmt.Println(str)
	}
}

func makeMatrix(row, column int) [][]int {
	if row > 0 && column > 0 {
		matrix := make([][]int, row)
		for i, _ := range matrix {
			matrix[i] = make([]int, column)
		}
		return matrix
	}
	return nil
}

func numbStrMaxCoat(matrix [][]int) (max, nMax int) {
	//строка с максимальным покрытием
	for i, row := range matrix {
		sum := 0
		for _, val := range row {
			sum += val
		}
		if sum == len(matrix) {
			return sum, i
		} else if sum > max {
			max = sum
			nMax = i
		}
	}
	return max, nMax
}

func getSlice(newMatrix []int, matrix []int, mask []int) {
	j := 0
	for i, val := range matrix {
		if mask[i] == 0 {
			newMatrix[j] = val
			j++
		}
	}
}

func matrixWithoutInter(matrix [][]int, numbStrMaxCoat, maxCoat int) [][]int {
	if len(matrix) > 0 {
		newMatrix := makeMatrix(len(matrix)-1, len(matrix[0])-maxCoat)
		if newMatrix != nil {
			j := 0
			for i, _ := range matrix {
				if i == numbStrMaxCoat {
					continue
				}
				getSlice(newMatrix[j], matrix[i], matrix[numbStrMaxCoat])
				j++
			}
			return newMatrix
		}
	}
	return nil
}

func matrixPower(matrix [][]int) (result int) {
	//слайс с минимальным количеством номеров строк для покрытия матрицы
	for matrix != nil {
		maxCoat, numbStrMaxCoat := numbStrMaxCoat(matrix)
		result++
		if maxCoat == len(matrix[0]) {

		}
		if maxCoat == 0 {
			return -1
		}
		matrix = matrixWithoutInter(matrix, numbStrMaxCoat, maxCoat)
	}
	return result
}

func printResults(matrixPower int) {
	if matrixPower == -1 {
		fmt.Println("невозвожно найти мощность")
	} else {
		fmt.Println("минимальная мощность: ", matrixPower)
	}
}

func main() {
	var matrix [][]int
	arguments := os.Args
	if len(arguments) > 1 {
		bl, err := regexp.MatchString(".txt$", os.Args[1])
		if !bl || err != nil {
			fmt.Println("Для заполнения матрицы из файла соблюдайте правила:\n1)Никаких пустых строк,\n2)Никак символов кроме 0 и 1\n3)Разрешение .txt")
			os.Exit(1)
		}
		matrix, err = fillMatrixByFile(os.Args[1])
		if !bl || err != nil {
			fmt.Println("Для заполнения матрицы из файла соблюдайте правила:\n1)В конце пустая строка," +
				"\n2)Никак символов кроме 0 и 1\n3)Разрешение .txt")
			fmt.Println(err.Error())
			os.Exit(1)
		}
	} else {
		matrix = makeMatrix(4, 4)
		matrix = fillMatrixRandom(matrix, 2, 0)
	}
	//printMatrix(matrix)
	matrixPower := matrixPower(matrix)
	printResults(matrixPower)
}
