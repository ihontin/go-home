package main

import (
	"encoding/json"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
	"github.com/guptarohit/asciigraph"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

// Val значение (актива) валюты в главно меню - пустая строка, в остальных: 1.BTC, 2. LTC, 3. ETH
var Val string

// Cache - структура для конкурентного доступа к данным
type Cache struct {
	mutex sync.RWMutex
	data  map[string][]float64
}

// NewCash Конструктор для создания экземпляра Cache
func NewCash() *Cache {
	return &Cache{data: make(map[string][]float64)}
}

// SetCash - метод добавляющий данные в Cache
func (c *Cache) SetCash(key string, v float64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = append(c.data[key], v)
}

// GetCash - метод возвращает данные из Cache по ключу
func (c *Cache) GetCash(key string) []float64 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.data[key]
}

// Delete - метод удаляет данные из Cache по ключу
func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)
}
func main() {
	// обновление вывода в реальном времени
	writer := uilive.New()
	writer.Start()
	defer writer.Stop() // Закрыть writer в конце работы
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	var recreate = make(chan bool) // канал для обновления консоли при вводе пользователя
	var graphC = NewCash()         // создание новой структуры
	// запускает функцию обновления выводимого графика и времени в отдельной горутине.
	go updateLiveTime(recreate, writer, graphC)
	//цикл отслеживающий ввод пользователя
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if char == 49 { // если ввод == 1
			graphC.Delete(Val) // удаляем данные в структуре графика
			Val = "BTC"        // меняем значение Валюты
		} else if char == 50 { // если ввод == 2
			graphC.Delete(Val) // удаляем данные в структуре графика
			Val = "LTC"        // меняем значение Валюты
		} else if char == 51 { // если ввод == 3
			graphC.Delete(Val) // удаляем данные в структуре графика
			Val = "ETH"        // меняем значение Валюты
		} else if char == 113 { // если ввод == q
			break
		}
		// если ввод == Backspace
		if key == keyboard.KeyBackspace2 || key == keyboard.KeyBackspace {
			graphC.Delete(Val) // удаляем данные в структуре графика
			Val = ""           // меняем значение Валюты
		}
		// если нажата одна из клавиш пишем в канал
		recreate <- true
	}
}

// worker получает данные о ценах на актив ежесекундно
func worker(gCache *Cache, pointG chan<- float64) {

	url := fmt.Sprintf("https://cex.io/api/ticker/%s/USD", Val) // форматирует url
	client := &http.Client{}                                    // использует структуру Client для подготовки запроса

	req, err := http.NewRequest("GET", url, nil) // создает новый http запрос
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}
	resp, err := client.Do(req) // отправляет запрос на сервер, возвращает ответ
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()                // закрывает ответ после выхода из worker
	body, err := ioutil.ReadAll(resp.Body) // читает все данные из тела ответа в []byte
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}
	var result = make(map[string]interface{}, 12) // создаем карту сокровищ из дюжины позиций
	err = json.Unmarshal(body, &result)           // декодирует из []byte json в map
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}
	// достает нужную единицу данных и переводит её from str to float64
	resFloat, err1 := strconv.ParseFloat(result["last"].(string), 64)
	if err1 != nil {
		fmt.Println("strconv.ParseFloat error:", err)
		return
	}
	gCache.SetCash(Val, resFloat) // добавляет полученную единицу в срез с которым работает график
	pointG <- resFloat            // отправляет полученную единицу в канал для её вывода в меню валюты на консоли
}

// updateLiveTime выводит в консоль и ежесекундно обновляет график данных актива, время и дату
func updateLiveTime(ch <-chan bool, w *uilive.Writer, gCache *Cache) {
	t := time.Tick(time.Second * 1)
	var pointG = make(chan float64)                                                                      // канал цены на актив
	var mainMenu = "1. BTC_USD\n2. LTC_USD\n3. ETH_USD\n\nPress 1-3 to change symbol, press q to exit\n" // основное меню
	//var menuNow string // меню сейчас
	for {
		// форматирует текущее время и дату
		currentTime := time.Now().Format("15:04:05")
		currentYear := time.Now().Format("2006-01-02")
		newGraph := ""      // обновляем граф
		menuNow := mainMenu // выводимому меню присваивает основное
		if Val != "" {      // если актив (валюта) выбран
			go worker(gCache, pointG)                             // достает данные о ценах на актив
			Point := <-pointG                                     // присваивает последнюю единицу данных на актив
			var subMenu = fmt.Sprintf("%s_USD: %f\n", Val, Point) // форматирует вывод меню в консоль
			menuNow = subMenu                                     // присваивает меню новый формат
			//меняет график на основе последних данных на актив
			newGraph = asciigraph.Plot(gCache.GetCash(Val), asciigraph.Width(100), asciigraph.Height(10))
		}

		// Обновит выводимое меню время год и график
		fmt.Fprintf(w, "%s\n\033[31m%v\033[0m\nТекущая дата: %v\nТекущее время: %s\n",
			menuNow, newGraph, currentYear, currentTime)

		// обновление на вывод
		w.Flush()

		// ежесекундный тикер или очистка консоли при выборе меню пользователем
		select {
		case <-t:
		case <-ch:
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}
}
