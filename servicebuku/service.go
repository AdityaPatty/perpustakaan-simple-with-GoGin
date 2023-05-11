package servicebuku

import (
	"perpustakaan/modelbuku"
	"perpustakaan/repositorybuku"
)

type PerpustakaanResponse struct {
	ID            int    `json:"id"`
	Judulbuku     string `json:"judul_buku"`
	Deskripsibuku string `json:"deskripsi_buku"`
	Isbn          string `json:"isbn"`
	Issn          string `Json:"issn"`
	Bahasabuku    string `Json:"bahasabuku"`
}

type BukuService interface {
	Get(id int) (*PerpustakaanResponse, error)
	Gets() ([]*PerpustakaanResponse, error)
	Create(perpustakaan modelbuku.Perpustakaan) (*PerpustakaanResponse, error)
	Update(id int, perpustakaan modelbuku.Perpustakaan) (*PerpustakaanResponse, error)
	Delete(id int) (*PerpustakaanResponse, error)
}

type bukuService struct {
	db repositorybuku.Database
}

func NewBukuService(db repositorybuku.Database) *bukuService {
	return &bukuService{db: db}
}

func (s *bukuService) Create(perpustakaan modelbuku.Perpustakaan) (*PerpustakaanResponse, error) {
	buku, err := s.db.Create(perpustakaan)

	if err != nil {
		return nil, err
	}

	return PerpustakaanToPerpustakaanResponse(*buku), nil
}

func (s *bukuService) Get(id int) (*PerpustakaanResponse, error) {
	perpustakaan, err := s.db.Get(id)

	if err != nil {
		return nil, err
	}

	return PerpustakaanToPerpustakaanResponse(*perpustakaan), nil
}

func (s *bukuService) Gets() ([]*PerpustakaanResponse, error) {
	perpustakaans, err := s.db.Gets()

	if err != nil {
		return nil, err
	}

	var perpus []*PerpustakaanResponse
	for _, perpustakaan := range perpustakaans {
		perpus = append(perpus, PerpustakaanToPerpustakaanResponse(*perpustakaan))
	}

	return perpus, nil
}

func (s *bukuService) Update(id int, perpustakaan modelbuku.Perpustakaan) (*PerpustakaanResponse, error) {
	if err := perpustakaan.IsValid(); err != nil {
		return nil, err
	}

	buku, err := s.db.Update(id, perpustakaan)

	if err != nil {
		return nil, err
	}

	return PerpustakaanToPerpustakaanResponse(*buku), nil
}

func (s *bukuService) Delete(id int) (*PerpustakaanResponse, error) {
	buku, err := s.db.Delete(id)

	if err != nil {
		return nil, err
	}

	return PerpustakaanToPerpustakaanResponse(*buku), nil
}

func PerpustakaanToPerpustakaanResponse(buku modelbuku.Perpustakaan) *PerpustakaanResponse {
	return &PerpustakaanResponse{
		ID:            buku.ID,
		Judulbuku:     buku.Judulbuku,
		Deskripsibuku: buku.Deskripsibuku,
		Isbn:          buku.Isbn,
		Issn:          buku.Issn,
		Bahasabuku:    buku.Bahasabuku,
	}
}
