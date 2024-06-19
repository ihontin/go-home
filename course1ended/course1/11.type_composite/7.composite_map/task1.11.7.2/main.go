package main

// mergeMaps , которая принимает две карты map1 и map2 типа map[string]int
// и объединяет их в одну карту. Функция должна вернуть объединенную карту.
func mergeMaps(map1, map2 map[string]int) map[string]int {
	var complexMaps = make(map[string]int)
	for key, val := range map1 {
		complexMaps[key] = val
	}
	for key, val := range map2 {
		complexMaps[key] = val
	}
	return complexMaps
}

//func main() {
//	map1 := map[string]int{"apple": 3, "banana": 2}
//	map2 := map[string]int{"orange": 5, "grape": 4}
//	mergedMap := mergeMaps(map1, map2)
//	for key, value := range mergedMap {
//		fmt.Printf("%s: %d\n", key, value)
//	}
//}
