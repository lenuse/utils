package utils

import "testing"

func TestToUnderScore(t *testing.T) {
	str := "XyID"
	target := "xy_id"
	if actually := ToUnderScore(str); actually != target {
		t.Errorf("ToUnderScore: [%v], actually: [%v]", target, actually)
	}
}

func BenchmarkToUnderScore(b *testing.B) {
	str := "CmsSubject"
	for i := 0; i < b.N; i++  {
		ToUnderScore(str)
	}
}

func BenchmarkToUnderScoreParallel(b *testing.B) {
	str := "ZcXyId"
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ToUnderScore(str)
		}
	})
}

func TestToCamelCase(t *testing.T) {
	str := "xy_id"
	target := "XyId"
	if actually := ToCamelCase(str); actually != target {
		t.Errorf("ToUnderScore: [%v], actually: [%v]", target, actually)
	}
}

func BenchmarkToCamelCase(b *testing.B) {
	str := "xy_id"
	for i := 0; i < b.N; i++  {
		ToCamelCase(str)
	}
}