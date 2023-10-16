package storage

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/storage/postgres"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/storage/repo"
)

type StorageI interface {
	UnicalPro() repo.MedicalStorageI
}

type StoragePg struct {
	unicalPro repo.MedicalStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &StoragePg{
		unicalPro: postgres.NewArtMedical(db),
	}
}

func (s *StoragePg) UnicalPro() repo.MedicalStorageI {
	return s.unicalPro
}
