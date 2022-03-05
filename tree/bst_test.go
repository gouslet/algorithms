/*
 * File: \tree\bst_test.go                                                     *
 * Project: algorithms                                                         *
 * Created At: Saturday, 2022/03/5 , 14:21:45                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/03/5 , 22:21:03                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package tree

import (
	"testing"
)

func TestPutAndGet(t *testing.T) {

	tests := []element[int]{
		{1, "a"},
		{2, "b"},
		{3, "c"},
		{4, "d"},
		{5, "e"},
	}

	bst := NewBST(func(a, b int) int {
		return a - b
	})

	for _, test := range tests {
		bst.Put(test.key, test.val)
		t.Logf("Put(%d, %s)", test.key, test.val)

	}

	for _, test := range tests {
		if v := bst.Get(test.key); v != test.val {
			t.Errorf("Get(%d) = %s,want %s", test.key, v, test.val)
		}
	}

}
