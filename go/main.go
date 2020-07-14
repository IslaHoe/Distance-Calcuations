package main

import "sort"

func main() {

	officeLat := 53.339428
	officeLng := -6.257664
	var radiusDist float64 = 100

	customers := fetchCustomers()
	customers = filterByDistance(customers, officeLat, officeLng, radiusDist)
	sort.Sort(ByID(customers))
	writeToFile(customers)
}
