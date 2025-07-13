package graph_algo

import (
	"fmt"
	"math"
	"testing"
)

func calcPath(eqArr [][]string, valArr []float64) (map[string]int, [][]float64) {
	serial := 0
	idMap := make(map[string]int)

	for _, eq := range eqArr {
		u, v := eq[0], eq[1]

		_, uOK := idMap[u]
		if !uOK {
			idMap[u] = serial
			serial++
		}
		_, vOK := idMap[v]
		if !vOK {
			idMap[v] = serial
			serial++
		}
	}

	wMatrix := make([][]float64, serial)
	for idx := range wMatrix {
		wMatrix[idx] = make([]float64, serial)
		for jdx := range serial {
			wMatrix[idx][jdx] = math.MaxInt
		}
	}

	for idx, eq := range eqArr {
		u, v := eq[0], eq[1]
		uId, vId := idMap[u], idMap[v]
		wMatrix[uId][vId] = valArr[idx]
		wMatrix[vId][uId] = 1 / valArr[idx]
	}

	// unlike floyed-warshall-SP, we are trying floyed-warshall-matrix
	for k := 0; k < serial; k++ {
		for i := 0; i < serial; i++ {
			for j := 0; j < serial; j++ {
				if wMatrix[i][k] != math.MaxInt && wMatrix[k][j] != math.MaxInt {
					wMatrix[i][j] = wMatrix[i][k] * wMatrix[k][j]
					wMatrix[j][i] = 1 / wMatrix[i][j]
				}
			}
		}
	}

	return idMap, wMatrix
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	idMap, wMatrix := calcPath(equations, values)
	ret := make([]float64, len(queries))

	//fmt.Println(idMap)

	for idx, query := range queries {
		u, v := query[0], query[1]
		uId, uOK := idMap[u]
		vId, vOK := idMap[v]
		if !uOK || !vOK {
			//fmt.Println("# -1 for id not found", u, ":", uId, v, ":", vId)
			ret[idx] = -1.0
			continue
		}
		if wMatrix[uId][vId] == math.MaxInt {
			//fmt.Println("#-1 for path not found: ", u, v)
			ret[idx] = -1.0
			continue
		}
		ret[idx] = wMatrix[uId][vId] //roundFloat(wMatrix[uId][vId], 5)
	}
	return ret
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func TestDiv(t *testing.T) {
	eqArr := [][]string{{"x2", "x1"}, {"x2", "x3"}, {"x1", "x4"}, {"x2", "x5"}}
	valArr := []float64{1e-05, 1e-05, 3.4, 5.6}
	queryArr := [][]string{{"x1", "x3"}} // {"x2", "x4"}, {"x3", "x4"}, {"x4", "x3"}

	result := calcEquation(eqArr, valArr, queryArr)
	fmt.Println(result)
}
