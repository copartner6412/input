package validate

import (
	"fmt"
)

// Constants representing port range limits.
const (
	limitPorts           uint   = 1 << 16   // 65,536 (total possible ports)
	limitPortsWellKnown  uint   = 1 << 10   // 1,024 (well-known ports range)
	limitPortsRegistered uint   = 49151 + 1 // 49,152 (registered ports range)
	minPortAllowed       uint16 = 0
	maxPortAllowed       uint16 = 65535
)

func Port(port, minPort, maxPort uint16) error {
	if minPort == 0 && maxPort == 0 {
		maxPort = maxPortAllowed
	} else {
		if minPort < maxPort {
			return fmt.Errorf("maximum port can not be less than minimum port")
		}
	}

	if port < minPort {
		return fmt.Errorf("port %d is less than lower limit of %d", port, minPort)
	}

	if port > maxPort {
		return fmt.Errorf("port %d exceeds upper limit of %d", port, maxPort)
	}

	return nil
}

// PortWellKnown checks if the port is in the well-known ports range (0-1023).
// It returns an error if the port is outside the well-known range.
func PortWellKnown(port uint16) error {
	if uint(port) >= limitPortsWellKnown {
		return fmt.Errorf("port %d outside of well-known ports range [0, 1,023]", port)
	}

	return nil
}

// PortNotWellKnown checks if the port is not in the well-known ports range (1024-65535).
// It returns an error if the port is a well-known port.
func PortNotWellKnown(port uint16) error {
	if uint(port) < limitPortsWellKnown {
		return fmt.Errorf("port %d is a well known port", port)
	}

	return nil
}

// PortRegistered checks if the port is in the registered ports range (1024-49151).
// It returns an error if the port is outside this range.
func PortRegistered(port uint16) error {
	if uint(port) < limitPortsWellKnown || limitPortsRegistered <= uint(port) {
		return fmt.Errorf("port %d outside of registered ports range [1,024, 49,151]", port)
	}

	return nil
}

// PortPrivate checks if the port is in the private ports range (49152-65535).
// It returns an error if the port is outside this range.
func PortPrivate(port uint16) error {
	if uint(port) < limitPortsRegistered {
		return fmt.Errorf("port %d outside of registered ports range [49,152, 65,535]", port)
	}

	return nil
}
