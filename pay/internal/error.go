package internal

import (
	"errors"
)

type ReturnError struct {
	ReturnCode string `xml:"return_code"` //返回状态码
	ReturnMsg  string `xml:"return_msg"`  //返回信息
}

func (r *ReturnError) Check() error {
	if r.ReturnCode == "FAIL" {
		return errors.New(r.ReturnMsg)
	}
	return nil
}

type ResultError struct {
	ResultCode string `xml:"result_code"`  //业务结果
	ErrCode    string `xml:"err_code"`     //错误代码
	ErrCodeDes string `xml:"err_code_des"` //错误代码描述
}

func (r *ResultError) Check() error {
	if r.ResultCode == "FAIL" {
		return errors.New(r.ErrCode)
	}
	return nil
}

type ResponseError struct {
	ReturnError
	ResultError
}

func (r *ResponseError) Check() error {
	if err := r.ReturnError.Check(); err != nil {
		return err
	}
	if err := r.ResultError.Check(); err != nil {
		return err
	}
	return nil
}
