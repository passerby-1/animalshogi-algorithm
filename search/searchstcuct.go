package search

import "animalshogi/models"

type HashTree struct {
	Hash    uint64         `json:"hash"`
	WinLose int            `json:"winlose"`
	Boards  []models.Board `json:"boards"`
	Next    []uint64       `json:"next"` // ポインタでやるのちょっと面倒そうだったのでとりあえず
	Prev    []uint64       `json:"prev"`
}
