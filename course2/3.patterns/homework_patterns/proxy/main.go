package main

import "fmt"

//Пример 1. Прокси для доступа к удаленному сервису
//В этом примере мы создадим прокси для доступа к удаленному сервису, который предоставляет информацию о погоде.
//Прокси будет выполнять запрос к удаленному сервису только при первом вызове метода получения погоды,
//а затем — кэшировать результаты для последующих вызовов.

//type WeatherService interface {
//	GetWeather(city string) string
//}
//
//type RealWeatherService struct{}
//
//func (rws *RealWeatherService) GetWeather(city string) string {
//	// Здесь был бы код для запроса погоды у удаленного сервиса
//	return "Солнечно"
//}
//
//type WeatherServiceProxy struct {
//	realService RealWeatherService
//	cache       map[string]string
//}
//
//func (wsp *WeatherServiceProxy) GetWeather(city string) string {
//	if weather, ok := wsp.cache[city]; ok {
//		fmt.Println("Используется кэш")
//		return weather
//	}
//
//	weather := wsp.realService.GetWeather(city)
//	wsp.cache[city] = weather
//	fmt.Println("Запрос к удаленному сервису")
//	return weather
//}
//
//func main() {
//	proxy := WeatherServiceProxy{
//		realService: RealWeatherService{},
//		cache:       make(map[string]string),
//	}
//
//	fmt.Println(proxy.GetWeather("Москва"))
//	fmt.Println(proxy.GetWeather("Москва"))
//}

//Пример 2. Прокси для ограничения доступа
//В этом примере мы создадим прокси для ограничения доступа к определенным методам объекта.
//Прокси будет проверять, имеет ли пользователь достаточные права для вызова метода,
//и разрешать или запрещать доступ в зависимости от результатов проверки.

//type UserService interface {
//	GetUserInfo(userID int) string
//}
//
//type RealUserService struct{}
//
//func (rus *RealUserService) GetUserInfo(userID int) string {
//	// Здесь был бы код для получения информации о пользователе
//	return "Информация о пользователе"
//}
//
//type UserServiceProxy struct {
//	realService RealUserService
//	authorized  bool
//}
//
//func (usp *UserServiceProxy) GetUserInfo(userID int) string {
//	if !usp.authorized {
//		return "Доступ запрещен"
//	}
//
//	return usp.realService.GetUserInfo(userID)
//}
//
//func main() {
//	proxy := UserServiceProxy{
//		realService: RealUserService{},
//		authorized:  true,
//	}
//
//	fmt.Println(proxy.GetUserInfo(123))
//	proxy.authorized = false
//	fmt.Println(proxy.GetUserInfo(123))
//}

//Пример 3. Прокси для ленивой инициализации объекта
//В этом примере мы создадим прокси, который будет откладывать создание объекта до момента,
//когда он действительно понадобится.

//type ExpensiveObject interface {
//	DoSomething()
//}
//
//type RealExpensiveObject struct{}
//
//func (reo *RealExpensiveObject) DoSomething() {
//	fmt.Println("Выполнение дорогостоящей операции")
//}
//
//type ExpensiveObjectProxy struct {
//	realObject *RealExpensiveObject
//}
//
//func (eop *ExpensiveObjectProxy) DoSomething() {
//	if eop.realObject == nil {
//		fmt.Println("Создание дорогостоящего объекта")
//		eop.realObject = &RealExpensiveObject{}
//	}
//
//	eop.realObject.DoSomething()
//}
//
//func main() {
//	proxy := ExpensiveObjectProxy{}
//
//	proxy.DoSomething()
//	proxy.DoSomething()
//}

//Пример 4. Прокси для логирования действий
//В этом примере мы создадим прокси для логирования действий при вызове методов объекта.

type Logger interface {
	Log(message string)
}

type RealLogger struct{}

func (rl *RealLogger) Log(message string) {
	fmt.Println("Лог:", message)
}

type LoggerProxy struct {
	realLogger RealLogger
}

func (lp *LoggerProxy) Log(message string) {
	fmt.Println("Начало логирования")
	lp.realLogger.Log(message)
	fmt.Println("Конец логирования")
}

