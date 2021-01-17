package main

import "fmt"

func main() {

	var coordinates = [][]int{{1, 1}, {2, 2}, {4, 4}, {5, 5}, {6, 6}, {7, 7}}
	var res = checkStraightLine(coordinates)
	fmt.Print(res)
}

func checkStraightLine(coordinates [][]int) bool {
	/*
	 * abc三点，a点和c点的xy的差比值等于b点和c点的比值，则为一条直线
	 * (c.x-a.x)/(c.y-a.y)=(c.x-b.x)/(c.y-b.y)
	 * (c.x-a.x)*(c.y-b.y)=(c.x-b.x)*(c.y-a.y)
	 * c.x*c.y-c.x*b.y-a.x*c.y+a.x*b.y=c.x*c.y-c.x*a.y-b.x*c.y+b.x*a.y
	 * -c.x*b.y-a.x*c.y+c.x*a.y+b.x*c.y=b.x*a.y-a.x*b.y
	 * c.x*(a.y-b.y)+c.y*(b.x-a.x)=b.x*a.y-a.x*b.y
	 */
	var abx int = coordinates[1][0] - coordinates[0][0]
	var aby int = coordinates[0][1] - coordinates[1][1]
	var abxy int = coordinates[1][0]*coordinates[0][1] - coordinates[0][0]*coordinates[1][1]

	for i := 2; i < len(coordinates); i++ {
		if coordinates[i][0]*aby+coordinates[i][1]*abx != abxy {
			return false
		}
	}
	return true
}
