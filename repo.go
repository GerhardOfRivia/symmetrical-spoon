package main

import "fmt"

var currentID int

var transactions []Transaction

// Give us some seed data
func init() {
	RepoCreateTransaction(Transaction{Name: "Write presentation"})
	RepoCreateTransaction(Transaction{Name: "Host meetup"})
}

// RepoFindTransaction find a transaction
func RepoFindTransaction(id int) Transaction {
	for _, t := range transactions {
		if t.ID == id {
			return t
		}
	}
	// return empty Transaction if not found
	return Transaction{}
}

// RepoCreateTransaction this is bad, I don't think it passes race condtions
func RepoCreateTransaction(t Transaction) Transaction {
	currentID++
	t.ID = currentID
	transactions = append(transactions, t)
	return t
}

// RepoDestroyTransaction remove a trasaction
func RepoDestroyTransaction(id int) error {
	for i, t := range transactions {
		if t.ID == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Transaction with id of %d to delete", id)
}