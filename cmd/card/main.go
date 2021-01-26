package main

import (
	"fmt"
	"github.com/Geniuskaa/Task5.1_BGO-3/pkg/card"
	"github.com/Geniuskaa/Task5.1_BGO-3/pkg/transaction"
	"github.com/Geniuskaa/Task5.1_BGO-3/pkg/transfer"
	"sort"
	"time"
)

func main() {

	bank := card.NewService([]*card.Card{},"Tinkoff")
	bank.AddCard(1,"VISA", "RUB", 14_800_00, "4724 3728 3929 5030")
	bank.AddCard(2, "MASTER", "RUB", 28_750_00, "6930 2857 3892 2967")
	bank.AddCard(3, "VISA", "RUB", 352_362_00, "4626 9205 2859 2852")


	transfers := transfer.NewService(bank, 0, 0.5, 10_00, 1.5, 30_00)
	_, err := transfers.Card2Card("4724 3728 3929 5030", "6930 2857 3892 2967", 5_425, time.Now())
	if err != nil {
		switch err {
		case transfer.ErrMoneyOnCardOfSenderDontEnough:
			fmt.Println("Недостаточно средств на балансе для перевода.")
		case transfer.ErrTooLowSumOfTransfer:
			fmt.Println("Слишком маленькая сумма перевода.")
		default:
			fmt.Println("Возникла непредвиденная ошибка.")
		}
	}

	transfers.Purchase(1_204, 0, time.Date(2021,2,14,6,0,0,0, time.Local))
	transfers.Purchase(13_146, 0, time.Date(2021,2,24,6,0,0,0, time.Local))
	transfers.Purchase(106, 0, time.Date(2021,2,25,6,0,0,0, time.Local))
	transfers.Purchase(746, 0, time.Date(2021,3,14,6,0,0,0, time.Local))
	transfers.Purchase(2_546, 0, time.Date(2021,3,4,6,0,0,0, time.Local))
	transfers.Purchase(73_416, 0, time.Date(2021,4,14,6,0,0,0, time.Local))


	//for _, sample := range bank.StoreOfCards[0].Transactions {
	//	fmt.Println(sample)
	//}
	//fmt.Println(" ")

	//SortSumOfTransactions(bank.StoreOfCards[0].Transactions)

	//for _, sample := range bank.StoreOfCards[0].Transactions {
	//	fmt.Println(sample)
	//}


	//transfers.Card2Card("4724 3708 3929 5030", "6930 2857 3892 2967", 50_425)
	//transfers.Card2Card("4724 3728 3929 5030", "6930 2857 3812 2967", 725)
	//transfers.Card2Card("2424 3728 2829 5030", "97030 2857 3892 2967", 38_425)
	bank.StoreOfCards[0].SumConcurrently(1, time.Date(2021,1,1,0,0,0,0, time.Local), time.Date(2021,5,1,0, 0,0,0, time.Local))

	//fmt.Println(len(bank.StoreOfCards[0].Transactions))


}



func SortSumOfTransactions(transactions []*transaction.Transaction) []*transaction.Transaction {
	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].Amount > transactions[j].Amount
	})
	return transactions
}
