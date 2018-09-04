package backend

import (
	"github.com/solcates/gobwa/pkg/bwa"
	"github.com/solcates/gobwa/pkg/bwa/messages"
)

type Spa struct {
	address string
	port    int
	status  *messages.Status
}
type Finder interface {
	Find() (err error)
	Status(address string) (status *messages.Status)
}

func NewSpaFinder() *SpaFinder {
	return &SpaFinder{
		spas: make(map[string]*Spa),
	}
}

type SpaFinder struct {
	spas map[string]*Spa
}

func (sp *SpaFinder) Find() (err error) {

	// Replace the current spas map
	// First find all BWA Spas on the network...
	sp.spas, err = sp.findBWASpas()

	// Respond

	return
}

func (sp *SpaFinder) findBWASpas() (spas map[string]*Spa, err error) {
	spas = make(map[string]*Spa)
	// Find all BWA
	var foundbwas []string
	bc := bwa.NewDiscoverer()
	foundbwas, err = bc.Discover()
	if err != nil {
		return
	}
	for _, spa := range foundbwas {
		c := bwa.NewBalbowClient(spa, 4257)
		if err = c.Connect(); err != nil {
			return
		}
		spas[spa] = &Spa{
			address: spa,
			port:    4257,
		}
	}
	return
}

func (sp *SpaFinder) Status(address string) (status *messages.Status) {

	spa, exists := sp.spas[address]
	if !exists {
		return
	}
	return spa.status
}
