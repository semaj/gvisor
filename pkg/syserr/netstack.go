// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syserr

import (
	"fmt"

	"gvisor.dev/gvisor/pkg/abi/linux"
	"gvisor.dev/gvisor/pkg/sync"
	"gvisor.dev/gvisor/pkg/tcpip"
)

// Mapping for tcpip.Error types.
var (
	ErrUnknownProtocol       = New(tcpip.ErrUnknownProtocol.String(), linux.EINVAL)
	ErrUnknownNICID          = New(tcpip.ErrUnknownNICID.String(), linux.ENODEV)
	ErrUnknownDevice         = New(tcpip.ErrUnknownDevice.String(), linux.ENODEV)
	ErrUnknownProtocolOption = New(tcpip.ErrUnknownProtocolOption.String(), linux.ENOPROTOOPT)
	ErrDuplicateNICID        = New(tcpip.ErrDuplicateNICID.String(), linux.EEXIST)
	ErrDuplicateAddress      = New(tcpip.ErrDuplicateAddress.String(), linux.EEXIST)
	ErrBadLinkEndpoint       = New(tcpip.ErrBadLinkEndpoint.String(), linux.EINVAL)
	ErrAlreadyBound          = New(tcpip.ErrAlreadyBound.String(), linux.EINVAL)
	ErrInvalidEndpointState  = New(tcpip.ErrInvalidEndpointState.String(), linux.EINVAL)
	ErrAlreadyConnecting     = New(tcpip.ErrAlreadyConnecting.String(), linux.EALREADY)
	ErrNoPortAvailable       = New(tcpip.ErrNoPortAvailable.String(), linux.EAGAIN)
	ErrPortInUse             = New(tcpip.ErrPortInUse.String(), linux.EADDRINUSE)
	ErrBadLocalAddress       = New(tcpip.ErrBadLocalAddress.String(), linux.EADDRNOTAVAIL)
	ErrClosedForSend         = New(tcpip.ErrClosedForSend.String(), linux.EPIPE)
	ErrClosedForReceive      = New(tcpip.ErrClosedForReceive.String(), nil)
	ErrTimeout               = New(tcpip.ErrTimeout.String(), linux.ETIMEDOUT)
	ErrAborted               = New(tcpip.ErrAborted.String(), linux.EPIPE)
	ErrConnectStarted        = New(tcpip.ErrConnectStarted.String(), linux.EINPROGRESS)
	ErrDestinationRequired   = New(tcpip.ErrDestinationRequired.String(), linux.EDESTADDRREQ)
	ErrNotSupported          = New(tcpip.ErrNotSupported.String(), linux.EOPNOTSUPP)
	ErrQueueSizeNotSupported = New(tcpip.ErrQueueSizeNotSupported.String(), linux.ENOTTY)
	ErrNoSuchFile            = New(tcpip.ErrNoSuchFile.String(), linux.ENOENT)
	ErrInvalidOptionValue    = New(tcpip.ErrInvalidOptionValue.String(), linux.EINVAL)
	ErrBroadcastDisabled     = New(tcpip.ErrBroadcastDisabled.String(), linux.EACCES)
	ErrNotPermittedNet       = New(tcpip.ErrNotPermitted.String(), linux.EPERM)
)

var initOnce sync.Once
var errTranslationMu sync.RWMutex
var netstackErrorTranslations map[string]*Error

// Precondition: calls to this function are synchronized.
func addErrMapping(key string, val *Error) {
	if _, ok := netstackErrorTranslations[key]; ok {
		panic(fmt.Sprintf("duplicate error key: %s", key))
	}
	netstackErrorTranslations[key] = val
}

