package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.expect) {
				t.Fatalf("\nwant: %v \ngot:  %v", tt.expect, tt.nums)
			}
		})
	}
}
