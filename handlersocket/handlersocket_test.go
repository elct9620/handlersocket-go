// Copyright 2010  The "handlersocket-go" Authors
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

package handlersocket

import (
	"testing"
	"fmt"
	//	"syscall"
)


func TestOpenIndex(t *testing.T) {
	target := HandlerSocketTarget{host: "127.0.0.1", port: 9999, index: 1, database: "hstest", table: "hstest_table1", indexname: "PRIMARY", columns: []string{"k", "v"}}

	if c := NewHandlerSocketConnection(target); c != nil {
		defer c.Close()
		c.OpenIndex(target)
		if c.lastError.Code != "0" {
			t.Errorf("Last Error Code = %s, want %s.", c.lastError.Code, "0")
		}
	}

}
/*
func TestWrite(t *testing.T) {
	target := HandlerSocketTarget{host: "127.0.0.1", port: 9999, index: 1, database: "hstest", table: "hstest_table1", indexname: "PRIMARY", columns: []string{"k", "v"}}

	if c := NewHandlerSocketConnection(target); c != nil {
		defer c.Close()
		c.OpenIndex(target)
		fmt.Println(c.lastError)
		if c.lastError.Code != "0" {
			t.Errorf("Last Error Code = %s, want %s.", c.lastError.Code, "0")
		}
	}
}
*/

func TestRead(t *testing.T) {
	target := HandlerSocketTarget{host: "127.0.0.1", port: 9999, index: 1, database: "hstest", table: "hstest_table1", indexname: "PRIMARY", columns: []string{"k", "v"}}

	if c := NewHandlerSocketConnection(target); c != nil {
		defer c.Close()
		c.OpenIndex(target)
		if c.lastError.Code != "0" {
			t.Errorf("Last Error Code = %s, want %s.", c.lastError.Code, "0")
		}
		resp := c.Find("1", "=", "1", "0", "blue")
		fmt.Println(resp)

	}

}

func BenchmarkRead(b *testing.B) {
	target := HandlerSocketTarget{host: "127.0.0.1", port: 9999, index: 1, database: "hstest", table: "hstest_table1", indexname: "PRIMARY", columns: []string{"k", "v"}}
	if c := NewHandlerSocketConnection(target); c != nil {
		defer c.Close()
		c.OpenIndex(target)

		for i := 0; i < b.N; i++ {

			if c.lastError.Code != "0" {
				fmt.Println("Last Error Code = %s, want %s.", c.lastError.Code, "0")
			}
			c.Find("1", "=", "1", "0", "blue")

		}
	}
}
