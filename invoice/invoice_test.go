package invoice

import (
	"testing"
	"time"
)

func mustDate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

func Test_OnTime_NoFee(t *testing.T) {
	due := mustDate(2025, 9, 1)
	paid := mustDate(2025, 9, 1)
	got := CalculateLateFee(100, due, paid)
	if got != 0 {
		t.Fatalf("want 0, got %d", got)
	}
}

func Test_TwoDaysLate_FeeIs4(t *testing.T) {
	due := mustDate(2025, 9, 1)
	paid := mustDate(2025, 9, 3) // 2 days late
	got := CalculateLateFee(100, due, paid)
	if got != 4 {
		t.Fatalf("want 4, got %d", got)
	}
}

func Test_FeeIsCappedAt10Percent(t *testing.T) {
	due := mustDate(2025, 9, 1)
	paid := mustDate(2025, 9, 30) // 29 days late → 58
	got := CalculateLateFee(50, due, paid) // cap 5
	if got != 5 {
		t.Fatalf("want 5 (cap), got %d", got)
	}
}

func Test_NoFeeForZeroAmount(t *testing.T) {
	due := mustDate(2025, 9, 1)
	paid := mustDate(2025, 9, 10)
	got := CalculateLateFee(0, due, paid)
	if got != 0 {
		t.Fatalf("want 0 for zero amount, got %d", got)
	}
}

func Test_PaidBeforeDue_NoFee(t *testing.T) {
	due := mustDate(2025, 9, 10)
	paid := mustDate(2025, 9, 5)
	got := CalculateLateFee(100, due, paid)
	if got != 0 {
		t.Fatalf("want 0 when paid before due, got %d", got)
	}
}
