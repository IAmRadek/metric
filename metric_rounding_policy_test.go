package metric_test

import (
	"testing"

	"github.com/IAmRadek/metric"
	isser "github.com/matryer/is"
)

func TestRoundingPolicy(t *testing.T) {
	tests := []struct {
		name   string
		policy metric.RoundingPolicy
		check  func(is *isser.I, policy metric.RoundingPolicy)
	}{
		{
			name:   "roundup_positive",
			policy: metric.RoundUp(1),
			check: func(is *isser.I, policy metric.RoundingPolicy) {
				is.Equal(policy.Name(), "ROUND_UP")
				is.Equal(policy.Round(4.45), 4.5)
			},
		},
		{
			name:   "roundup_negative",
			policy: metric.RoundUp(1),
			check: func(is *isser.I, policy metric.RoundingPolicy) {
				is.Equal(policy.Name(), "ROUND_UP")
				is.Equal(policy.Round(-4.45), -4.5)
			},
		},
		{
			name:   "roundup_positive_2",
			policy: metric.RoundUp(2),
			check: func(is *isser.I, policy metric.RoundingPolicy) {
				is.Equal(policy.Name(), "ROUND_UP")
				is.Equal(policy.Round(4.445), 4.45)
			},
		},
		{
			name:   "roundup_negative_2",
			policy: metric.RoundUp(2),
			check: func(is *isser.I, policy metric.RoundingPolicy) {
				is.Equal(policy.Name(), "ROUND_UP")
				is.Equal(policy.Round(-4.445), -4.45)
			},
		},
		{
			name:   "rounddown_positive",
			policy: metric.RoundDown(1),
			check: func(is *isser.I, policy metric.RoundingPolicy) {
				is.Equal(policy.Name(), "ROUND_DOWN")
				is.Equal(policy.Round(4.45), 4.4)
			},
		},
		{
			name:   "rounddown_negative",
			policy: metric.RoundDown(1),
			check: func(is *isser.I, policy metric.RoundingPolicy) {
				is.Equal(policy.Name(), "ROUND_DOWN")
				is.Equal(policy.Round(-4.45), -4.4)
			},
		},
		{
			name:   "round_positive",
			policy: metric.Round(1, 5),
			check: func(is *isser.I, policy metric.RoundingPolicy) {
				is.Equal(policy.Name(), "ROUND")
				is.Equal(policy.Round(4.45), 4.5)
			},
		},
		{
			name:   "round_positive_2",
			policy: metric.Round(1, 6),
			check: func(is *isser.I, policy metric.RoundingPolicy) {
				is.Equal(policy.Name(), "ROUND")
				is.Equal(policy.Round(4.45), 4.4)
			},
		},
		{
			name:   "round_negative",
			policy: metric.Round(1, 5),
			check: func(is *isser.I, policy metric.RoundingPolicy) {
				is.Equal(policy.Name(), "ROUND")
				is.Equal(policy.Round(-4.45), -4.5)
			},
		},
		{
			name:   "round_negative_2",
			policy: metric.Round(1, 6),
			check: func(is *isser.I, policy metric.RoundingPolicy) {
				is.Equal(policy.Name(), "ROUND")
				is.Equal(policy.Round(-4.45), -4.4)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is := isser.New(t)
			tt.check(is, tt.policy)
		})
	}
}
