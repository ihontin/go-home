package main

import "fmt"

//Пример 1
//В этом примере мы имеем два несовместимых интерфейса: Square и Circle. Мы создаем адаптер ShapeAdapter,
//который оборачивает объект Circle и предоставляет методы, совместимые с интерфейсом Square.

//type Square interface {
//	DrawSquare()
//}
//
//type Circle interface {
//	DrawCircle()
//}
//
//type CircleShape struct{}
//
//func (c *CircleShape) DrawCircle() {
//	fmt.Println("Drawing a circle")
//}
//
//type ShapeAdapter struct {
//	Circle Circle
//}
//
//func (a *ShapeAdapter) DrawSquare() {
//	a.Circle.DrawCircle()
//}
//
//func main() {
//	circle := &CircleShape{}
//	adapter := &ShapeAdapter{
//		Circle: circle,
//	}
//	adapter.DrawSquare()
//}

//Пример 2
//В этом примере у нас есть два различных интерфейса: Printer и Scanner. Мы создаем адаптер PrinterAdapter,
//который оборачивает объект Scanner и предоставляет методы, совместимые с интерфейсом Printer.

//type Printer interface {
//	Print()
//}
//
//type Scanner interface {
//	Scan()
//}
//
//type ScannerDevice struct{}
//
//func (s *ScannerDevice) Scan() {
//	fmt.Println("Scanning a document")
//}
//
//type PrinterAdapter struct {
//	Scanner Scanner
//}
//
//func (a *PrinterAdapter) Print() {
//	a.Scanner.Scan()
//}
//
//func main() {
//	scanner := &ScannerDevice{}
//	adapter := &PrinterAdapter{
//		Scanner: scanner,
//	}
//	adapter.Print()
//}

//Пример 3
//В этом примере у нас есть интерфейс Database и структура MySQLDatabase, которая не совместима с интерфейсом.
//Мы создаем адаптер PostgreSQLAdapter, который оборачивает объект MySQLDatabase и предоставляет методы,
//совместимые с интерфейсом Database.

//type Database interface {
//	Connect()
//	Query(query string)
//}
//
//type MySQLDatabase struct{}
//
//func (m *MySQLDatabase) Connect(connection string) {
//	fmt.Println("Connecting to MySQL database", connection)
//}
//
//func (m *MySQLDatabase) Query(query string) {
//	fmt.Println("Executing MySQL query:", query)
//}
//
//type PostgreSQLAdapter struct {
//	MySQLDB *MySQLDatabase
//}
//
//func (a *PostgreSQLAdapter) Connect() {
//	a.MySQLDB.Connect("some:connection:string")
//}
//
//func (a *PostgreSQLAdapter) Query(query string) {
//	a.MySQLDB.Query(query)
//}
//
//func main() {
//	mysqlDB := &MySQLDatabase{}
//	adapter := &PostgreSQLAdapter{
//		MySQLDB: mysqlDB,
//	}
//	adapter.Connect()
//	adapter.Query("SELECT * FROM users")
//}

//Пример 4
//В этом примере у нас есть интерфейс Logger и структура FileLogger, которая не совместима с интерфейсом.
//Мы создаем адаптер DatabaseLoggerAdapter, который оборачивает объект FileLogger и предоставляет методы,
//совместимые с интерфейсом Logger.

//type Logger interface {
//	Log(message string)
//}
//
//type FileLogger struct{}
//
//func (f *FileLogger) Log(message string) error {
//	fmt.Println("Logging message to file:", message)
//
//	return nil
//}
//
//type DatabaseLoggerAdapter struct {
//	FileLogger *FileLogger
//}
//
//func (a *DatabaseLoggerAdapter) Log(message string) {
//	a.FileLogger.Log(message)
//}
//
//func main() {
//	fileLogger := &FileLogger{}
//	adapter := &DatabaseLoggerAdapter{
//		FileLogger: fileLogger,
//	}
//	adapter.Log("Error occurred")
//}

//Пример 5
//В этом примере у нас есть интерфейс PaymentProcessor и структура BankPaymentProcessor,
//которая не совместима с интерфейсом. Мы создаем адаптер PaymentGatewayAdapter, который оборачивает
//объект BankPaymentProcessor и предоставляет методы, совместимые с интерфейсом PaymentProcessor.

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type BankPaymentProcessor struct{}

func (b *BankPaymentProcessor) ProcessPayment(amount int) {
	fmt.Println("Processing payment via bank:", amount)
}

type PaymentGatewayAdapter struct {
	BankProcessor BankPaymentProcessor
}

func (a *PaymentGatewayAdapter) ProcessPayment(amount float64) {
	a.BankProcessor.ProcessPayment(int(amount))
}

func main() {
	bankProcessor := &BankPaymentProcessor{}
	adapter := &PaymentGatewayAdapter{
		BankProcessor: *bankProcessor,
	}
	adapter.ProcessPayment(100.50)
}
