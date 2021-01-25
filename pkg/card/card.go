package card

import (
	"errors"
	"fmt"
	"github.com/Geniuskaa/Task5.1_BGO-3/pkg/transaction"
	"math/rand"
	"strings"
	"sync"
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

func (t *Card) SumConcurrently(goroutines int) int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	transactionsByMonths := make([][]*transaction.Transaction, 12, 12)


	stop := 0
	for i := range transactionsByMonths {

			switch i {

			case 0:
				for _, sample := range t.Transactions {
					if sample.Date >= 1612137600 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break

			case 1:
				for _, sample := range t.Transactions[stop :] {
					if sample.Date >= 1614556800 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break

			case 2:
				for _, sample := range t.Transactions[stop :] {
					if sample.Date >= 1617235200 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break

			case 3:
				for _, sample := range t.Transactions[stop :] {
					if sample.Date >= 1619827200 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break

			case 4:
				for _, sample := range t.Transactions[stop :] {
					if sample.Date >= 1622505600 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break

			case 5:
				for _, sample := range t.Transactions[stop :] {
					if sample.Date >= 1625097600 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break

			case 6:
				for _, sample := range t.Transactions[stop :] {
					if sample.Date >= 1627776000 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break

			case 7:
				for _, sample := range t.Transactions[stop :] {
					if sample.Date >= 1630454400 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break

			case 8:
				for _, sample := range t.Transactions[stop :] {
					if sample.Date >= 1633046400 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break

			case 9:
				for _, sample := range t.Transactions[stop :] {
					if sample.Date >= 1635724800 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break

			case 10:
				for _, sample := range t.Transactions[stop :] {
					if sample.Date >= 1638316800 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break

			case 11:
				for _, sample := range t.Transactions[stop :] {
					if sample.Date >= 1640995200 {
						break
					}
					stop++
					transactionsByMonths[i] = append(transactionsByMonths[i], sample)
				}
				break
			}

	}

	fmt.Println("длина января ", len(transactionsByMonths[0]))

	total := int64(0)
	partSize := len(transactionsByMonths) / goroutines // Динамически поменяем в For
	for i := 0; i < goroutines; i++ {

		part := transactionsByMonths[i*partSize : (i+1)*partSize]
		go func() {
			for _, element := range part {
				sum := sum(element)
				fmt.Println(sum / 100)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return total
}


