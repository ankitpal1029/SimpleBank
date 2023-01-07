package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ankitpal1029/SimpleBank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) (Account, CreateAccountParams){
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)

	return account, arg 
}

func TestCreateAccount(t *testing.T){
	account, arg:= createRandomAccount(t)

	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T){
	account1, _:= createRandomAccount(t)
	account2, err:= testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.ID, account2.ID)
}

func TestUpdateAccount(t *testing.T){
	account1, _:= createRandomAccount(t)

	arg := UpdateAccountParams{
		ID: account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg, )

	require.NoError(t, err)

	require.NotEmpty(t, account2)

	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.ID, account2.ID)
}

func TestDeleteAccount(t * testing.T){
	account1, _ := createRandomAccount(t)

	account2, err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	account3, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account3)
}

func TestListAccounts(t *testing.T){
	for i := 0; i < 10; i++{
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(accounts), 5)

	for _, account := range accounts{
		require.NotEmpty(t, account)
	}
}