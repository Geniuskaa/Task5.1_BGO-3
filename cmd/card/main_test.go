package main

import (
	"github.com/Geniuskaa/Task5.1_BGO-3/pkg/card"
	"github.com/Geniuskaa/Task5.1_BGO-3/pkg/transaction"
	"github.com/Geniuskaa/Task5.1_BGO-3/pkg/transfer"
	"reflect"
	"testing"
)

func TestSortSumOfTransactions(t *testing.T) {
	b := card.NewService([]*card.Card{}, "Cobolt")
	b.AddCard(1,"VISA", "RUB", 14_800_00, "4724 3728 3929 5030")
	transfers := transfer.NewService(b, 0, 0.5, 10_00, 1.5, 30_00)
	transfers.Purchase(1_204, 0, 0)
	transfers.Purchase(13_146, 0, 0)
	transfers.Purchase(106, 0, 0)
	transfers.Purchase(746, 0, 0)

	type args struct {
		bank *card.Service
	}

	tests := []struct {
		name string
		args args
		want []*transaction.Transaction
	}{
		{"First", args{b}, []*transaction.Transaction{{
			Id:     20,
			Amount: 1314600,
			MCC:    "5090",
			Date: 	0,
			Status: "Completed" },

		{	Id:     20,
			Amount: 120400,
			MCC:    "5090",
			Date: 	0,
			Status: "Completed" },

		{	Id:     20,
			Amount: 74600,
			MCC:    "5090",
			Date: 	0,
			Status: "Completed" },

		{	Id:     20,
			Amount: 10600,
			MCC:    "5090",
			Date: 	0,
			Status: "Completed" },
		}},
	}

	for _, tt := range tests {
		if got := SortSumOfTransactions(tt.args.bank.StoreOfCards[0].Transactions); !reflect.DeepEqual(tt.args.bank.StoreOfCards[0].Transactions, tt.want) {
			t.Errorf("Sum() = %v, want %v", got, tt.want)
		}
	}
}