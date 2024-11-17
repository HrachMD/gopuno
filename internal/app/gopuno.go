package gopuno

import "github.com/HrachMD/gopuno/internal/config"

type gopunoService interface {
}

type Gopuno struct {
	gopunoService gopunoService
}

func New(cfg *config.Config) *Gopuno {
	return &Gopuno{}
}

func (g *Gopuno) Start() {

}

func (g *Gopuno) Stop() {

}
