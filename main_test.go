package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEnsureDomainRecordExist(t *testing.T) {
	debug = false
	type args struct {
		domain string
		rr     string
		ip     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				domain: "everonet.com",
				rr:     "drdemo",
				ip:     "47.242.55.241",
			},
			wantErr: false,
		},
		{
			args: args{
				domain: "everonet.com",
				rr:     "drdemo2",
				ip:     "47.242.55.241",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EnsureDomainRecordExist(tt.args.domain, tt.args.rr, tt.args.ip); (err != nil) != tt.wantErr {
				t.Errorf("EnsureDomainRecordExist() error = %v, wantErr %v", err, tt.wantErr)
			}
			record, err := GetDomainRecord(tt.args.domain, tt.args.rr)
			require.NoError(t, err)
			require.Equal(t, record.Value, tt.args.ip)
		})
	}
}
