package store

import (
	"easy-life-back-go/internal/pkg/store_codes"
	pkgStore "easy-life-back-go/pkg/store"
	"time"
)

type Store struct {
	store      pkgStore.Store
	storeCodes store_codes.StoreCodes
}

func NewStore(redis pkgStore.Store, storeCodes store_codes.StoreCodes) *Store {
	return &Store{
		store:      redis,
		storeCodes: storeCodes,
	}
}

func (r *Store) SetRegistrationCode(email, code string, attempt int) error {
	return r.storeCodes.SetWithTTL(
		GetKeyUserRegistrationCode(email),
		code,
		attempt,
		time.Minute*10,
	)
}

func (r *Store) GetRegistrationCode(email string) (string, int, error) {
	return r.storeCodes.GetWithTTL(GetKeyUserRegistrationCode(email))
}

func (r *Store) DelRegistrationCode(email string) error {
	return r.store.Del(GetKeyUserRegistrationCode(email))
}

func (r *Store) SetForgotRegistrationCode(email, code string) error {
	return nil
}

func (r *Store) SetJWTPairCode(id, code string) error {
	return nil
}
