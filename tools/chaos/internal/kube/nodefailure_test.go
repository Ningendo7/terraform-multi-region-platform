package kube

import (
	"slices"
	"testing"
)

func TestSelectByCount(t *testing.T) {
	candidates := []string{"a", "b", "c", "d", "e"}

	cases := []struct {
		name  string
		count int
		want  int
	}{
		{name: "count zero returns all", count: 0, want: 5},
		{name: "count negative returns all", count: -1, want: 5},
		{name: "count exceeds candidates returns all", count: 10, want: 5},
		{name: "count within range samples down", count: 2, want: 2},
		{name: "count equals length returns all", count: 5, want: 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := selectByCount(candidates, c.count, 42)

			if len(got) != c.want {
				t.Fatalf("selectByCount(%v, %d, 42) = %v, want length %d", candidates, c.count, got, c.want)
			}

			for _, name := range got {
				if !slices.Contains(candidates, name) {
					t.Fatalf("selectByCount returned %q, not in original candidates %v", name, candidates)
				}
			}
		})
	}
}

func TestSelectByCountDeterministic(t *testing.T) {
	candidates := []string{"a", "b", "c", "d", "e"}

	first := selectByCount(candidates, 2, 42)
	second := selectByCount(candidates, 2, 42)

	if !slices.Equal(first, second) {
		t.Fatalf("selectByCount with the same seed returned different results: %v vs %v", first, second)
	}
}

func TestCheckBlastRadius(t *testing.T) {
	cases := []struct {
		name        string
		targetCount int
		clusterSize int
		force       bool
		wantErr     bool
	}{
		{name: "under cap", targetCount: 1, clusterSize: 4},
		{name: "exactly at cap", targetCount: 2, clusterSize: 4},
		{name: "over cap", targetCount: 3, clusterSize: 4, wantErr: true},
		{name: "over cap but forced", targetCount: 3, clusterSize: 4, force: true},
		{name: "targeting everything", targetCount: 4, clusterSize: 4, wantErr: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := checkBlastRadius(c.targetCount, c.clusterSize, c.force)

			if c.wantErr && err == nil {
				t.Fatalf("checkBlastRadius(%d, %d, force=%v) = nil, want error", c.targetCount, c.clusterSize, c.force)
			}
			if !c.wantErr && err != nil {
				t.Fatalf("checkBlastRadius(%d, %d, force=%v) unexpected error: %v", c.targetCount, c.clusterSize, c.force, err)
			}
		})
	}
}