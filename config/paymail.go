package config

import (
	"errors"

	"github.com/mrz1836/go-sanitize"
	"github.com/mrz1836/go-validate"
)

// Validate checks the configuration for specific rules
func (p *paymailConfig) Validate() error {

	// Only if enabled
	if p.Enabled {
		if len(p.Domains) == 0 {
			return errors.New("domain is required for paymail")
		}

		var err error
		for _, domain := range p.Domains {
			domain, err = sanitize.Domain(domain, false, true)
			if err != nil {
				return err
			}
			if !validate.IsValidHost(domain) {
				return errors.New("domain [" + domain + "] is not a valid hostname")
			}
		}

		// Todo: validate the default_from_paymail and default_note values
	}

	return nil
}