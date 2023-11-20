package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/MorZLE/GoParseTSV/constants"
	"github.com/MorZLE/GoParseTSV/internal/model"
	"github.com/MorZLE/GoParseTSV/internal/service/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetGuid(t *testing.T) {
	type mckS func(r *mocks.Service)
	req := func(t []byte) *http.Request {
		return httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8080/", bytes.NewBuffer(t))
	}

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
		m mckS
		t any
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantRes  [][]model.Guid
	}{
		{
			name: "positiveTest1",
			args: args{
				w: &httptest.ResponseRecorder{},
				m: func(r *mocks.Service) {
					r.On("GetAllGuid", model.RequestGetGuid{UnitGUID: "124124j243-f32r", Page: 2, Limite: 6}).Return([][]model.Guid{
						{model.Guid{Number: "1"}, model.Guid{Number: "2"}, model.Guid{Number: "3"}, model.Guid{Number: "4"}, model.Guid{Number: "5"}, model.Guid{Number: "5"}},
					}, nil)
				},
				t: model.RequestGetGuid{UnitGUID: "124124j243-f32r", Page: 2, Limite: 6},
			},
			wantRes: [][]model.Guid{
				{model.Guid{Number: "1"}, model.Guid{Number: "2"}, model.Guid{Number: "3"}, model.Guid{Number: "4"}, model.Guid{Number: "5"}, model.Guid{Number: "5"}},
			},
			wantCode: 200,
		},
		{
			name: "positiveTest2",
			args: args{
				w: &httptest.ResponseRecorder{},
				m: func(r *mocks.Service) {
					r.On("GetAllGuid", model.RequestGetGuid{UnitGUID: "124124j243-f32r", Page: 2, Limite: 2}).Return([][]model.Guid{
						{model.Guid{Number: "1"}, model.Guid{Number: "2"}},
					}, nil)
				},
				t: model.RequestGetGuid{UnitGUID: "124124j243-f32r", Page: 2, Limite: 2},
			},
			wantRes: [][]model.Guid{
				{model.Guid{Number: "1"}, model.Guid{Number: "2"}},
			},
			wantCode: 200,
		},
		{
			name: "testNotFound",
			args: args{
				w: &httptest.ResponseRecorder{},
				m: func(r *mocks.Service) {
					r.On("GetAllGuid", model.RequestGetGuid{UnitGUID: "124124j243-f32r", Page: 2, Limite: 2}).Return([][]model.Guid{}, constants.ErrNotFound)
				},
				t: model.RequestGetGuid{UnitGUID: "124124j243-f32r", Page: 2, Limite: 2},
			},
			wantRes:  nil,
			wantCode: 409,
		},
		{
			name: "negativeTest2",
			args: args{
				w: &httptest.ResponseRecorder{},
				m: func(r *mocks.Service) {
					r.On("GetAllGuid", model.RequestGetGuid{UnitGUID: "124124j243-f32r", Page: 2}).Return(nil, constants.ErrEnabledData)
				},
				t: model.RequestGetGuid{UnitGUID: "124124j243-f32r", Page: 2},
			},
			wantRes:  nil,
			wantCode: 400,
		},
		{
			name: "negativeTest3",
			args: args{
				w: &httptest.ResponseRecorder{},
				m: func(r *mocks.Service) {
					r.On("GetAllGuid", model.RequestGetGuid{UnitGUID: "124124j243-f32r", Page: 2}).Return(nil, errors.New("error"))
				},
				t: model.RequestGetGuid{UnitGUID: "124124j243-f32r", Page: 2},
			},
			wantRes:  nil,
			wantCode: 500,
		},
		{
			name: "negativeTest4",
			args: args{
				w: &httptest.ResponseRecorder{},
				m: func(r *mocks.Service) {
					r.On("GetAllGuid", model.RequestGetGuid{Page: 2}).Return(nil, constants.ErrEnabledData)
				},
				t: model.RequestGetGuid{Page: 2},
			},
			wantRes:  nil,
			wantCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()

			logic := mocks.NewService(t)
			tt.args.m(logic)

			body, _ := json.Marshal(&tt.args.t)
			tt.args.r = req(body)

			tt.args.r.Header.Set("Content-Type", "application/json")
			controller := &Handler{
				s: logic,
			}

			app.Post("/", controller.GetGuid)
			resp, err := app.Test(tt.args.r)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantCode, resp.StatusCode)
			if tt.wantRes == nil {
				return
			}
			wantbody, _ := json.Marshal(tt.wantRes)
			res, _ := io.ReadAll(resp.Body)
			assert.Equal(t, wantbody, res)
		})
	}
}
