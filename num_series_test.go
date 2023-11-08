package nummatch

import "testing"

func TestNumSeries_RoundDown(t *testing.T) {
	type fields struct {
		begin   int64
		offset  int64
		exclude []int64
	}
	tests := map[string]struct {
		fields  fields
		target  int64
		wantOut int64
	}{
		`target in exclude list`: {
			fields: fields{
				begin:   1,
				offset:  2,
				exclude: []int64{8, 14, 28},
			},
			target:  14,
			wantOut: 12,
		},
		`target in exclude list of values`: {
			fields: fields{
				begin:   0,
				offset:  2,
				exclude: []int64{2, 10, 14, 16, 18, 20, 28},
			},
			target:  18,
			wantOut: 12,
		},
		`target not in exclude list(above)`: {
			fields: fields{
				begin:   1,
				offset:  2,
				exclude: []int64{2, 10, 14, 16, 18, 20, 28},
			},
			target:  40,
			wantOut: 40,
		},
		`target not in exclude list (below)`: {
			fields: fields{
				begin:   8,
				offset:  2,
				exclude: []int64{10, 14, 16, 18, 20, 28},
			},
			target:  2,
			wantOut: 2,
		},
		`target in exclude begin`: {
			fields: fields{
				begin:   2,
				offset:  2,
				exclude: []int64{2, 4, 8, 14, 28},
			},
			target:  4,
			wantOut: 4,
		},
		`target with unneccassary exclude list`: {
			fields: fields{
				begin:   6,
				offset:  2,
				exclude: []int64{2, 4, 8, 14, 28},
			},
			target:  4,
			wantOut: 4,
		},
		`target with unneccassary exclude whole list`: {
			fields: fields{
				begin:   6,
				offset:  2,
				exclude: []int64{2, 4},
			},
			target:  4,
			wantOut: 4,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			n := NewNumSeries(
				tt.fields.begin,
				tt.fields.offset,
				func(i int64) bool {
					for _, val := range tt.fields.exclude {
						if i == val {
							return true
						}
					}
					return false
				},
			)
			if gotOut := n.RoundDown(tt.target); gotOut != tt.wantOut {
				t.Errorf("NewNumSeries.RoundDown() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestNumSeries_RoundUp(t *testing.T) {
	type fields struct {
		begin   int64
		offset  int64
		exclude []int64
	}
	tests := map[string]struct {
		fields  fields
		target  int64
		wantOut int64
	}{
		`target in exclude list`: {
			fields: fields{
				begin:   2,
				offset:  2,
				exclude: []int64{8, 14, 28},
			},
			target:  14,
			wantOut: 16,
		},
		`target not in the series`: {
			fields: fields{
				begin:   2,
				offset:  2,
				exclude: []int64{8, 14, 28},
			},
			target:  7,
			wantOut: 6,
		},
		`target in exclude list of values`: {
			fields: fields{
				begin:   0,
				offset:  2,
				exclude: []int64{2, 16, 18, 20, 28},
			},
			target:  18,
			wantOut: 22,
		},
		`target not in exclude list(above)`: {
			fields: fields{
				begin:   2,
				offset:  2,
				exclude: []int64{2, 10, 14, 16, 18, 20, 28},
			},
			target:  40,
			wantOut: 40,
		},
		`target not in exclude list (below)`: {
			fields: fields{
				begin:   8,
				offset:  2,
				exclude: []int64{10, 14, 16, 18, 20, 28},
			},
			target:  2,
			wantOut: 8,
		},
		`target in exclude begin`: {
			fields: fields{
				begin:   2,
				offset:  2,
				exclude: []int64{2, 4, 8, 14, 28},
			},
			target:  2,
			wantOut: 6,
		},
		`target with unneccassary exclude whole list`: {
			fields: fields{
				begin:   6,
				offset:  2,
				exclude: []int64{2, 4},
			},
			target:  4,
			wantOut: 6,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			n := NewNumSeries(
				tt.fields.begin,
				tt.fields.offset,
				func(i int64) bool {
					for _, val := range tt.fields.exclude {
						if i == val {
							return true
						}
					}
					return false
				},
			)
			if gotOut := n.RoundUp(tt.target); gotOut != tt.wantOut {
				t.Errorf("NewNumSeries.RoundDown() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
