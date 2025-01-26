package server

import (
	"context"
	"fmt"
	"marilyn_manson_bot/internal/model"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v4"
)

func helloHandler(c tele.Context) error {
	return c.Send("Venom!")
}

func listOfDebts(c tele.Context) error {
	collector := c.Message().Sender.ID
	debts, err := Repository.GetDebtsByCollector(context.Background(), collector)
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Reply("Venom " + err.Error())
	}
	ans := "Должны тебе: \n"
	for _, debt := range debts {
		ans += fmt.Sprintf("%v %v %v\n", debt.DebtorId, debt.Amount, debt.Currency)
	}
	return c.Reply(ans)
}

func parseCreateDebtMessage(message string) (*string, *int, error) {
	data := strings.Split(message, "-")
	if len(data) == 2 {
		debtor := data[1]
		amount, err := strconv.Atoi(data[0])
		if err != nil {
			return nil, nil, err
		}
		return &debtor, &amount, nil
	}
	return nil, nil, tele.Err("Cant parse request message")
}

func createDebt(c tele.Context) error {
	debtor, amount, err := parseCreateDebtMessage(c.Text())
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Reply("Venom " + err.Error())
	}
	debt, err := Repository.GetDebtByCollectorAndDebtor(context.Background(), c.Sender().ID, *debtor)
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Reply("Venom " + err.Error())
	}
	if debt == nil {
		new_debt := model.NewDebt(*debtor, c.Sender().ID, *amount)
		err = Repository.AddDebt(context.Background(), new_debt)
		if err != nil {
			Log.Error("failed to get debts {}", map[string]interface{}{
				"error": err,
			})
			return c.Reply("Venom " + err.Error())
		}
	} else {
		debt.AddAmount(*amount)
		err = Repository.UpdateDebt(context.Background(), debt)
		if err != nil {
			Log.Error("failed to get debts {}", map[string]interface{}{
				"error": err,
			})
			return c.Reply("Venom " + err.Error())
		}
	}
	return c.Reply("Success")
}
