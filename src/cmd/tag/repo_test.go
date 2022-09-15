package tag

import (
	"context"
	"reflect"
	"testing"

	httperrors "github.com/myrachanto/erroring"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_tagrepository_getuno(t *testing.T) {
	type fields struct {
		Mongodb *mongo.Database
		Bizname string
		Ctx     context.Context
		Cancel  context.CancelFunc
	}
	type args struct {
		code string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult *Tag
		wantErr    httperrors.HttpErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &tagrepository{
				Mongodb: tt.fields.Mongodb,
				Bizname: tt.fields.Bizname,
				Cancel:  tt.fields.Cancel,
			}
			gotResult, gotErr := r.getuno(tt.args.code)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("tagrepository.getuno() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("tagrepository.getuno() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
