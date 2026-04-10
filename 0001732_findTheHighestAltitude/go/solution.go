package solution

func LargestAltitude(gain []int) int {
	altitude := 0
	highest := 0

	for _, point := range gain {
		altitude = altitude + point
		highest = max(highest, altitude)
	}

	return highest
}
