package middleware

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"furina/errorCode"
)

type RespStruct struct {
	Code int
	Data interface{}
}

var (
	ErrMustValid         = errors.New("method must be valid")
	ErrMustFunc          = errors.New("method must be func")
	ErrMustPtr           = errors.New("param must be ptr")
	ErrMustPointToStruct = errors.New("param must point to struct")
	ErrMustHasThreeParam = errors.New("method must has three input")
	ErrMustOneOut        = errors.New("method must has one out")
	ErrMustRespErrorPtr  = errors.New("method ret must be *errorCode.RespError")

	RespErrorType = reflect.TypeOf((*errorCode.RespError)(nil))
)

func checkMethod(method interface{}) (mV reflect.Value, reqT, respT reflect.Type, err error) {
	mV = reflect.ValueOf(method)
	if !mV.IsValid() {
		err = ErrMustValid
		return
	}
	mT := mV.Type()
	if mT.Kind() != reflect.Func {
		err = ErrMustFunc
		return
	}
	if mT.NumIn() != 3 {
		err = ErrMustHasThreeParam
		return
	}
	reqT = mT.In(0)
	if reqT.Kind() != reflect.Ptr {
		err = ErrMustPtr
		return
	}
	reqT = reqT.Elem()
	if reqT.Kind() != reflect.Struct {
		err = ErrMustPointToStruct
		return
	}
	respT = mT.In(1)
	if respT.Kind() != reflect.Ptr {
		err = ErrMustPtr
		return
	}
	respT = respT.Elem()
	if respT.Kind() != reflect.Struct {
		err = ErrMustPointToStruct
		return
	}
	ctxT := mT.In(2)
	if ctxT.Kind() != reflect.Ptr {
		err = ErrMustPtr
		return
	}
	ctxT = ctxT.Elem()
	if ctxT.Kind() != reflect.Struct {
		err = ErrMustPointToStruct
		return
	}
	if mT.NumOut() != 1 {
		err = ErrMustOneOut
		return
	}
	retT := mT.Out(0)
	if retT != RespErrorType {
		err = ErrMustRespErrorPtr
		return
	}
	return
}

func CreateHandlerFunc(method interface{}) gin.HandlerFunc {
	mV, reqT, respT, err := checkMethod(method)
	if err != nil {
		panic(any(err))
	}

	return func(c *gin.Context) {
		req := reflect.New(reqT)
		if err := c.ShouldBind(req.Interface()); err != nil {
			responseError := errorCode.NewError(errorCode.CODE_PARAM_WRONG, err.Error())
			c.JSON(http.StatusBadRequest, responseError)
			return
		}

		resp := reflect.New(respT)
		ctx := reflect.ValueOf(c)
		rets := mV.Call([]reflect.Value{req, resp, ctx})
		errValue := rets[0]
		if errValue.Interface() != nil {
			if respError := errValue.Interface().(*errorCode.RespError); respError != nil {
				if respError.Msg == "" {
					respError.AutoErrMsg()
				}
				c.JSON(http.StatusOK, respError)
				return
			}
		}

		c.PureJSON(
			http.StatusOK, RespStruct{
				Code: 0,
				Data: resp.Interface(),
			},
		)
	}
}
