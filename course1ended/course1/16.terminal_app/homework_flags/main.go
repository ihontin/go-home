package main

//Пример 1
//В этом примере показано, как определить флаг командной строки с целочисленным значением.
//go run main.go -num=50

//func main() {
//	var num int
//	flag.IntVar(&num, "num", 0, "Целочисленное значение")
//	flag.Parse()
//
//	fmt.Println("Значение флага num:", num)
//}

//Пример 2
//В этом примере показано, как определить флаг командной строки с булевым значением.
//go run main.go -flag false

//func main() {
//	var flagValue bool
//	flag.BoolVar(&flagValue, "flag", false, "Булево значение")
//	flag.Parse()
//
//	fmt.Println("Значение флага flag:", flagValue)
//}

//Пример 3
//В этом примере показано, как определить флаг командной строки с текстовым значением.
//go run main.go -text='В добрый путь!'

//func main() {
//	var text string
//	flag.StringVar(&text, "text", "", "Текстовое значение")
//	flag.Parse()
//
//	fmt.Println("Значение флага text:", text)
//}

//Пример 4
//В этом примере показано, как определить флаг командной строки с кратким и полным именем.
//go run main.go -name "I am Antoshka"
//go run main.go -n "I am Antoshka"
//func main() {
//	var name string
//	flag.StringVar(&name, "n", "", "Имя")
//	flag.StringVar(&name, "name", "", "Имя")
//	flag.Parse()
//
//	fmt.Println("Значение флага name:", name)
//}

//Пример 5
//В этом примере показано, как определить флаг командной строки с аргументом.

//func main() {
//	var argValue string
//	flag.StringVar(&argValue, "arg", "", "Аргумент")
//	flag.Parse()
//
//	fmt.Println("Значение аргумента arg:", argValue)
//}

//Пример 6
//В этом примере показано, как использовать значение по умолчанию для флага командной строки.

//func main() {
//	var name string
//	flag.StringVar(&name, "name", "John", "Имя")
//	flag.Parse()
//
//	fmt.Println("Значение флага name:", name)
//}

//Пример 7
//В этом примере показано, как обрабатывать ошибки, связанные с флагами командной строки.

//func main() {
//	var name string
//	flag.StringVar(&name, "name", "", "Имя")
//	flag.Parse()
//
//	if name == "" {
//		fmt.Println("Ошибка: не указано имя")
//		flag.Usage()
//		os.Exit(1)
//	}
//
//	fmt.Println("Привет,", name)
//}

//Пример 8
//В этом примере показано, как использовать флаги командной строки для настройки поведения приложения.

//func main() {
//	var verbose bool
//	flag.BoolVar(&verbose, "verbose", false, "Вывод подробной информации")
//	flag.Parse()
//
//	if verbose {
//		fmt.Println("Подробный вывод включен")
//	} else {
//		fmt.Println("Подробный вывод выключен")
//	}
//}

//Пример 9 Использование флагов с использованием пакета Cobra
//В этом примере мы покажем, как использовать флаги с помощью пакета Cobra,
//который предоставляет более мощные возможности для работы с командной строкой.

// В этом примере мы создаем корневую команду с помощью cobra.Command. Затем мы определяем флаг “name” с помощью
// rootCmd.Flags().String(). Значение флага по умолчанию установлено как “World”. В функции Run мы используем значение
// флага для вывода приветствия.
//func main() {
//	// Создание корневой команды
//	rootCmd := &cobra.Command{
//		Use:   "app",
//		Short: "Пример приложения с использованием Cobra",
//		Run: func(cmd *cobra.Command, args []string) {
//			// Использование значения флага
//			name, _ := cmd.Flags().GetString("name")
//			fmt.Printf("Привет, %s!\n", name)
//		},
//	}
//
//	// Определение флага
//	rootCmd.Flags().String("name", "World", "Имя для приветствия")
//
//	// Запуск команды
//	rootCmd.Execute()
//}

//type Color string
//
//const (
//	ColorBlack  Color = "\u001b[30m"
//	ColorRed          = "\u001b[31m"
//	ColorGreen        = "\u001b[32m"
//	ColorYellow       = "\u001b[33m"
//	ColorBlue         = "\u001b[34m"
//	ColorReset        = "\u001b[0m"
//)
//
//func colorize(color Color, message string) {
//	fmt.Println(string(color), message, string(ColorReset))
//}
//
//func main() {
//	useColor := flag.Bool("color", false, "display colorized output")
//	flag.Parse()
//
//	if *useColor {
//		colorize(ColorBlue, "Hello, DigitalOcean!")
//		return
//	}
//	fmt.Println("Hello, DigitalOcean!")
//}
