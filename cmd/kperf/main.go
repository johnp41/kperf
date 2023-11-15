// Copyright 2020 The Knative Authors
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

package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"

	"knative.dev/kperf/core"
)

func main() {
	var mask uintptr

	// Get the current CPU affinity of the process
	if _, _, err := syscall.RawSyscall(syscall.SYS_SCHED_GETAFFINITY, 0, uintptr(unsafe.Sizeof(mask)), uintptr(unsafe.Pointer(&mask))); err != 0 {
		fmt.Println("Failed to get CPU affinity:", err)
		return
	}
	fmt.Println("Current CPU affinity:", mask)

	// Set the new CPU affinity
	mask = 1
	if _, _, err := syscall.RawSyscall(syscall.SYS_SCHED_SETAFFINITY, 0, uintptr(unsafe.Sizeof(mask)), uintptr(unsafe.Pointer(&mask))); err != 0 {
		fmt.Println("Failed to set CPU affinity:", err)
		return
	}

	if err := core.NewPerfCommand().Execute(); err != nil {
		log.Println("failed to execute kperf command:", err)
		os.Exit(1)
	}
}
