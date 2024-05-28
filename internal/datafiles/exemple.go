package datafiles

import (
	"time"

	"github.com/betoth/geradorarquivos/internal/util"
	"github.com/google/uuid"
)

type ExempleDetails struct {
	ExempleDetails      int    `json:"composition"`
	ExempleDetailsType  string `json:"composition_type"`
	ExempleDetailsValue int    `json:"composition_value"`
}

type ExempleComposition struct {
	Compositions           int              `json:"instalment"`
	CompositionsValue      int              `json:"value_instalment"`
	CompositionsCreateDate time.Time        `json:"create_date"`
	CompositionsDueDate    time.Time        `json:"due_date"`
	CompositionsDetails    []ExempleDetails `json:"instalment_composition"`
}

type Exemple struct {
	ExempleId          uuid.UUID            `json:"id_external"`
	ExempleCreateDate  time.Time            `json:"create_date"`
	ExempleDueDate     time.Time            `json:"due_date"`
	ExempleValue       int                  `json:"value"`
	ExempleComposition []ExempleComposition `json:"installments"`
}

var ExempleData = &Exemple{
	ExempleId:         uuid.New(),
	ExempleCreateDate: time.Now(),
	ExempleDueDate:    util.DatePlus30,
	ExempleValue:      3000,
	ExempleComposition: []ExempleComposition{
		{
			Compositions:           1,
			CompositionsValue:      1000,
			CompositionsCreateDate: time.Now(),
			CompositionsDueDate:    util.DatePlus30,
			CompositionsDetails: []ExempleDetails{
				{ExempleDetails: 1, ExempleDetailsType: "Type1", ExempleDetailsValue: 100},
				{ExempleDetails: 2, ExempleDetailsType: "Type2", ExempleDetailsValue: 200},
			},
		},
	},
}
