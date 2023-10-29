package handler_test

import (
	"testing"
	"textcom/handler"
	mock_file "textcom/mock/file"

	"github.com/golang/mock/gomock"
)

func TestHandler(t *testing.T) {
	f := mock_file.NewMockIFile(gomock.NewController(t))
	h := handler.NewHandler(f)

	f.EXPECT().Add("MOCK")
	f.EXPECT().Add("mOcK")
	f.EXPECT().Write().Return("PATH", nil)
	err := h.Run("mock")
	if err != nil {
		t.Errorf("Want: nil got [%v]", err)
	}
}
