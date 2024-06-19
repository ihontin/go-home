package main

//Пример 1: Без принципа DRY и с ним
//В этом примере у нас есть две функции, которые выполняют одну и ту же логику - вычисление суммы двух чисел.
//	Они дублируют код и усложняют поддержку.

//func calculateSum2(a int, b int) int {
//	return a + b
//}
//
//func calculateSum3(a, b, c int) int {
//	return a + b
//}
//
//func main() {
//	result1 := calculateSum2(2, 3)
//	result2 := calculateSum3(5, 7, 9)
//	// ...
//}

//Применим принцип DRY.
//В этом исправленном примере мы нашли общую логику вычисления суммы и вынесли ее в отдельную функцию,
//чтобы избежать дублирования кода. Теперь код стал более поддерживаемым и избежал дублирования.
//
//То есть, метод dry стремится к тому чтобы, в конечном итоге у нас был общий случай,
//избавляясь от частных случаев в данном примере. Так же, надо взять во внимание переиспользование методов в структурах,
//когда в части, какого либо метода вы переиспользуете логику другого метода.
//func calculateSum(xs ...int) int {
//	var res int
//	for _, x := range xs {
//		res += x
//	}
//
//	return res
//}
//
//func main() {
//	result1 := calculateSum(2, 3)
//	result2 := calculateSum(5, 7)
//	// ...
//}

//Пример 2: Без принципа KISS и с ним
//В этом примере у нас есть функция, которая реализует сложную формулу с множеством условий и операций.
//	Код становится сложным и трудночитаемым.

//func getAverageROI(buys []float64, sells []float64) int {
//	var sum float64
//	var averageBuys float64
//	for i := range buys {
//		sum += buys[i]
//	}
//	averageBuys = sum / len(buys)
//
//	sum = 0
//	var averageSells float64
//	for i := range sells {
//		sum += sells[i]
//	}
//	averageSells = sum / len(sells)
//
//	return averageSells - averageBuys
//}

//Применим принцип KISS.
//В этом исправленном примере мы использовали более простой стиль без излишней сложности.
//	Логика разбита, на мелкие более простые функции, нежели всю логику хранить в одной большой функции.
//	Код стал более понятным и легким для поддержки.

//func getAverageROI(buys []float64, sells float64) int {
//	return average(sells) - average(buys)
//}
//
//func average(xs []float64) float64 {
//	return sum(xs) / len(xs)
//}
//
//func calculateSum(xs ...float64) int {
//	var res int
//	for _, x := range xs {
//		res += x
//	}
//
//	return res
//}

//Пример 3: Без принципа YAGNI и с ним
//Например нам дали задачу высчитать средний доход с продажи товара. Мы решили сделать функцию,
//мы добавили еще некоторые показатели в функцию. Но в итоге нам понадобился только средний доход.
//Мы потратили время на реализацию лишней функциональности.

//func getProductStatistic(buys ...float64, sells ...float64) (float64, float64, float64, float64) {
//	// Логика вычисления суммы и произведения
//	// ...
//}

//Применим принцип YAGNI.
//В этом исправленном примере мы не реализовали лишнюю функциональность, а сделали функцию, которая принимает
//в качестве параметров функции, которые будут вычислять необходимые показатели. Таким образом мы можем использовать
//только необходимые нам функции. Так же, мы сделали код более гибким, теперь мы можем добавлять новые функции
//для вычисления показателей, не меняя сигнатуру функции getProductStatistic.

//package main
//
//import "fmt"
//
//func AverageROI(buys []float64, sells []float64) float64 {
//	// implement
//	return 0
//}
//
//func AverageProfit(buys []float64, sells []float64) float64 {
//	// implement
//	return 0
//}
//
//type AverageSellPrice func(buys []float64, sells []float64) float64 // Не используется
//type AverageBuyPrice func(buys []float64, sells []float64) float64  // Не используется
//
//type GeneralCalc func(buys []float64, sells []float64) float64
//
//func getProductStatistic(gc ...GeneralCalc) []float64 {
//	var res []float64
//
//	for _, c := range gc {
//		res = append(res, c([]float64{1, 2, 3}, []float64{4, 5, 6}))
//	}
//
//	return res
//}
//
//func main() {
//	res := getProductStatistic(AverageROI, AverageProfit)
//	fmt.Println(res)
//}
