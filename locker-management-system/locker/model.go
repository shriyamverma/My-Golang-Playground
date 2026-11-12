package locker

import (
	"errors"
	"time"
)

type LockerSize int

const (
	Small LockerSize = iota + 1
	Medium
	Large
	ExtraLarge
)

var (
	ErrNoAvailableLocker = errors.New("no available locker slot for package size")
	ErrInvalidOTP        = errors.New("invalid or expired OTP")
	ErrLockerNotFound    = errors.New("locker slot not found")
	ErrSlotOccupied      = errors.New("locker slot already occupied")
	ErrSlotEmpty         = errors.New("locker slot is empty")
)

type Package struct {
	ID        string
	Size      LockerSize
	CreatedAt time.Time
}

type Slot struct {
	ID         string
	Size       LockerSize
	IsOccupied bool
	Package    *Package
	OTP        string
	ExpiresAt  time.Time
}
