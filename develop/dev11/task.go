package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

/*
11.	HTTP-сервер

Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.

В рамках задания необходимо:
1.	Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
2.	Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
3.	Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
4.	Реализовать middleware для логирования запросов

Методы API:
●	POST /create_event
●	POST /update_event
●	POST /delete_event
●	GET /events_for_day
●	GET /events_for_week
●	GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09). В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."} в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
1.	Реализовать все методы.
2.	Бизнес логика НЕ должна зависеть от кода HTTP сервера.
3.	В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
*/
var RWmtx sync.RWMutex
var wg sync.WaitGroup

type Info struct {
	UserId      string    `json:"user_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	taskid      int
}

func Logs(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

type Event struct {
	UserId string
	inf    []Info
}

var Events = make(map[string]*Event, 5)

func AddEvent(event Event) error {
	RWmtx.Lock()
	defer RWmtx.Unlock()
	isAlready := false
	if _, ok := Events[event.UserId]; ok {
		for _, val := range Events[event.UserId].inf {
			if val.Description == event.inf[0].Description && val.Date == event.inf[0].Date {
				isAlready = true
				break
			}
		}
		if !isAlready {
			Events[event.UserId].inf = append(Events[event.UserId].inf, event.inf[0])
			return nil
		} else {
			return errors.New("Event already exists")
		}

	}

	Events[event.UserId] = &event
	return nil
}

func UpdateInfo(event Event) error {
	RWmtx.Lock()
	defer RWmtx.Unlock()
	isUpdate := false
	if _, ok := Events[event.UserId]; ok {
		for _, v := range Events[event.UserId].inf {
			for _, ev := range event.inf {
				if v.Date == ev.Date {
					Events[event.UserId] = &event
					isUpdate = true
					return nil
				}
			}

		}
		if !isUpdate {
			Events[event.UserId].inf = append(Events[event.UserId].inf, event.inf[0])
			return nil
		}

	}
	return errors.New("Event no find")
}

func DeleteEvent(event Event) error {
	RWmtx.Lock()
	defer RWmtx.Unlock()
	if _, ok := Events[event.UserId]; ok {
		delete(Events, event.UserId)
		return nil
	}
	return errors.New("Event no find")
}

func DayEvents(id string, date time.Time) []Info {
	var result []Info
	RWmtx.RLock()
	defer RWmtx.RUnlock()
	if val, ok := Events[id]; ok {
		for _, v := range val.inf {
			if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() && v.Date.Day() == date.Day() {
				result = append(result, v)
			}
		}
	}
	return result
}
func WeekEvents(id string, date time.Time) []Info {
	RWmtx.RLock()
	defer RWmtx.RUnlock()
	var result []Info
	year, week := date.ISOWeek()

	if val, ok := Events[id]; ok {
		for _, v := range val.inf {
			year2, week2 := v.Date.ISOWeek()
			if year == year2 && week == week2 {
				result = append(result, v)
			}
		}
	}
	return result
}

func MonthEvents(id string, date time.Time) []Info {
	RWmtx.RLock()
	defer RWmtx.RUnlock()
	var result []Info

	if val, ok := Events[id]; ok {
		for _, v := range val.inf {
			if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() {
				result = append(result, v)
			}
		}
	}
	return result
}

type ResponseResult_t struct {
	Message string `json:"message"`
	Events  []Info `json:"events"`
}

func ResponseResult(w http.ResponseWriter, mess string, e []Info, status int) {
	var respResult = ResponseResult_t{
		Message: mess,
		Events:  e}

	json, err := json.Marshal(respResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func RequestToJson(r *http.Request) (*Event, error) {
	var info Info
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		return nil, errors.New("Wrong json")
	}
	event.UserId = info.UserId
	event.inf = append(event.inf, info)
	return &event, nil

}

type Server_t struct {
	httpServ *http.Server
	mux      *http.ServeMux
	port     string
}

func ResponseError(w http.ResponseWriter, mess string, status int) {
	errResp := ResponseResult_t{Message: mess}
	json, err := json.Marshal(errResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}

func (s *Server_t) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	event, err := RequestToJson(r)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = event.Validate()
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := AddEvent(*event); err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	ResponseResult(w, "Event add", event.inf, http.StatusCreated)
}

func (e *Event) Validate() error {
	id, err := strconv.Atoi(e.UserId)
	if err != nil {
		return errors.New("Wrong id")
	}
	if e.UserId == "" || id < 0 {
		return errors.New("Wrong event")
	}

	for _, v := range e.inf {
		if v.Description == "" {
			return errors.New("Wrong event")
		}
	}

	return nil
}

func (s *Server_t) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	event, err := RequestToJson(r)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = event.Validate()
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = UpdateInfo(*event)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	ResponseResult(w, "Data update", event.inf, http.StatusOK)
	return
}

func (s *Server_t) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	event, err := RequestToJson(r)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = DeleteEvent(*event)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	ResponseResult(w, "Data delete", event.inf, http.StatusOK)
	return

}

func (s *Server_t) EventForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("user_id")

	result := DayEvents(id, date)
	ResponseResult(w, "Events", result, http.StatusOK)

}

func (s *Server_t) EventForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("user_id")
	result := WeekEvents(id, date)
	ResponseResult(w, "Events", result, http.StatusOK)
}

func (s *Server_t) EventForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseError(w, "Wrong method", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("user_id")

	result := MonthEvents(id, date)
	ResponseResult(w, "Events", result, http.StatusOK)
}

func NewServer(port string) Server_t {
	return Server_t{
		mux:  http.NewServeMux(),
		port: port,
	}
}
func (s *Server_t) Run() {
	defer wg.Done()
	s.mux.HandleFunc("/create_event", Logs(http.HandlerFunc(s.CreateEvent)))
	s.mux.HandleFunc("/update_event", Logs(http.HandlerFunc(s.UpdateEvent)))
	s.mux.HandleFunc("/delete_event", Logs(http.HandlerFunc(s.DeleteEvent)))

	s.mux.HandleFunc("/events_for_day", Logs(http.HandlerFunc(s.EventForDay)))
	s.mux.HandleFunc("/events_for_week", Logs(http.HandlerFunc(s.EventForWeek)))
	s.mux.HandleFunc("/events_for_month", Logs(http.HandlerFunc(s.EventForMonth)))
	s.httpServ = &http.Server{
		Addr:    s.port,
		Handler: s.mux,
	}

	fmt.Println("Server start")
	if err := s.httpServ.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}

}

func (s *Server_t) WaitStop() {
	defer wg.Done()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	select {
	case <-quit:
		fmt.Println("Server stopped")
		s.httpServ.Shutdown(nil)
	}
}

func main() {
	serv := NewServer(":16666")
	wg.Add(2)
	go serv.WaitStop()
	serv.Run()
	wg.Wait()
}
