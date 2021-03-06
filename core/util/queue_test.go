/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Push("first")
	q.Push("second")
	q.Push("third")
	t.Logf("Size: %d\n", q.Size())
	head := q.Peek()
	if head != "first" {
		t.Fatalf("head is not correct. Expected \"first\", got %v", head)
	}
	t.Logf("Pop: %v\n", q.Pop())
	if q.Size() != 2 {
		t.Fatalf("Wrong queue size != 2: %d", q.Size())
	}
	t.Logf("Pop: %v\n", q.Pop())
	t.Logf("Pop: %v\n", q.Pop())
	if q.Size() != 0 {
		t.Fatalf("Wrong queue size != 0: %d", q.Size())
	}

}
