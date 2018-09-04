package backend

import (
	"reflect"
	"testing"

	"github.com/solcates/gobwa/pkg/bwa/messages"
)

func TestNewSpaFinder(t *testing.T) {
	tests := []struct {
		name string
		want *SpaFinder
	}{
		{
			name: "Default",
			want: &SpaFinder{
				spas: make(map[string]*Spa),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpaFinder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpaFinder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpaFinder_Find(t *testing.T) {
	sp := NewSpaFinder()
	tests := []struct {
		name    string
		sp      *SpaFinder
		wantErr bool
	}{
		{
			name:    "ok",
			sp:      sp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sp.Find(); (err != nil) != tt.wantErr {
				t.Errorf("SpaFinder.Find() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpaFinder_Status(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name       string
		sp         *SpaFinder
		args       args
		wantStatus *messages.Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStatus := tt.sp.Status(tt.args.address); !reflect.DeepEqual(gotStatus, tt.wantStatus) {
				t.Errorf("SpaFinder.Status() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
