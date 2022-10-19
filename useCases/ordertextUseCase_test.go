package useCases

import (
	"reflect"
	"testing"
)

func TestOrderText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse []string
		wantErr      bool
	}{
		{
			name: "TestOrderText success",
			args: args{
				text: "Luis,Camilo,Andres,Laura",
			},
			wantResponse: []string{
				"Andres",
				"Camilo",
				"Laura",
				"Luis",
			},
			wantErr: false,
		},
		{
			name: "TestOrderText error ",
			args: args{
				text: "Luis,Camilo;Andres,Laura",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := OrderText(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("OrderText() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
