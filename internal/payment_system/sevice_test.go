package payment_system

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewPaymentSystem(t *testing.T) {
	ps := NewPaymentSystem()

	//сгенерировано 2 дефолтных счета
	assert.Equal(t, len(ps.accounts), 2)
}

func Test_CreateAccount(t *testing.T) {
	ps := NewPaymentSystem()

	account := Account{
		Name:   "testName",
		Status: "testStatus",
	}

	newAccount := ps.createAccount(account.Name, account.Status)
	assert.Equal(t, account.Name, newAccount.Name)
	assert.Equal(t, 0, newAccount.Balance.Kopecks)
	assert.Equal(t, true, newAccount.IsActive)
	assert.Equal(t, account.Status, newAccount.Status)
	assert.Equal(t, len(ps.accounts), 3)
}

func Test_DestroyMoney(t *testing.T) {
	ps := NewPaymentSystem()

	account := ps.createAccount("test", BASE_STATUS)

	initBalance := 1042
	account.Balance.Kopecks = initBalance

	initDestroy := 570
	err := ps.DestroyMoney(account.Iban, initDestroy)
	assert.NoError(t, err)

	assert.Equal(t, initDestroy, ps.destructionAccount.Balance.Kopecks)
	assert.Equal(t, initBalance-initDestroy, account.Balance.Kopecks)

	err = ps.DestroyMoney("00", 0)
	assert.Contains(t, err.Error(), "not found")
}

func Test_TransferMoney(t *testing.T) {
	ps := NewPaymentSystem()

	senderAccount := ps.createAccount("sender", BASE_STATUS)
	addKopecks := 5537
	senderAccount.Balance.Kopecks = addKopecks
	receiverAccount := ps.createAccount("receiver", BASE_STATUS)
	transferKopecks := 214

	err := ps.TransferMoney(senderAccount.Iban, receiverAccount.Iban, transferKopecks)
	assert.NoError(t, err)
	assert.Equal(t, addKopecks-transferKopecks, senderAccount.Balance.Kopecks)
	assert.Equal(t, transferKopecks, receiverAccount.Balance.Kopecks)

	err = ps.TransferMoney("00", receiverAccount.Iban, 0)
	assert.Contains(t, err.Error(), "not found")
	err = ps.TransferMoney(senderAccount.Iban, "00", 0)
	assert.Contains(t, err.Error(), "not found")

	senderAccount.IsActive = false
	err = ps.TransferMoney(senderAccount.Iban, receiverAccount.Iban, 0)
	assert.Contains(t, err.Error(), "is blocked")
	senderAccount.IsActive = true
	receiverAccount.IsActive = false
	err = ps.TransferMoney(senderAccount.Iban, receiverAccount.Iban, 0)
	assert.Contains(t, err.Error(), "is blocked")
	receiverAccount.IsActive = true

	err = ps.TransferMoney(senderAccount.Iban, receiverAccount.Iban, addKopecks+1020)
	assert.Contains(t, err.Error(), "enough funds")
}

func Test_TransferMoneyJSON(t *testing.T) {
	ps := NewPaymentSystem()

	senderAccount := ps.createAccount("sender", BASE_STATUS)
	addKopecks := 5000
	senderAccount.Balance.Kopecks = addKopecks
	receiverAccount := ps.createAccount("receiver", BASE_STATUS)

	requestJson := fmt.Sprintf("{\"sender_iban\": \"%s\",\"receiver_iban\": \"%s\",\"money\": \"5.75\"}", senderAccount.Iban, receiverAccount.Iban)

	err := ps.TransferMoneyJSON(requestJson)
	assert.NoError(t, err)
}

func TestPaymentSystem_ListAccounts(t *testing.T) {
	ps := NewPaymentSystem()

	account := ps.createAccount("test", BASE_STATUS)
	account.Balance.Kopecks = 100

	accountsString, err := ps.ListAccounts()
	assert.NoError(t, err)

	var accounts []*Account
	err = json.Unmarshal([]byte(accountsString), &accounts)
	assert.NoError(t, err)

	assert.Len(t, accounts, len(ps.accounts))
}
