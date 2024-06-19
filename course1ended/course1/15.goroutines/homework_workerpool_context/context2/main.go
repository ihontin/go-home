package main

//В этих примерах показано, как использовать контексты для управления временем выполнения операций,
//передачи значений между функциями и управления несколькими горутинами.Контексты в Golang предоставляют
//мощный инструмент для эффективного управления параллельными операциями и обеспечивают гибкость в разработке приложений.

// Пример 1
// В этом примере мы используем контекст для управления временем выполнения операции. Мы устанавливаем таймаут в 1 секунду,
// и если операция не завершится в течение этого времени, она будет отменена.
//func main() {
//	ctx := context.Background()
//
//	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
//	defer cancel()
//
//	go performOperation(ctx)
//
//	time.Sleep(2 * time.Second)
//}
//
//func performOperation(ctx context.Context) {
//	select {
//	case <-ctx.Done():
//		fmt.Println("Operation cancelled")
//	case <-time.After(2 * time.Second):
//		fmt.Println("Operation completed")
//	}
//}

//Пример 2
//В этом примере мы используем контекст для передачи значения между функциями.Мы создаем контекст с помощью
//функции context.WithValue и передаем в него значение “user_id”.Затем мы извлекаем это значение в другой функции.
//func main() {
//
//	ctx := context.WithValue(context.Background(), "user_id", 123)
//	processRequest(ctx)
//}
//func processRequest(ctx context.Context) {
//	userID := ctx.Value("user_id").(int)
//	fmt.Println("Processing request for user:", userID)
//}

// Пример 3
// В этом примере мы используем композицию контекстов для управления несколькими горутинами.Мы создаем родительский
// контекст и два дочерних контекста.Затем мы передаем дочерние контексты в горутины, которые выполняют свои задачи.
//func main() {
//	ctx := context.Background()
//
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel()
//
//	ctx1, cancel1 := context.WithTimeout(ctx, 1*time.Second)
//	defer cancel1()
//
//	ctx2, cancel2 := context.WithTimeout(ctx, 2*time.Second)
//	defer cancel2()
//
//	go performTask(ctx1, "Task 1")
//	go performTask(ctx2, "Task 2")
//
//	time.Sleep(3 * time.Second)
//}
//func performTask(ctx context.Context, taskName string) {
//	select {
//	case <-ctx.Done():
//		fmt.Println(taskName, "cancelled")
//	case <-time.After(2 * time.Second):
//		fmt.Println(taskName, "completed")
//	}
//}

// Пример 4
// В данном примере создается контекст и запускается горутина, которая выполняет задачу.
// При вызове функции cancel, контекст отменяется, и горутина завершает свою работу.
//func main() {
//	ctx := context.Background() // Создание контекста
//
//	ctx, cancel := context.WithCancel(ctx) // Создание нового контекста с функцией отмены
//	defer cancel()                         // Отмена контекста при завершении функции
//
//	go performTask(ctx) // Запуск горутины с передачей контекста
//
//	time.Sleep(2 * time.Second) // Ожидание для демонстрации работы контекста
//}
//
//func performTask(ctx context.Context) {
//	for {
//		select {
//		case <-ctx.Done(): // Проверка на отмену контекста
//			fmt.Println("Task cancelled")
//			return
//		default:
//			// Выполнение задачи
//			fmt.Println("Performing task...")
//			time.Sleep(500 * time.Millisecond)
//		}
//	}
//}

//HelloHandler Пример 5
//Пример использования контекста в http handler
//В этом примере мы создаем контекст с таймаутом в 2 секунды и выполняем некоторую работу, которая занимает 3 секунды.
//После выполнения работы мы проверяем, была ли отмена контекста до завершения работы.Если контекст был отменен,
//мы отправляем клиенту ошибку http.StatusRequestTimeout.Если контекст не был отменен, мы отправляем клиенту успешный ответ.
//func HelloHandler(w http.ResponseWriter, r *http.Request) {
//	// Получаем контекст из запроса (если он уже существует) или создаем новый контекст.
//	ctx := r.Context()
//
//	// Создаем новый контекст с таймаутом в 2 секунды.
//	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
//	defer cancel() // Важно вызвать функцию `cancel` после выполнения обработчика.
//
//	// Выполняем некоторую работу, например, имитируем длительную операцию.
//	time.Sleep(3 * time.Second)
//
//	// Проверяем, была ли отмена контекста до завершения работы.
//	if ctx.Err() == context.DeadlineExceeded {
//		http.Error(w, "Request timeout", http.StatusRequestTimeout)
//		return
//	}
//
//	// Если контекст не был отменен, отправляем успешный ответ.
//	w.Write([]byte("Hello, World!"))
//}
