package validate_test

import (
	"math/rand/v2"
	"testing"

	"github.com/copartner6412/input/pseudorandom"
	"github.com/copartner6412/input/validate"
)

const (
	minPortAllowed uint = 0
	maxPortAllowed uint = 65535
)

func FuzzPortForValidPseudorandomInput(f *testing.F) {
	f.Fuzz(func (t *testing.T, seed1, seed2 uint64, min, max uint)  {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		minPort := minPortAllowed + min%(maxPortAllowed - minPortAllowed + 1)
		maxPort := minPort + max%(maxPortAllowed - minPort + 1)
		port1, err := pseudorandom.Port(r1, uint16(minPort), uint16(maxPort))
		if err != nil {
			t.Fatalf("error generating a pseudo-random port: %v", err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		port2, err := pseudorandom.Port(r2, uint16(minPort), uint16(maxPort))
		if err != nil {
			t.Fatalf("error regenerating the pseudo-random port: %v", err)
		}
		if port1 != port2 {
			t.Fatal("not deterministic")
		}
	})
}

func FuzzPortWellKnownSuccessfulForValidPseudorandomPort(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		port1 := pseudorandom.PortWellKnown(r1)
		err := validate.PortWellKnown(port1)
		if err != nil {
			t.Errorf("expected no error for valid pseudo-random well-known port %d, but got error: %v", port1, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		port2 := pseudorandom.PortWellKnown(r2)
		if port1 != port2 {
			t.Errorf("not deterministic, expected: %d, got: %d", port1, port2)
		}
	})
}


func FuzzPortNotWellKnownSuccessfulForValidPseudorandomWellKnownPort(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		port1 := pseudorandom.PortNotWellKnown(r1)
		err := validate.PortNotWellKnown(port1)
		if err != nil {
			t.Errorf("expected no error for valid pseudo-random not-well-known port %d, but got error: %v", port1, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		port2 := pseudorandom.PortNotWellKnown(r2)
		if port1 != port2 {
			t.Errorf("not deterministic, expected: %d, got: %d", port1, port2)
		}
	})
}

func FuzzPortRegisteredSuccessfulForValidPseudorandomRegisteredPort(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		port1 := pseudorandom.PortRegistered(r1)
		err := validate.PortRegistered(port1)
		if err != nil {
			t.Errorf("expected no error for valid pseudo-random registered port %d, but got error: %v", port1, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		port2 := pseudorandom.PortRegistered(r2)
		if port1 != port2 {
			t.Errorf("not deterministic, expected: %d, got: %d", port1, port2)
		}
	})
}

func FuzzPortPrivateSuccessfulForValidPseudorandomPrivatePort(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed1, seed2 uint64) {
		r1 := rand.New(rand.NewPCG(seed1, seed2))
		port1 := pseudorandom.PortPrivate(r1)
		err := validate.PortPrivate(port1)
		if err != nil {
			t.Fatalf("expected no error for valid pseudo-random private port %d, but got error: %v", port1, err)
		}
		r2 := rand.New(rand.NewPCG(seed1, seed2))
		port2 := pseudorandom.PortPrivate(r2)
		if port1 != port2 {
			t.Fatalf("not deterministic")
		}
	})
}

func TestPortWellKnownSuccessfulForValidInput(t *testing.T) {
	testCaseGroups := map[string][]uint16{
		"Common Services": {
			80,  // HTTP
			443, // HTTPS
			22,  // SSH
			25,  // SMTP
			110, // POP3
			143, // IMAP
			53,  // DNS
			21,  // FTP
		},
		"File Transfer": {
			20,  // FTP data transfer
			69,  // TFTP
			873, // rsync
		},
		"Web and Mail": {
			587, // SMTP with encryption (submission)
			465, // SMTP over SSL
			993, // IMAP over SSL
			995, // POP3 over SSL
		},
	}

	for name, testCaseGroup := range testCaseGroups {
		t.Run(name, func(t *testing.T) {
			for _, testCase := range testCaseGroup {
				err := validate.PortWellKnown(testCase)
				if err != nil {
					t.Errorf("expected no error for valid input %d, but got error: %v", testCase, err)
				}
			}
		})
	}
}

