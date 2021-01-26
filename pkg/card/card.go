package card

import (
	"errors"
	"fmt"
	"github.com/Geniuskaa/Task5.1_BGO-3/pkg/transaction"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Card struct {
	Id int64
	Issuer string
	Currency string
	Balance int64
	Number string
	Transactions []*transaction.Transaction
}

type Service struct {
	bank string
	StoreOfCards []*Card
}

type part struct {
	monthTimestamp int64
	transactions []*transaction.Transaction
}

func NewService(storeOfCards []*Card, bankName string) *Service {
	return &Service{
		bank: bankName,
		StoreOfCards: storeOfCards}
}

func (s *Service) AddCard(id int64, issuer string, currency string, balance int64, number string) {
	s.StoreOfCards = append(s.StoreOfCards, &Card{
		Id:       id,
		Issuer:   issuer,
		Currency: currency,
		Balance:  balance,
		Number:   number,
	})
}

var ErrCardNotInOurBase = errors.New("Данной карты нет в нашей базе данных.")


func (s *Service) SearchCards(number string) (err error, index int) {
	for i, _ := range s.StoreOfCards {
		if s.StoreOfCards[i].Number == number {
			return nil, i
		}
	}
	if strings.HasPrefix(number, "5106 21") {
		s.StoreOfCards = append(s.StoreOfCards, &Card{
			Id:           rand.Int63n(1000),
			Issuer:       "VISA",
			Currency:     "RUB",
			Balance:      rand.Int63n(10000000),
			Number:       number,
			Transactions: nil,
		})
		return nil, len(s.StoreOfCards) - 1
	}
	return ErrCardNotInOurBase, -1
}

func sum(transactions []*transaction.Transaction) int64 {
	result := int64(0)
	for _, transaction := range transactions {
		result += transaction.Amount
	}
	return result
}

func (t *Card) SumConcurrently(goroutines int, from time.Time, to time.Time) int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	months := make([]*part, 0)

	next := from
	for next.Before(to) {
		months = append(months, &part{
			monthTimestamp: next.Unix(),
		})
		next = next.AddDate(0, 1, 0)
	}
	months = append(months, &part{
		monthTimestamp: to.Unix(),
	})

	for j, transaction := range t.Transactions {
		if months[0].monthTimestamp <= transaction.Date.Unix() && transaction.Date.Unix() < months[len(months) - 1].monthTimestamp {
			for i := 1; i < len(months); i++ {
				if t.Transactions[j].Date.Unix() < months[i].monthTimestamp {
					months[i - 1].transactions = append(months[i - 1].transactions, t.Transactions[j])
					break
				}
			}
		}
	}

	total := int64(0)
	partSize := len(months) / goroutines // Динамически поменяем в For
	for i := 0; i < goroutines; i++ {

		part := months[i*partSize : (i+1)*partSize]
		go func() {
			for _, element := range part {
				sum := sum(element.transactions)
				total += sum
				fmt.Println(sum / 100)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("За выбранный промежуток времени было потрачено ", total / 100, " рублей")
	return total
}


