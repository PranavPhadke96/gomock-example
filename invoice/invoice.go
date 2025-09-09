package invoice

import "time"

// CalculateLateFee computes the late fee in pounds (integers only).
// Rules:
//   - £2 per calendar day late.
//   - Max fee is 10% of invoice amount (floor).
//   - If amount <= 0 or paid on/before due date -> fee = 0.
//   - Dates are compared as local calendar dates (time-of-day ignored).
func CalculateLateFee(amount int, dueDate, paidDate time.Time) int {
	if amount <= 0 {
		return 0
	}
	lateDays := daysLate(dueDate, paidDate)
	if lateDays <= 0 {
		return 0
	}

	fee := lateDays * 2
	capFee := amount / 10 // 10% cap (integer division)

	if fee > capFee {
		return capFee
	}
	return fee
}

// daysLate returns whole calendar days (paid - due).
// Positive means "late", zero or negative means "on time or early".
func daysLate(dueDate, paidDate time.Time) int {
	due := truncateToDate(dueDate)
	paid := truncateToDate(paidDate)
	return int(paid.Sub(due).Hours() / 24)
}

func truncateToDate(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}
