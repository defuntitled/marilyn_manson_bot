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
	if collector != c.Chat().ID {
		return c.Reply("нельзя юзать в групповых чатах по соображениям приватности")
	}
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

func getDebtor(message string) (debtor string, err error) {
	data := strings.Split(message, " ")
	if len(data) < 2 {
		err = FormatError{"message should contains debtor's nickname"}
		return
	}
	debtor = data[1]
	return
}

func getAmount(message string) (amount int, err error) {
	Log.Info(message, map[string]interface{}{})
	data := strings.Split(message, " ")
	if len(data) < 3 {
		err = FormatError{"message should contains amount"}
		return
	}
	amount, err = strconv.Atoi(data[2])
	return
}

func createDebt(c tele.Context) error {
	Log.Info(c.Text(), map[string]interface{}{})
	debtor, err := getDebtor(c.Text())
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Reply("Venom " + err.Error())
	}
	amount, err := getAmount(c.Text())
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Reply("Venom " + err.Error())
	}
	Log.Info(fmt.Sprintf("debtor %v, amount %v", debtor, amount), map[string]interface{}{})
	debt, err := Repository.GetDebtByCollectorAndDebtor(context.Background(), c.Sender().ID, debtor)
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Reply("Venom " + err.Error())
	}
	if debt == nil {
		new_debt := model.NewDebt(debtor, c.Sender().ID, amount)
		err = Repository.AddDebt(context.Background(), new_debt)
		if err != nil {
			Log.Error("failed to get debts {}", map[string]interface{}{
				"error": err,
			})
			return c.Reply("Venom " + err.Error())
		}
	} else {
		Log.Info("sosal?", map[string]interface{}{})
		debt.AddAmount(amount)
		err = Repository.UpdateDebt(context.Background(), debt)
		if err != nil {
			Log.Error("failed to update debts {}", map[string]interface{}{
				"error": err,
			})
			return c.Reply("Venom " + err.Error())
		}
	}
	return c.Reply("Записал")
}

func closeDebt(c tele.Context) error {
	debtor, err := getDebtor(c.Text())
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Reply("Venom " + err.Error())
	}
	debt, err := Repository.GetDebtByCollectorAndDebtor(context.Background(), c.Sender().ID, debtor)
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Reply("Venom " + err.Error())
	}
	if debt == nil {
		return c.Reply("Нет такого долга")
	}
	debt.SetPaid()
	err = Repository.UpdateDebt(context.Background(), debt)
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Reply("Venom " + err.Error())
	}
	return c.Reply("Записал")
}

func getDebt(c tele.Context) error {
	debtor, err := getDebtor(c.Text())
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Reply("Venom " + err.Error())
	}
	debt, err := Repository.GetDebtByCollectorAndDebtor(context.Background(), c.Sender().ID, debtor)
	if err != nil {
		Log.Error("failed to get debts {}", map[string]interface{}{
			"error": err,
		})
		return c.Reply("Venom " + err.Error())
	}
	if debt == nil {
		return c.Reply("Долгов нет")
	}
	ans := fmt.Sprintf("%v тебе должен %v %v", debt.DebtorId, debt.Amount, debt.Currency)
	return c.Reply(ans)
}
