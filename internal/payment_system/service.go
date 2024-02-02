package payment_system

import (
	"encoding/json"
	"fmt"
	"task-payment-system/pkg/iban"
)

type PaymentSystem struct {
	accounts           map[string]*Account
	emissionAccount    *Account
	destructionAccount *Account
}

func NewPaymentSystem() *PaymentSystem {
	ps := &PaymentSystem{
		accounts: make(map[string]*Account),
	}

	ps.emissionAccount = ps.createAccount(EMISSION_NAME, SPECIAL_STATUS)
	ps.destructionAccount = ps.createAccount(DESTRUCTION_NAME, SPECIAL_STATUS)

	return ps
}

func (ps *PaymentSystem) GetEmissionAccountIban() string {
	return ps.emissionAccount.Iban
}

func (ps *PaymentSystem) GetDestructionAccountIban() string {
	return ps.destructionAccount.Iban
}

func (ps *PaymentSystem) EmitMoney(kopecks int) {
	ps.emissionAccount.Balance.Kopecks += kopecks
}

func (ps *PaymentSystem) DestroyMoney(accountIban string, kopecks int) error {
	account, ok := ps.getAccount(accountIban)
	if !ok {
		return fmt.Errorf("not found account by iban")
	}

	//TODO: может ли быть баланс отрицательным при этой операции?
	account.Balance.Kopecks -= kopecks
	ps.destructionAccount.Balance.Kopecks += kopecks

	return nil
}

// ищет счет по IBAN номеру (поиск по ключу в мапе), возвращает вторым параметром ok (найдено-не найдено)
func (ps *PaymentSystem) getAccount(iban string) (*Account, bool) {
	account, ok := ps.accounts[iban]
	return account, ok
}

func (ps *PaymentSystem) CreateBaseAccountAndGetIban(name string) string {
	newBaseAccount := ps.createAccount(name, BASE_STATUS)
	return newBaseAccount.Iban
}

func (ps *PaymentSystem) createAccount(name, status string) *Account {
	newIban := iban.GenerateIban()

	newAccount := &Account{
		Iban:     newIban,
		Name:     name,
		IsActive: true,
		Status:   status,
	}

	ps.accounts[newIban] = newAccount
	return newAccount
}

func (ps *PaymentSystem) TransferMoney(senderAccountIban, receiverAccountIban string, kopecks int) error {
	senderAccount, ok := ps.getAccount(senderAccountIban)
	if !ok {
		return fmt.Errorf("not found sender account by iban")
	}
	receiverAccount, ok := ps.getAccount(receiverAccountIban)
	if !ok {
		return fmt.Errorf("not found receiver account by iban")
	}

	if !senderAccount.IsActive {
		return fmt.Errorf("sender account is blocked")
	}
	if !receiverAccount.IsActive {
		return fmt.Errorf("receiver account is blocked")
	}

	if senderAccount.Balance.Kopecks < kopecks {
		return fmt.Errorf("sender account does not have enough funds")
	}

	senderAccount.Balance.Kopecks -= kopecks
	receiverAccount.Balance.Kopecks += kopecks

	return nil
}

func (ps *PaymentSystem) TransferMoneyJSON(requestJSON string) error {
	var request TransferRequest
	if err := json.Unmarshal([]byte(requestJSON), &request); err != nil {
		return fmt.Errorf("error parsing json: %w", err)
	}

	if err := ps.TransferMoney(request.SenderIban, request.ReceiverIban, request.Money.Kopecks); err != nil {
		return err
	}

	return nil
}

func (ps *PaymentSystem) ListAccounts() (string, error) {
	accounts := make([]*Account, 0, len(ps.accounts))
	for _, account := range ps.accounts {
		accounts = append(accounts, account)
	}

	result, err := json.Marshal(accounts)
	if err != nil {
		return "", fmt.Errorf("error marshal list accounts to json: %w", err)
	}

	return string(result), nil
}
