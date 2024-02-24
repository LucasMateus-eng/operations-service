package address

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

type BrazilianState int64

const (
	UNDEFINED BrazilianState = iota
	AC
	AL
	AP
	AM
	BA
	CE
	ES
	GO
	MA
	MT
	MS
	MG
	PA
	PB
	PR
	PE
	PI
	RJ
	RN
	RS
	RO
	RR
	SC
	SP
	SE
	TO
	DF
)

var (
	brazilianStateMap = map[string]BrazilianState{
		"UNDEFINED":           UNDEFINED,
		"ACRE":                AC,
		"ALAGOAS":             AL,
		"AMAPÁ":               AP,
		"AMAZONAS":            AM,
		"BAHIA":               BA,
		"CEARÁ":               CE,
		"ESPÍRITO SANTO":      ES,
		"GOIÁS":               GO,
		"MARANHÃO":            MA,
		"MATO GROSSO":         MT,
		"MATO GROSSO DO SUL":  MS,
		"MINAS GERAIS":        MG,
		"PARÁ":                PA,
		"PARAÍBA":             PB,
		"PARANÁ":              PR,
		"PERNAMBUCO":          PE,
		"PIAUÍ":               PI,
		"RIO DE JANEIRO":      RJ,
		"RIO GRANDE DO NORTE": RN,
		"RIO GRANDE DO SUL":   RS,
		"RONDÔNIA":            RO,
		"RORAIMA":             RR,
		"SANTA CATARINA":      SC,
		"SÃO PAULO":           SP,
		"SERGIPE":             SE,
		"TOCANTINS":           TO,
		"DISTRITO FEDERAL":    DF,
	}

	brazilianStateList = []BrazilianState{
		AC,
		AL,
		AP,
		AM,
		BA,
		CE,
		ES,
		GO,
		MA,
		MT,
		MS,
		MG,
		PA,
		PB,
		PR,
		PE,
		PI,
		RJ,
		RN,
		RS,
		RO,
		RR,
		SC,
		SP,
		SE,
		TO,
		DF,
	}
)

func GetBrazilianState(name string) (*BrazilianState, error) {
	for key, value := range brazilianStateMap {
		if strings.EqualFold(key, name) {
			return &value, nil
		}
	}

	return nil, fmt.Errorf("the given brazilian state [%s] is non-existent in the map of valid values", name)
}

func (bs BrazilianState) Change(new BrazilianState) (*BrazilianState, error) {
	if new == UNDEFINED {
		return nil, errors.New("the new brazilian state cannot be equal to undefined")
	}

	if new == bs {
		return nil, errors.New("the new brazilian state cannot be the same as the old one")
	}

	if !slices.Contains(brazilianStateList, new) {
		return nil, fmt.Errorf("the new brazilian state [%v] is not present in the available list of valid values", new)
	}

	return &new, nil
}

func (bs BrazilianState) String() string {
	switch bs {
	case AC:
		return "ACRE"
	case AL:
		return "ALAGOAS"
	case AP:
		return "AMAPÁ"
	case AM:
		return "AMAZONAS"
	case BA:
		return "BAHIA"
	case CE:
		return "CEARÁ"
	case ES:
		return "ESPÍRITO SANTO"
	case GO:
		return "GOIÁS"
	case MA:
		return "MARANHÃO"
	case MT:
		return "MATO GROSSO"
	case MS:
		return "MATO GROSSO DO SUL"
	case MG:
		return "MINAS GERAIS"
	case PA:
		return "PARÁ"
	case PB:
		return "PARAÍBA"
	case PR:
		return "PARANÁ"
	case PE:
		return "PERNAMBUCO"
	case PI:
		return "PIAUÍ"
	case RJ:
		return "RIO DE JANEIRO"
	case RN:
		return "RIO GRANDE DO NORTE"
	case RS:
		return "RIO GRANDE DO SUL"
	case RO:
		return "RONDÔNIA"
	case RR:
		return "RORAIMA"
	case SC:
		return "SANTA CATARINA"
	case SP:
		return "SÃO PAULO"
	case SE:
		return "SERGIPE"
	case TO:
		return "TOCATINS"
	case DF:
		return "DISTRITO FEDERAL"
	}

	return "UNDEFINED"
}

type Address struct {
	ID           int
	UserID       int
	Locality     string
	Number       string
	Complement   string
	Neighborhood string
	City         string
	State        BrazilianState
	CEP          string
	Country      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Reading interface {
	GetById(ctx context.Context, id int) (*Address, error)
	GetByUserID(ctx context.Context, userID int) (*Address, error)
}

type Writing interface {
	Create(ctx context.Context, a *Address) (int, error)
	Update(ctx context.Context, a *Address) error
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Reading
	Writing
}

type UseCase interface {
	GetById(ctx context.Context, id int) (*Address, error)
	GetByUserID(ctx context.Context, userID int) (*Address, error)
	Create(ctx context.Context, a *Address) (int, error)
	Update(ctx context.Context, a *Address) error
	Delete(ctx context.Context, id int) error
}
