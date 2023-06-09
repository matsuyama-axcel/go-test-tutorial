package chapter1

import (
	"testing"
)

func TestAddition(t *testing.T) {
	// 正常系のテストパターン
	success := map[string]struct {
		numA int
		numB int
		want int
	}{
		// FIXME: テストケースを追加
		"tc": {
			numA: 50,
			numB: 12,
			want: 62,
		},
	}
	// エラー系のテストパターン
	fail := map[string]struct {
		numA       int
		numB       int
		wantErrStr string
	}{    
		// FIXME: テストケースを追加
		"テストケース名": {  
			numA: ,
			numB: ,
			want: []string{"a","b","c"},
		},
	}

	for tt, tc := range success {
		t.Run(tt, func(t *testing.T) {
			got, err := addtion(tc.numA, tc.numB)
			if err != nil {
				t.Errorf("err is not nil: %s", err)
			}
			if tc.want != got {
				t.Errorf("unexpected return. want:%d actual:%d", tc.want, got)
			}
		})
	}
	for tt, tc := range fail {
		t.Run(tt, func(t *testing.T) {
			got, err := addtion(tc.numA, tc.numB)
			if got != 0 {
				t.Errorf("unexpected return. want:0 actual:%d", got)
			}
			if tc.wantErrStr != err.Error() {
				t.Errorf("unexpected err. want:%s actual:%s", tc.wantErrStr, err)
			}
		})
	}
}
