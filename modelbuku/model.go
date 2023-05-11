package modelbuku

import (
	"errors"
)

type IPerpustakaan interface {
	Get(id int) (*Perpustakaan, error)
	Gets() (Perpustakaans, error)
	Create(perpustakaan Perpustakaan) error
	Update(id int, perpustakaan Perpustakaan) error
}

type Perpustakaans []*Perpustakaan

type Perpustakaan struct {
	ID            int
	Judulbuku     string
	Deskripsibuku string
	Isbn          string
	Issn          string
	Bahasabuku    string
}

func (buku *Perpustakaan) TableName() string {
	return "Perpustakaans"
}

func (buku *Perpustakaan) IsValid() error {
	if buku.ID == 0 ||
		buku.Judulbuku == "" ||
		buku.Deskripsibuku == "" ||
		buku.Isbn == "" ||
		buku.Issn == "" ||
		buku.Bahasabuku == "" {
		return errors.New("Data Buku Invalid")
	}

	return nil
}
