package validate_test

import (
	"testing"

	"github.com/copartner6412/input/validate"
)

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
