package useCases

import (
	"fmt"
	"testing"
)

func TestFriendlyChains(t *testing.T) {
	const (
		X          string = "TOKYO"
		Y          string = "KYOTO"
		DONOTMATCH string = "yotok"
	)

	type args struct {
		x string
		y string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantErr      bool
	}{
		{
			name: "FriendlyChains they are friends",
			args: args{
				x: X,
				y: Y,
			},
			wantResponse: fmt.Sprintf(THEYAREFRIENDS, X, Y),
			wantErr:      false,
		},
		{
			name: "FriendlyChains they not are friends",
			args: args{
				x: X,
				y: DONOTMATCH,
			},
			wantResponse: fmt.Sprintf(THEYNOTAREFRIENDS, X, DONOTMATCH),
			wantErr:      false,
		},
		{
			name: "FriendlyChains error text strings must be same length",
			args: args{
				x: X,
				y: "CHI",
			},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := FriendlyChains(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendlyChains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("FriendlyChains() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
