package server

import (
	"context"
	"fmt"

	tele "gopkg.in/telebot.v4"
)

func helloHandler(c tele.Context) error {
	return c.Send("Venom!")
}

func listOfDebts(c tele.Context) error {
	collector := c.Message().Sender.ID
	debts, err := Repository.GetDebtsByDabtorsAndCollector(context.Background(), collector)
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Send("Venom " + err.Error())
	}
	ans := "Должны тебе: \n"
	for _, debt := range debts {
		ans += fmt.Sprintf("%v %v %v\n", debt.DebtorId, debt.Amount, debt.Currency)
	}
	return c.Reply(ans)
}
