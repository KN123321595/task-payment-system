package main

import (
	"fmt"
	"task-payment-system/internal/payment_system"
)

func main() {
	ps := payment_system.NewPaymentSystem()

	fmt.Println("1. Вывод IBAN счета эмиссии: ", ps.GetEmissionAccountIban())
	fmt.Println("2. Вывод IBAN счета уничтожения: ", ps.GetDestructionAccountIban())
	fmt.Println("------------------------------------------------------------------------")

	fmt.Println("3. Осуществление эмиссии. Сумма 555 рублей 11 копеек")
	ps.EmitMoney(55511)
	fmt.Println(ps.ListAccounts())
	fmt.Println("------------------------------------------------------------------------")

	fmt.Println("4. Отправка определенной суммы денег с указанного счета на счет уничтожения.")
	fmt.Println("Успешный перевод. Отправка 100 рублей 50 копеек с эмиссионого счета на счет уничтожения")
	if err := ps.DestroyMoney(ps.GetEmissionAccountIban(), 10050); err != nil {
		fmt.Println("error transfer: ", err)
	}

	fmt.Println()
	fmt.Println("Неудачный перевод (счет с которого нужно отправить не найден)")
	if err := ps.DestroyMoney("", 10050); err != nil {
		fmt.Println("error transfer: ", err)
	}

	fmt.Println()
	fmt.Println(ps.ListAccounts())
	fmt.Println("------------------------------------------------------------------------")

	fmt.Println("5. Открытие нового счета. Название - test1")
	accountIban1 := ps.CreateBaseAccountAndGetIban("test1")
	fmt.Println("IBAN нового счета: ", accountIban1)
	fmt.Println(ps.ListAccounts())
	fmt.Println("------------------------------------------------------------------------")

	fmt.Println("6. Перевод заданной суммы денег между двумя указанными счетами")
	fmt.Println("Открытие нового счета с названием - test2")
	accountIban2 := ps.CreateBaseAccountAndGetIban("test2")
	fmt.Println("IBAN нового счета: ", accountIban2)

	fmt.Println()
	fmt.Println("Успешный перевод. Перевод 50 рублей 14 копеек с эмиссионого счета на счет test1")
	if err := ps.TransferMoney(ps.GetEmissionAccountIban(), accountIban1, 5014); err != nil {
		fmt.Println("error transfer: ", err)
	}
	fmt.Println(ps.ListAccounts())

	fmt.Println()
	fmt.Println("Неудачный перевод (счет не найден)")
	if err := ps.TransferMoney(accountIban1, "", 10000); err != nil {
		fmt.Println("error transfer: ", err)
	}

	fmt.Println()
	fmt.Println("Неудачный перевод (недостаточно средства). Перевод 60 рублей с test1 счета на счет test2")
	if err := ps.TransferMoney(accountIban1, accountIban2, 6000); err != nil {
		fmt.Println("error transfer: ", err)
	}

	fmt.Println()
	fmt.Println("Успешный перевод. Перевод 5.75 со счета test1 на счет test2 через параметр JSON")
	if err := ps.TransferMoneyJSON(fmt.Sprintf("{\"sender_iban\": \"%s\",\"receiver_iban\": \"%s\",\"money\": \"5.75\"}", accountIban1, accountIban2)); err != nil {
		fmt.Println("error transfer JSON: ", err)
	}
	fmt.Println(ps.ListAccounts())

	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("7. Вывод списка всех счетов в JSON формате")
	fmt.Println(ps.ListAccounts())
}