// Precondition: calls to this function are synchronized.
func initErrorTranslations() {
	netstackErrorTranslations = make(map[string]*Error)
	addErrMapping(tcpip.ErrUnknownProtocol.String(), ErrUnknownProtocol)
	addErrMapping(tcpip.ErrUnknownNICID.String(), ErrUnknownNICID)
	addErrMapping(tcpip.ErrUnknownDevice.String(), ErrUnknownDevice)
	addErrMapping(tcpip.ErrUnknownProtocolOption.String(), ErrUnknownProtocolOption)
	addErrMapping(tcpip.ErrDuplicateNICID.String(), ErrDuplicateNICID)
	addErrMapping(tcpip.ErrDuplicateAddress.String(), ErrDuplicateAddress)
	addErrMapping(tcpip.ErrNoRoute.String(), ErrNoRoute)
	addErrMapping(tcpip.ErrBadLinkEndpoint.String(), ErrBadLinkEndpoint)
	addErrMapping(tcpip.ErrAlreadyBound.String(), ErrAlreadyBound)
	addErrMapping(tcpip.ErrInvalidEndpointState.String(), ErrInvalidEndpointState)
	addErrMapping(tcpip.ErrAlreadyConnecting.String(), ErrAlreadyConnecting)
	addErrMapping(tcpip.ErrAlreadyConnected.String(), ErrAlreadyConnected)
	addErrMapping(tcpip.ErrNoPortAvailable.String(), ErrNoPortAvailable)
	addErrMapping(tcpip.ErrPortInUse.String(), ErrPortInUse)
	addErrMapping(tcpip.ErrBadLocalAddress.String(), ErrBadLocalAddress)
	addErrMapping(tcpip.ErrClosedForSend.String(), ErrClosedForSend)
	addErrMapping(tcpip.ErrClosedForReceive.String(), ErrClosedForReceive)
	addErrMapping(tcpip.ErrWouldBlock.String(), ErrWouldBlock)
	addErrMapping(tcpip.ErrConnectionRefused.String(), ErrConnectionRefused)
	addErrMapping(tcpip.ErrTimeout.String(), ErrTimeout)
	addErrMapping(tcpip.ErrAborted.String(), ErrAborted)
	addErrMapping(tcpip.ErrConnectStarted.String(), ErrConnectStarted)
	addErrMapping(tcpip.ErrDestinationRequired.String(), ErrDestinationRequired)
	addErrMapping(tcpip.ErrNotSupported.String(), ErrNotSupported)
	addErrMapping(tcpip.ErrQueueSizeNotSupported.String(), ErrQueueSizeNotSupported)
	addErrMapping(tcpip.ErrNotConnected.String(), ErrNotConnected)
	addErrMapping(tcpip.ErrConnectionReset.String(), ErrConnectionReset)
	addErrMapping(tcpip.ErrConnectionAborted.String(), ErrConnectionAborted)
	addErrMapping(tcpip.ErrNoSuchFile.String(), ErrNoSuchFile)
	addErrMapping(tcpip.ErrInvalidOptionValue.String(), ErrInvalidOptionValue)
	addErrMapping(tcpip.ErrNoLinkAddress.String(), ErrHostDown)
	addErrMapping(tcpip.ErrBadAddress.String(), ErrBadAddress)
	addErrMapping(tcpip.ErrNetworkUnreachable.String(), ErrNetworkUnreachable)
	addErrMapping(tcpip.ErrMessageTooLong.String(), ErrMessageTooLong)
	addErrMapping(tcpip.ErrNoBufferSpace.String(), ErrNoBufferSpace)
	addErrMapping(tcpip.ErrBroadcastDisabled.String(), ErrBroadcastDisabled)
	addErrMapping(tcpip.ErrNotPermitted.String(), ErrNotPermittedNet)
	addErrMapping(tcpip.ErrAddressFamilyNotSupported.String(), ErrAddressFamilyNotSupported)
}

// TranslateNetstackError converts an error from the tcpip package to a sentry
// internal error.
func TranslateNetstackError(err *tcpip.Error) *Error {
	if netstackErrorTranslations == nil {
		errTranslationMu.Lock()
		initOnce.Do(initErrorTranslations)
		errTranslationMu.Unlock()
	}
	if err == nil {
		return nil
	}
	errTranslationMu.RLock()
	se, ok := netstackErrorTranslations[err.String()]
	errTranslationMu.RUnlock()
	if !ok {
		panic("Unknown error: " + err.String())
	}
	return se
}
