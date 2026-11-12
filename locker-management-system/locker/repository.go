package locker

import "sync"

type LockerRepository interface {
	AddSlot(slot *Slot)
	FindAvailableSlot(size LockerSize) (*Slot, error)
	GetSlotByID(slotID string) (*Slot, error)
	UpdateSlot(slot *Slot) error
}

type InMemoryLockerRepository struct {
	mu    sync.RWMutex
	slots map[string]*Slot
}

func NewInMemoryLockerRepository() *InMemoryLockerRepository {
	return &InMemoryLockerRepository{
		slots: make(map[string]*Slot),
	}
}

func (r *InMemoryLockerRepository) AddSlot(slot *Slot) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.slots[slot.ID] = slot
}

func (r *InMemoryLockerRepository) FindAvailableSlot(requiredSize LockerSize) (*Slot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var bestFit *Slot
	for _, slot := range r.slots {
		if !slot.IsOccupied && slot.Size >= requiredSize {
			if bestFit == nil || slot.Size < bestFit.Size {
				bestFit = slot
			}
		}
	}

	if bestFit == nil {
		return nil, ErrNoAvailableLocker
	}
	return bestFit, nil
}

func (r *InMemoryLockerRepository) GetSlotByID(slotID string) (*Slot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	slot, exists := r.slots[slotID]
	if !exists {
		return nil, ErrLockerNotFound
	}
	return slot, nil
}

func (r *InMemoryLockerRepository) UpdateSlot(slot *Slot) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.slots[slot.ID] = slot
	return nil
}
