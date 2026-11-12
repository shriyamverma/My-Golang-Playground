package locker

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

type LockerService interface {
	DepositPackage(pkg *Package) (slotID string, otp string, err error)
	PickupPackage(slotID string, otp string) (*Package, error)
	RegisterSlot(id string, size LockerSize)
}

type Service struct {
	mu     sync.Mutex
	repo   LockerRepository
	otpTTL time.Duration
}

func NewService(repo LockerRepository, otpTTL time.Duration) *Service {
	return &Service{
		repo:   repo,
		otpTTL: otpTTL,
	}
}

func (s *Service) RegisterSlot(id string, size LockerSize) {
	s.repo.AddSlot(&Slot{
		ID:         id,
		Size:       size,
		IsOccupied: false,
	})
}

func (s *Service) DepositPackage(pkg *Package) (string, string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	slot, err := s.repo.FindAvailableSlot(pkg.Size)
	if err != nil {
		return "", "", err
	}

	otp, err := generateOTP(6)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate otp: %w", err)
	}

	slot.IsOccupied = true
	slot.Package = pkg
	slot.OTP = otp
	slot.ExpiresAt = time.Now().Add(s.otpTTL)

	if err := s.repo.UpdateSlot(slot); err != nil {
		return "", "", err
	}

	return slot.ID, otp, nil
}

func (s *Service) PickupPackage(slotID string, otp string) (*Package, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	slot, err := s.repo.GetSlotByID(slotID)
	if err != nil {
		return nil, err
	}

	if !slot.IsOccupied || slot.Package == nil {
		return nil, ErrSlotEmpty
	}

	if slot.OTP != otp || time.Now().After(slot.ExpiresAt) {
		return nil, ErrInvalidOTP
	}

	retrievedPackage := slot.Package

	// Reset slot
	slot.Package = nil
	slot.IsOccupied = false
	slot.OTP = ""
	slot.ExpiresAt = time.Time{}

	if err := s.repo.UpdateSlot(slot); err != nil {
		return nil, err
	}

	return retrievedPackage, nil
}

func generateOTP(length int) (string, error) {
	const digits = "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[num.Int64()]
	}
	return string(otp), nil
}
