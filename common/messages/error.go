// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Copyright (c) 2013-2015 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package messages

import (
	"fmt"
	"time"
)

// MessageError describes an issue with a message.
// An example of some potential issues are messages from the wrong bitcoin
// network, invalid commands, mismatched checksums, and exceeding max payloads.
//
// This provides a mechanism for the caller to type assert the error to
// differentiate between general io errors such as io.EOF and issues that
// resulted from malformed messages.
type MessageError struct {
	error
	Func        string // Function name
	Description string // Human readable description of the issue
}

var _ error = (*MessageError)(nil)

// Error satisfies the error interface and prints human-readable errors.
func (e *MessageError) Error() string {
	callTime := time.Now().UnixNano()
	defer messagesMessageErrorError.Observe(float64(time.Now().UnixNano() - callTime))
	if e.Func != "" {
		return fmt.Sprintf("%v: %v", e.Func, e.Description)
	}
	return e.Description
}

// messageError creates an error for the given function and description.
func messageError(f string, desc string) *MessageError {
	callTime := time.Now().UnixNano()
	defer messagesMessageErrormessageError.Observe(float64(time.Now().UnixNano() - callTime))
	return &MessageError{Func: f, Description: desc}
}
