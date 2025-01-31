/*
Copyright 2018 Christian Banse

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExpense(t *testing.T) {
	expense := NewExpense()
	expense.Amount = 2.0
	expense.Person = "Christian"

	// assert default currency
	assert.Equal(t, expense.Currency, DefaultCurrency, "A new expense should have the default currency")

	fmt.Printf("%s spent %f %s", expense.Person, expense.Amount, expense.Currency)
}
