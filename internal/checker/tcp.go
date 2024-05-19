package checker

import (
	"context"
	"fmt"
	"go-tcp-checker/internal/config"
	"go-tcp-checker/internal/results"
	"net"
	"sync"
	"time"
)

// ServiceChecker defines the behavior for checking network services.
type ServiceChecker interface {
	CheckService(ctx context.Context, address string) (results.ServiceResult, error)
}

// TCPServiceChecker checks the availability of TCP services.
type TCPServiceChecker struct{}

// CheckService tries to establish a TCP connection to the specified address.
func (t *TCPServiceChecker) CheckService(ctx context.Context, address string) (results.ServiceResult, error) {
	start := time.Now() // Start time measurement
	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", address)
	if err != nil {
		return results.ServiceResult{Address: address, Available: false, Duration: time.Since(start)}, &ServiceCheckError{Address: address, Reason: err.Error()}
	}
	err = conn.Close()
	if err != nil {
		return results.ServiceResult{Address: address, Available: false, Duration: time.Since(start)}, &ServiceCheckError{Address: address, Reason: err.Error()}
	}
	return results.ServiceResult{Address: address, Available: true, Duration: time.Since(start)}, nil
}

// ServiceCheckError is a custom error type that provides more detailed information about a service check failure.
type ServiceCheckError struct {
	Address string
	Reason  string
}

func (e *ServiceCheckError) Error() string {
	return fmt.Sprintf("Service unavailable at %s: %s", e.Address, e.Reason)
}

// CheckServicesConcurrentlyWithContext checks the availability of multiple network services concurrently.
func CheckServicesConcurrentlyWithContext(checker ServiceChecker, cfg *config.Config) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	rst := make([]results.ServiceResult, len(cfg.Services))
	semaphore := make(chan struct{}, cfg.ConcurrencyLevel)

	for i, service := range cfg.Services {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire a token

		go func(i int, addr string) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release the token

			result, err := checker.CheckService(ctx, addr)

			if err != nil {
				result.Error = err.Error()
			}
			rst[i] = result
		}(i, service.Address)
	}

	wg.Wait()
	results.PrintResults(rst, results.OutputFormat(cfg.OutputFormat), cfg.OutputPath)
}
