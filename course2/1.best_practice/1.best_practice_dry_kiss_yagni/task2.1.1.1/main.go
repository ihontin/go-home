package main

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Product struct {
	ProductID     int       // номер продукта
	Sells         []float64 // Выручка
	Buys          []float64 // Затраты
	CurrentPrice  float64   // текущая стоимость за единицу
	ProfitPercent float64   // процент прибыли
}

type Profitable interface {
	SetProduct(p *Product)
	GetAverageProfit() float64
	GetAverageProfitPercent() float64
	GetCurrentProfit() float64
	GetDifferenceProfit() float64
	GetAllData() []float64
	Average(prices []float64) float64
	Sum(prices []float64) float64
}

type StatisticProfit struct {
	product                 *Product
	getAverageProfit        func() float64
	getAverageProfitPercent func() float64
	getCurrentProfit        func() float64
	getDifferenceProfit     func() float64
	getAllData              func() []float64
}

//func WithAllData(s *StatisticProfit) {
//	s.getAllData = func() []float64 {
//		res := make([]float64, 0, 4)
//		if s.getAverageProfit != nil {
//			res = append(res, s.getAverageProfit())
//		}
//		if s.getAverageProfitPercent != nil {
//			res = append(res, s.getAverageProfitPercent())
//		}
//		if s.getCurrentProfit != nil {
//			res = append(res, s.getCurrentProfit())
//		}
//		if s.getDifferenceProfit != nil {
//			res = append(res, s.getDifferenceProfit())
//		}
//		return res
//	}
//}

func (s *StatisticProfit) GetAllData() []float64 {
	res := make([]float64, 0, 4)
	res = append(res, s.GetAverageProfit(), s.GetAverageProfitPercent(), s.GetCurrentProfit(), s.GetDifferenceProfit())
	return res
}

//func NewStatisticProfit(p *Product, gad func() []float64, g ...func() float64) *StatisticProfit {
//	return &StatisticProfit{
//		product:                 p,
//		getAverageProfit:        g[0],
//		getAverageProfitPercent: g[1],
//		getCurrentProfit:        g[2],
//		getDifferenceProfit:     g[3],
//		getAllData:              gad,
//	}
//}

func (s *StatisticProfit) SetProduct(p *Product) {
	s.product = p
}

// GetAverageProfit возвращает значение средней прибыли структуры Product
func (s *StatisticProfit) GetAverageProfit() float64 {
	averageSells := s.Average(s.product.Sells)
	averageBuys := s.Average(s.product.Buys)

	return averageSells - averageBuys
}

// GetAverageProfitPercent возвращает средний процент прибыли структуры Product
func (s *StatisticProfit) GetAverageProfitPercent() float64 {
	averageSells := s.Average(s.product.Sells)
	averageBuys := s.Average(s.product.Buys)
	if averageSells == 0 || averageBuys == 0 {
		return 0
	}
	return (averageBuys / averageSells) * 100
}

// GetCurrentProfit возвращает текущей прибыли структуры Product
func (s *StatisticProfit) GetCurrentProfit() float64 {
	currentSells := s.Sum(s.product.Sells)
	currentBuys := s.Sum(s.product.Buys)
	return currentSells - currentBuys
}

// GetDifferenceProfit - возвращает разницу между текущей и средней прибылью
func (s *StatisticProfit) GetDifferenceProfit() float64 {
	currentProfit := s.GetCurrentProfit()
	averageProfit := s.GetAverageProfit()
	return currentProfit - averageProfit
}

func (s *StatisticProfit) Average(prices []float64) float64 {
	if len(prices) == 0 {
		return 0.0
	}
	sum := s.Sum(prices)
	return sum / float64(len(prices))
}

func (s *StatisticProfit) Sum(prices []float64) float64 {
	if len(prices) < 1 {
		return 0
	}
	var sum float64
	for _, price := range prices {
		sum += price
	}
	return sum
}

//func main() {
//	statistic := &StatisticProfit{}
//
//	colla := &Product{
//		ProductID:     1,
//		Sells:         []float64{100, 200, 300},
//		Buys:          []float64{50, 75, 100},
//		CurrentPrice:  250,
//		ProfitPercent: 10,
//	}
//
//	statistic.SetProduct(colla)
//
//	// Вызов методов интерфейса Profitable на переменной StatisticProfit
//	averageProfit := statistic.GetAverageProfit()
//	averageProfitPercent := statistic.GetAverageProfitPercent()
//	currentProfit := statistic.GetCurrentProfit()
//	differenceProfit := statistic.GetDifferenceProfit()
//	allData := statistic.GetAllData()
//
//	fmt.Println("Средняя прибыль:", averageProfit)
//	fmt.Println("Средний процент прибыли:", averageProfitPercent)
//	fmt.Println("Текущая прибыль:", currentProfit)
//	fmt.Println("Разница в прибыли:", differenceProfit)
//	fmt.Println("Все данные:", allData)
//}
