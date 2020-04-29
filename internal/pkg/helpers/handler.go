package helpers

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
	"reflect"
)

type BindRenderHandler struct {
	inputType  *reflect.Type
	outputType *reflect.Type
	fv         reflect.Value
}

func NewBindRenderHandler(f interface{}) *BindRenderHandler {
	fv := reflect.ValueOf(f)

	ft := fv.Type()
	if ft.Kind() != reflect.Func {
		panic("not a func")
	}

	h := BindRenderHandler{
		inputType:  nil,
		outputType: nil,
		fv:         fv,
	}

	if ft.NumIn() == 0 {
		h.inputType = nil
	} else if ft.NumIn() == 1 {
		h.inputType = new(reflect.Type)
		*h.inputType = ft.In(0).Elem()
	} else {
		panic("f must have 0 or 1 arguments")
	}

	if ft.NumOut() == 1 {
		h.outputType = nil
	} else if ft.NumOut() == 2 {
		h.outputType = new(reflect.Type)
		*h.outputType = ft.Out(0).Elem()
	} else {
		panic("f must return 1 or 2 values")
	}

	return &h
}

func (h *BindRenderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	args := []reflect.Value{}

	var err error
	if h.inputType != nil {
		input := reflect.New(*h.inputType)
		args = append(args, input)

		if binder, ok := input.Interface().(render.Binder); ok {
			err = render.Bind(r, binder)
		} else {
			err = render.Decode(r, input.Interface())
		}

		if err != nil {
			fmt.Println(err)
			// TODO
			return
		}
	}

	returned := h.fv.Call(args)

	var returnedErr interface{}
	if h.outputType != nil {
		returnedErr = returned[1].Interface()
	} else {
		returnedErr = returned[0].Interface()
	}

	if returnedErr != nil {
		fmt.Println(err)
		return
	}

	if h.outputType == nil {
		return
	}

	output := returned[0].Interface()
	if renderer, ok := output.(render.Renderer); ok {
		err = render.Render(w, r, renderer)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	render.Respond(w, r, output)
}
