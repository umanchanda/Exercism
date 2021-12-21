package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, num := range birdsPerDay {
		sum += num
	}
	return sum
	// panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	start := (week - 1) * 7
	end := start + 7
	for i := start; i < end; i++ {
		sum += birdsPerDay[i]
	}
	return sum
	// panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	newBirdCount := birdsPerDay
	for i := 0; i < len(birdsPerDay); i += 2 {
		newBirdCount[i] = birdsPerDay[i] + 1
	}
	return newBirdCount
	// panic("Please implement the FixBirdCountLog() function")
}
