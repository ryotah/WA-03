package roman

import "testing"

func TestToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				num: 1,
			},
			want: "I",
		},
		{
			name: "3999",
			args: args{
				num: 3999,
			},
			want: "MMMCMXCIX",
		},
		{
			name: "0",
			args: args{
				num: 0,
			},
			wantErr: true,
		},
		{
			name: "4000",
			args: args{
				num: 0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToRoman(tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToRoman() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
