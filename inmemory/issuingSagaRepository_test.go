package inmemory

import (
	"sync"
	"testing"

	covid "github.com/daneidmark/ddd-covid-pass-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCanSaveIssuingSaga(t *testing.T) {
	repo := IssuingSagaRepository{DB: map[covid.PatientReference]covid.CovidPassId{}, mx: sync.RWMutex{}}

	ref := covid.PatientReference("Dan")
	id := covid.CovidPassId(uuid.New())
	repo.Store(ref, id)

	id2 := repo.Find(ref)

	assert.Equal(t, id, id2)
}
