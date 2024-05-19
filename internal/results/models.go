package results

import "time"

// ServiceResult represents the result of checking a single service.
type ServiceResult struct {
	Address   string        `json:"address"`
	Available bool          `json:"available"`
	Error     string        `json:"error,omitempty"` // Include the error message if present
	Duration  time.Duration `json:"duration"`
}
