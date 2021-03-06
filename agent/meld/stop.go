package meld

import (
	"github.com/factionlabs/meld-agent/pkg/rust"
)

func (m *Meld) StopRust(args int, success *bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	r, err := rust.NewRustServer(nil)
	if err != nil {
		*success = false
		return err
	}

	if err := r.Stop(); err != nil {
		*success = false
		return err
	}

	*success = true
	return nil
}
