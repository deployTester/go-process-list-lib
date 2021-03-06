// +build dragonfly

package ps

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Path returns path process executable
func (p *UnixProcess) Path() (string, error) {
         return filepath.EvalSymlinks(fmt.Sprintf("/proc/%d/file", p.pid))
}

// Refresh reloads data associated with this process
func (p *UnixProcess) Refresh() error {
	statPath := fmt.Sprintf("/proc/%d/status", p.pid)
	dataBytes, err := ioutil.ReadFile(statPath)
	if err != nil {
		return err
	}

	data := string(dataBytes)

	_, err = fmt.Sscanf(data,
		"%s %d %d %d %d",
		&p.binary,
		&p.pid,
		&p.ppid,
		&p.pgrp,
		&p.sid)

	return err
}