package datafiles

import (
	"time"

	"github.com/betoth/geradorarquivos/internal/util"
	"github.com/google/uuid"
)

type Composition struct {
	Composition      int    `json:"composition"`
	CompositionType  string `json:"composition_type"`
	CompositionValue int    `json:"composition_value"`
}

type Installment struct {
	Instalment            int           `json:"instalment"`
	ValueInstalment       int           `json:"value_instalment"`
	CreateDate            time.Time     `json:"create_date"`
	DueDate               time.Time     `json:"due_date"`
	InstalmentComposition []Composition `json:"instalment_composition"`
}

type Loan struct {
	IdExternal   uuid.UUID     `json:"id_external"`
	CreateDate   time.Time     `json:"create_date"`
	DueDate      time.Time     `json:"due_date"`
	Value        int           `json:"value"`
	Installments []Installment `json:"installments"`
}

var LoanData = &Loan{
	IdExternal: uuid.New(),
	CreateDate: time.Now(),
	DueDate:    util.DatePlus30,
	Value:      3000,
	Installments: []Installment{
		{
			Instalment:      1,
			ValueInstalment: 1000,
			CreateDate:      time.Now(),
			DueDate:         util.DatePlus30,
			InstalmentComposition: []Composition{
				{Composition: 1, CompositionType: "Type1", CompositionValue: 100},
				{Composition: 2, CompositionType: "Type2", CompositionValue: 200},
			},
		},
	},
}
