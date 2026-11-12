package main

import (
	"fmt"
	"time"

	"locker-management-system/locker"
)

func main() {
	repo := locker.NewInMemoryLockerRepository()
	service := locker.NewService(repo, 24*time.Hour)

	// Configure locker slots
	service.RegisterSlot("SLOT-S1", locker.Small)
	service.RegisterSlot("SLOT-M1", locker.Medium)
	service.RegisterSlot("SLOT-L1", locker.Large)

	// --- Scenario 1: Standard Delivery & Customer Pickup ---
	fmt.Println("=== Scenario 1: Standard Delivery & Pickup ===")
	pkg := &locker.Package{
		ID:        "PKG-98213",
		Size:      locker.Medium,
		CreatedAt: time.Now(),
	}

	slotID, otp, err := service.DepositPackage(pkg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Delivery] Deposited %s to %s with OTP: %s\n", pkg.ID, slotID, otp)

	collectedPkg, err := service.PickupPackage(slotID, otp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Customer] Collected package: %s (Size: %d)\n\n", collectedPkg.ID, collectedPkg.Size)

	// --- Scenario 2: Customer Return & Delivery Agent Collection ---
	fmt.Println("=== Scenario 2: Return Order Flow ===")
	returnPkg := &locker.Package{
		ID:        "RET-44109",
		Size:      locker.Small,
		CreatedAt: time.Now(),
	}

	// 1. Customer initiates a return and deposits the item into an allocated locker
	retSlotID, retOTP, err := service.DepositPackage(returnPkg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Customer Return] Deposited return package %s into %s. Return OTP generated for courier: %s\n",
		returnPkg.ID, retSlotID, retOTP)

	// 2. Delivery agent arrives at the locker and collects the return item using the OTP
	pickedUpReturn, err := service.PickupPackage(retSlotID, retOTP)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Delivery Pickup] Courier collected return package %s from %s. Slot is now free.\n",
		pickedUpReturn.ID, retSlotID)
}
