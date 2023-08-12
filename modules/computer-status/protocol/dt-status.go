/* For license and copyright information please see the LEGAL file in the code repository */

package csp

type Status uint8

const (
	Status_Unset Status = iota

	Status_Ready
	Status_Shutdown

	Status_DiskPressure
	Status_MemoryPressure
	Status_ProcessPressure
	Status_GPUPressure
	Status_TPUPressure
	Status_NetworkUnavailable
	// TODO::: add more status
)
