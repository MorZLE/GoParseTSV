package service

import (
	"errors"
	"github.com/MorZLE/GoParseTSV/constants"
	"github.com/MorZLE/GoParseTSV/internal/model"
	"github.com/MorZLE/GoParseTSV/internal/repository/mocks"
	"reflect"
	"testing"
)

func Test_serviceImpl_GetAllGuid(t *testing.T) {

	type mckS func(r *mocks.Repository)

	type args struct {
		req model.RequestGetGuid
		m   mckS
	}
	tests := []struct {
		name    string
		args    args
		want    [][]model.Guid
		wantErr error
	}{
		{name: "PosiveTest1",
			args: args{
				req: model.RequestGetGuid{
					Limite:   2,
					Page:     1,
					UnitGUID: "unitguid",
				},
				m: func(r *mocks.Repository) {
					r.On("Get", "unitguid").Return([]model.Guid{
						{Number: "1"}, {Number: "2"}, {Number: "3"},
						{Number: "4"}, {Number: "5"},
					}, nil)
				},
			},
			want: [][]model.Guid{
				{model.Guid{Number: "2"}, model.Guid{Number: "3"}},
				{model.Guid{Number: "4"}, model.Guid{Number: "5"}},
			},
			wantErr: nil,
		},
		{name: "PosiveTest2",
			args: args{
				req: model.RequestGetGuid{
					Limite:   4,
					Page:     1,
					UnitGUID: "unitguid",
				},
				m: func(r *mocks.Repository) {
					r.On("Get", "unitguid").Return([]model.Guid{
						{Number: "1"}, {Number: "2"}, {Number: "3"},
						{Number: "4"}, {Number: "5"},
						{Number: "6"}, {Number: "7"}, {Number: "8"},
						{Number: "9"}, {Number: "10"},
					}, nil)
				},
			},
			want: [][]model.Guid{
				{model.Guid{Number: "2"}, model.Guid{Number: "3"},
					model.Guid{Number: "4"}, model.Guid{Number: "5"}},
				{model.Guid{Number: "6"}, model.Guid{Number: "7"},
					model.Guid{Number: "8"}, model.Guid{Number: "9"}},
				{model.Guid{Number: "10"}},
			},
			wantErr: nil,
		},
		{name: "NegativeTest1",
			args: args{
				req: model.RequestGetGuid{
					Limite:   0,
					Page:     1,
					UnitGUID: "unitguid",
				},
				m: func(r *mocks.Repository) {},
			},
			wantErr: constants.ErrEnabledData,
		},
		{name: "NegativeTest2",
			args: args{
				req: model.RequestGetGuid{
					Limite:   2,
					Page:     1,
					UnitGUID: "",
				},
				m: func(r *mocks.Repository) {},
			},
			wantErr: constants.ErrEnabledData,
		},
		{name: "NegativeTest3",
			args: args{
				req: model.RequestGetGuid{
					Limite:   2,
					Page:     -1,
					UnitGUID: "",
				},
				m: func(r *mocks.Repository) {},
			},
			wantErr: constants.ErrEnabledData,
		},
		{name: "NegativeTest4",
			args: args{
				req: model.RequestGetGuid{
					Limite:   -2,
					Page:     1,
					UnitGUID: "",
				},
				m: func(r *mocks.Repository) {},
			},
			wantErr: constants.ErrEnabledData,
		},
		{name: "ErrDBTest5",
			args: args{
				req: model.RequestGetGuid{
					Limite:   2,
					Page:     1,
					UnitGUID: "afawf",
				},
				m: func(r *mocks.Repository) {
					r.On("Get", "afawf").Return([]model.Guid{}, constants.ErrNotFound)
				},
			},
			wantErr: constants.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := mocks.NewRepository(t)
			tt.args.m(rep)
			s := &serviceImpl{
				r: rep,
			}
			got, err := s.GetAllGuid(tt.args.req)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetAllGuid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllGuid() got = %v, want %v", got, tt.want)
			}
		})
	}
}