func main() {
	proxy := LoggerProxy{
		realLogger: RealLogger{},
	}

	proxy.Log("Выполнение операции")
}

//Пример 5. Прокси для счетчика вызовов методов
//В этом примере мы создадим прокси для подсчета количества вызовов методов объекта.

//type Counter interface {
//	Increment()
//	GetCount() int
//}
//
//type RealCounter struct {
//	count int
//}
//
//func (rc *RealCounter) Increment() {
//	rc.count++
//}
//
//func (rc *RealCounter) GetCount() int {
//	return rc.count
//}
//
//type CounterProxy struct {
//	realCounter RealCounter
//}
//
//func (cp *CounterProxy) Increment() {
//	fmt.Println("Увеличение счетчика")
//	cp.realCounter.Increment()
//}
//
//func (cp *CounterProxy) GetCount() int {
//	return cp.realCounter.GetCount()
//}
//
//func main() {
//	proxy := CounterProxy{
//		realCounter: RealCounter{},
//	}
//
//	proxy.Increment()
//	proxy.Increment()
//	fmt.Println(proxy.GetCount())
//}

//Пример 6. Прокси для проверки прав доступа
//В этом примере мы создадим прокси для проверки прав доступа к определенным методам объекта.

//type AccessControl interface {
//	HasAccess(userID int) bool
//}
//
//type RealAccessControl struct{}
//
//func (rac *RealAccessControl) HasAccess(userID int) bool {
//	// Здесь был бы код для проверки прав доступа пользователя
//	return userID == 1
//}
//
//type AccessControlProxy struct {
//	realAccessControl RealAccessControl
//}
//
//func (acp *AccessControlProxy) HasAccess(userID int) bool {
//	fmt.Println("Проверка прав доступа")
//	return acp.realAccessControl.HasAccess(userID)
//}
//
//func main() {
//	proxy := AccessControlProxy{
//		realAccessControl: RealAccessControl{},
//	}
//
//	fmt.Println(proxy.HasAccess(1))
//	fmt.Println(proxy.HasAccess(2))
//}

//Пример 7. Прокси для отложенной загрузки изображений
//В этом примере мы создадим прокси, который будет загружать изображение только при первом вызове метода отображения,
//а затем — кэшировать его для последующих вызовов.

//type Image interface {
//	Display()
//}
//
//type RealImage struct {
//	filename string
//}
//
//func (ri *RealImage) LoadImage() {
//	fmt.Println("Загрузка изображения", ri.filename)
//}
//
//func (ri *RealImage) Display() {
//	fmt.Println("Отображение изображения", ri.filename)
//}
//
//type ImageProxy struct {
//	realImage *RealImage
//	filename  string
//}
//
//func (ip *ImageProxy) Display() {
//	if ip.realImage == nil {
//		ip.realImage = &RealImage{filename: ip.filename}
//		ip.realImage.LoadImage()
//	}
//
//	ip.realImage.Display()
//}
//
//func main() {
//	proxy := ImageProxy{
//		filename: "image.jpg",
//	}
//
//	proxy.Display()
//	proxy.Display()
//}

//Пример 8. Прокси для ограничения количества вызовов метода
//В этом примере мы создадим прокси, который будет разрешать вызов метода только определенное количество раз.

//type LimitedAccessObject interface {
//	DoSomething()
//}
//
//type RealLimitedAccessObject struct{}
//
//func (rlo *RealLimitedAccessObject) DoSomething() {
//	fmt.Println("Выполнение операции")
//}
//
//type LimitedAccessObjectProxy struct {
//	realObject RealLimitedAccessObject
//	remaining  int
//}
//
//func (lap *LimitedAccessObjectProxy) DoSomething() {
//	if lap.remaining > 0 {
//		lap.realObject.DoSomething()
//		lap.remaining--
//	} else {
//		fmt.Println("Доступ запрещен")
//	}
//}
//
//func main() {
//	proxy := LimitedAccessObjectProxy{
//		realObject: RealLimitedAccessObject{},
//		remaining:  2,
//	}
//
//	proxy.DoSomething()
//	proxy.DoSomething()
//	proxy.DoSomething()
//}
