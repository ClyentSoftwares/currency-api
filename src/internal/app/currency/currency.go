package currency

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/clyentsoftwares/currency-api/src/internal/pkg/cache"
)

type Service struct {
	cache  *cache.Cache
	apiKey string
}

func NewService() *Service {
	apiKey := os.Getenv("OPEN_EXCHANGE_RATES_APP_ID")
	if apiKey == "" {
		panic("OPEN_EXCHANGE_RATES_APP_ID is not set")
	}
	c := cache.NewCache()
	s := &Service{
		cache:  c,
		apiKey: apiKey,
	}
	s.refreshRates()
	return s
}

func (s *Service) refreshRates() {
	go func() {
		for {
			s.updateRates()
			time.Sleep(1 * time.Hour)
		}
	}()
}

func (s *Service) updateRates() {
	resp, err := http.Get("https://openexchangerates.org/api/latest.json?app_id=" + s.apiKey)
	if err != nil {
		log.Println("Failed to fetch exchange rates:", err)
		return
	}
	defer resp.Body.Close()

	var data struct {
		Rates map[string]float64 `json:"rates"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("Failed to decode exchange rates:", err)
		return
	}

	s.cache.Set("rates", data.Rates, 2*time.Hour)
}

func round(num float64) float64 {
	return math.Round(num*100) / 100
}

func (s *Service) Convert(amount float64, from, to string) (float64, error) {
	rates, err := s.getRates()
	if err != nil {
		return 0, err
	}

	rateFrom, ok := rates[from]
	if !ok {
		return 0, fmt.Errorf("currency %s not supported", from)
	}
	rateTo, ok := rates[to]
	if !ok {
		return 0, fmt.Errorf("currency %s not supported", to)
	}

	amountInUSD := amount / rateFrom
	convertedAmount := round(amountInUSD * rateTo)

	return convertedAmount, nil
}

func (s *Service) GetRate(from, to string) (float64, error) {
	rates, err := s.getRates()
	if err != nil {
		return 0, err
	}

	rateFrom, ok := rates[from]
	if !ok {
		return 0, fmt.Errorf("currency %s not supported", from)
	}
	rateTo, ok := rates[to]
	if !ok {
		return 0, fmt.Errorf("currency %s not supported", to)
	}

	return round(rateTo / rateFrom), nil
}

func (s *Service) GetAllRates(base string) (map[string]float64, error) {
	rates, err := s.getRates()
	if err != nil {
		return nil, err
	}

	rateBase, ok := rates[base]
	if !ok {
		return nil, fmt.Errorf("currency %s not supported", base)
	}

	ratesConverted := make(map[string]float64)
	for currency, rate := range rates {
		ratesConverted[currency] = round(rate / rateBase)
	}

	return ratesConverted, nil
}

func (s *Service) getRates() (map[string]float64, error) {
	rates, exists := s.cache.Get("rates")
	if !exists {
		s.updateRates() // Update rates if not in cache
		rates, exists = s.cache.Get("rates")
		if !exists {
			return nil, fmt.Errorf("failed to get exchange rates")
		}
	}

	ratesMap, ok := rates.(map[string]float64)
	if !ok {
		return nil, fmt.Errorf("invalid rate format in cache")
	}

	return ratesMap, nil
}