func TestPortWellKnownFailsForInvalidInput(t *testing.T) {
	testCaseGroups := map[string][]uint16{
		"PortsAboveRange": {
			1024,  // Just above the upper boundary for well-known ports
			65535, // Maximum value for uint16, far above well-known range
		},
		"CommonInvalidPorts": {
			49152, // Starting range for dynamic/private ports
		},
	}

	for name, testCaseGroup := range testCaseGroups {
		t.Run(name, func(t *testing.T) {
			for _, testCase := range testCaseGroup {
				err := validate.PortWellKnown(testCase)
				if err == nil {
					t.Errorf("expected error for invalid input %d, but got no error", testCase)
				}
			}
		})
	}
}

func TestPortNotWellKnownSuccessfulForValidInput(t *testing.T) {
	testCaseGroups := map[string][]uint16{
		"EdgeCases": {
			1024,
			65535, // Maximum value for uint16
		},
		"CommonValidPorts": {
			49152, // Starting range for dynamic/private ports
		},
	}

	for name, testCaseGroup := range testCaseGroups {
		t.Run(name, func(t *testing.T) {
			for _, testCase := range testCaseGroup {
				err := validate.PortNotWellKnown(testCase)
				if err != nil {
					t.Errorf("expected no error for valid input %d, but got error: %v", testCase, err)
				}
			}
		})
	}
}

func TestPortNotWellKnownFailsForInvalidInput(t *testing.T) {
	testCaseGroups := map[string][]uint16{
		"PortsBelowRange": {
			0,
			1023,
		},
		"CommonInvalidPorts": {
			443,
		},
	}
	for name, testCaseGroup := range testCaseGroups {
		t.Run(name, func(t *testing.T) {
			for _, testCase := range testCaseGroup {
				err := validate.PortNotWellKnown(testCase)
				if err == nil {
					t.Errorf("expected error for invalid input %d, but got no error", testCase)
				}
			}
		})
	}
}

func TestPortRegisterSuccessfulForValidInput(t *testing.T) {
	testCaseGroups := map[string][]uint16{
		"EdgeCases": {
			1024,
			49151,
		},
		"CommonValidPorts": {
			30306, // MySQL
		},
	}

	for name, testCaseGroup := range testCaseGroups {
		t.Run(name, func(t *testing.T) {
			for _, testCase := range testCaseGroup {
				err := validate.PortRegistered(testCase)
				if err != nil {
					t.Errorf("expected no error for valid input %d, but got error: %v", testCase, err)
				}
			}
		})
	}
}

func TestPortRegisteredFailsForInvalidInput(t *testing.T) {
	testCaseGroups := map[string][]uint16{
		"PortsBelowRange": {
			0,
			1023,
		},
		"PortsAboveRange": {
			49152,
		},
		"CommonInvalidPorts": {
			443,
		},
	}
	for name, testCaseGroup := range testCaseGroups {
		t.Run(name, func(t *testing.T) {
			for _, testCase := range testCaseGroup {
				err := validate.PortRegistered(testCase)
				if err == nil {
					t.Errorf("expected error for invalid input %d, but got no error", testCase)
				}
			}
		})
	}
}

func TestPortPrivateSuccessfulForValidInput(t *testing.T) {
	testCaseGroups := map[string][]uint16{
		"EdgeCases": {
			49152,
			65535,
		},
	}

	for name, testCaseGroup := range testCaseGroups {
		t.Run(name, func(t *testing.T) {
			for _, testCase := range testCaseGroup {
				err := validate.PortPrivate(testCase)
				if err != nil {
					t.Errorf("expected no error for valid input %d, but got error: %v", testCase, err)
				}
			}
		})
	}
}

func TestPortPrivateFailsForInvalidInput(t *testing.T) {
	testCaseGroups := map[string][]uint16{
		"PortsBelowRange": {
			0,
			1023,
			49151,
		},
		"CommonInvalidPorts": {
			443,
			8443,
		},
	}
	for name, testCaseGroup := range testCaseGroups {
		t.Run(name, func(t *testing.T) {
			for _, testCase := range testCaseGroup {
				err := validate.PortPrivate(testCase)
				if err == nil {
					t.Errorf("expected error for invalid input %d, but got no error", testCase)
				}
			}
		})
	}
}
