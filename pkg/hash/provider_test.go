package hash

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewHashProvider(t *testing.T) {
	tests := []struct {
		name    string
		want    *Provider
		wantErr bool
	}{
		{
			name:    "Create Hash Provider",
			want:    &Provider{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHashProvider()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHashProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashProvider() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_CheckPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CheckPasswordHash",
			args: args{
				"123456",
				"$2a$14$w9bLCQseCdvXfqoXJ3Dn2.Pg7va5RxBHZUOW8R6NkZDEPtpFQxXs6",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{}
			got := p.CheckPasswordHash(tt.args.password, tt.args.hash)
			assert.Equal(t, tt.want, got, "CheckPasswordHash() = %v, want %v", got, tt.want)
		})
	}
}

func TestProvider_HashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "HashPassword",
			args: args{
				"123456",
			},
			want:    "$2a$14$w9bLCQseCdvXfqoXJ3Dn2.Pg7va5RxBHZUOW8R6NkZDEPtpFQxXs6",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{}
			got, err := p.HashPassword(tt.args.password)
			check := p.CheckPasswordHash(tt.args.password, got)
			assert.True(t, check, "HashPassword() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}
