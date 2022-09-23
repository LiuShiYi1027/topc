package cgroups

import (
	"testing"
)

func TestParseFileKeyValue(t *testing.T) {
	type args struct {
		dir  string
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "open mem stat",
			args: args{
				dir:  "/sys/fs/cgroup/memory/",
				file: "memory.stat",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFileKeyValue(tt.args.dir, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseKeyValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("ParseKeyValue() got = %v, want got length is not zero", got)
			}
		})
	}
}
