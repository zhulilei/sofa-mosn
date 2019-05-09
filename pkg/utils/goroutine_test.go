/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestGoWithRecover(t *testing.T) {
	panicStr := "test panic"
	panicHandler := func() {
		panic(panicStr)
	}
	output := ""
	recoverHandler := func(r interface{}) {
		output = fmt.Sprintf("%v", r)
	}
	GoWithRecover(panicHandler, recoverHandler, false)
	time.Sleep(time.Second) // wait panic goroutine
	if output != panicStr {
		t.Errorf("expected catch panic output, but got: %s", output)
	}
}

func TestGoWithRecoverRestart(t *testing.T) {
	count := 0
	noPanic := false
	panicHandler := func() {
		count++
		if count <= 2 {
			panic("panic")
		}
		noPanic = true
	}
	GoWithRecover(panicHandler, nil, true)
	time.Sleep(time.Second)
	if !(noPanic && count == 3) {
		t.Errorf("panic handler is not restart expectedly")
	}
}