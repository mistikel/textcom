package handler

import (
	"log"
	"textcom/file"
	"textcom/helper"
)

type Handler struct {
	f file.IFile
}

func NewHandler(f file.IFile) *Handler {
	return &Handler{
		f: f,
	}
}

func (h *Handler) Run(text string) error {
	h.f.Add(helper.UpperCase(text))
	h.f.Add(helper.CamelCase(text))
	return h.write()
}

func (h *Handler) write() error {
	_, err := h.f.Write()
	if err != nil {
		log.Fatal(err)
	}

	return err
}
