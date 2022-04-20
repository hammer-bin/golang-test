package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// An automated process to update secret values periodically
var (
	/*ttl, _ = strconv.Atoi(os.Getenv("VAULT_TTL")) // IN MINUTES

	leasePath                  = os.Getenv("EXAMPLE_VAULT_LEASE")*/
	ttl                        = 5
	leasePath                  = "aaa/data/oracleDB"
	leaseSecretPattern         = []string{"data", "data"}
	leaseSecretInjectionPoints = map[string]string{"data": "db_name", "password": "password"}

	step   = time.Duration(ttl) * time.Minute
	ticker = time.NewTicker(step)
	done   = make(chan bool)
)

func UpdateSecretValues() {
	// Update the values immediately, and then go into a ticker pattern
	updateValues()

	// Ticker
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println(t)
			updateValues()
		}
	}
}

// Simple high-level wrapper script
func updateValues() {
	// Run the injector at least once
	err := injectWrapper(leasePath, leaseSecretPattern, leaseSecretInjectionPoints)

	if err != nil {
		return
	}
}

// Injection wrapper
func injectWrapper(leasePath string, secretPattern []string, injectorPoints map[string]string) error {
	if secretPattern[0] == "" || secretPattern[1] == "" {
		return errors.New("secret pattern UNDEFINED")
	}

	secretValues, err := ReadSecretValues(leasePath, secretPattern[0], secretPattern[1])

	// Fail-safes 1/2
	if err != nil {
		return err
	}

	// Fail-safes 2/2
	if secretValues == nil {
		return errors.New("no secret values in path")
	}

	// Inject the values
	injErr := secretValuesInjector(secretValues, injectorPoints)

	if injErr != nil {
		return err
	}

	return nil
}

// Takes an arbitrary number of values and injection points and sets env vars values
func secretValuesInjector(values map[string]interface{}, injectionPoints map[string]string) error {
	if len(values) != len(injectionPoints) {
		return errors.New("length mismatch, aborting")
	}

	for i, envKey := range injectionPoints {
		err := os.Setenv(envKey, fmt.Sprintf("%v", values[i]))

		if err != nil {
			break
		}
	}

	return nil
}
