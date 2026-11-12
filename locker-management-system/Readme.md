# Locker Management Service (Go)

An in-process domain library implementing a thread-safe **Locker Management System** for e-commerce deliveries and customer pickups/returns.

---

## Problem Statement

Design and implement an in-memory Locker Management Service for an e-commerce platform. Customers can opt to have their packages delivered to a nearby locker system and collect them at their convenience. The same locker system also supports customer return drop-offs.

### Core Flows
1. **Delivery Flow (Drop-off):**
   - A delivery agent drops off a package at the locker station.
   - The system allocates an available locker slot that accommodates the package size (using a best-fit strategy).
   - A secure One-Time Password (OTP) is generated and shared with the user alongside the assigned locker details.
2. **Pickup Flow (Collection):**
   - The user visits the locker station and enters the slot ID and OTP.
   - Upon verification, the locker unlocks, the package is retrieved, and the locker slot is freed for future orders.
3. **Return Flow:**
   - Supports reverse logistics where a user initiates a return drop-off, generating an OTP for the delivery partner to collect the returned item.

### Constraints
- **No external dependencies or web frameworks:** Implemented as a plain in-process Go domain library (no REST APIs or external databases required).
- **Concurrency Safe:** Safe for concurrent deposit and pickup operations.

---

## What Was Implemented

### 1. Project Architecture
The solution follows clean separation of concerns across domain models, data persistence, and business logic:

```text
assignment/
├── go.mod
├── main.go                 # Runnable entry point demonstrating delivery & pickup flows
├── locker/
│   ├── model.go            # Domain entities (Package, Slot, LockerSize) & sentinel errors
│   ├── repository.go       # Thread-safe in-memory storage & best-fit slot lookup
│   └── service.go          # Core business logic, OTP generation, & expiry validation
└── README.md
```

### 2. Key Features & Design Decisions
- **Best-Fit Slot Allocation (`FindAvailableSlot`):**
  - Matches incoming packages (`Small`, `Medium`, `Large`, `ExtraLarge`) to the smallest available locker slot that can fit the package (`slot.Size >= requiredSize`), minimizing wasted capacity.
- **Cryptographically Secure OTPs (`crypto/rand`):**
  - Generates 6-digit numeric OTPs using `crypto/rand` paired with configurable TTL expiration (`ExpiresAt`).
- **Concurrency Control (`sync.Mutex` / `sync.RWMutex`):**
  - Ensures atomic state transitions during slot allocation, OTP verification, and slot release so two packages cannot claim the same slot simultaneously.
- **Clean State Reset:**
  - Once a package is picked up with a valid OTP, the slot's package reference, OTP, and occupancy flag are immediately cleared so the locker can be reused.
- **Idiomatic Error Handling:**
  - Exposes clear sentinel errors (`ErrNoAvailableLocker`, `ErrInvalidOTP`, `ErrLockerNotFound`, `ErrSlotEmpty`) for predictable testing and caller handling.

---

## Running the Project

```bash
# Build all packages
go build ./...

# Run the demonstration flow
go run main.go

# Run unit tests
go test -v ./...
```