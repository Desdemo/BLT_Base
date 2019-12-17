package base

import (
	"errors"
)

type Common struct {
	Table string
	Data interface{}
	Id int
	Name string
	Version int64
}

func (com Common)GetForId() (data interface{}, err error)  {
	data, err = Engine.Table(com.Table).ID(com.Id).Get(com.Data)
	return
}

func (com *Common)ExistForId() (status bool, err error)  {
	status, err = Engine.Table(com.Table).ID(com.Id).Exist(com.Data)
	return
}

func (com *Common)ExistForName() (status bool, err error)  {
	status, err = Engine.Table(com.Table).Where("name = ?", com.Name).Exist(com.Data)
	return
}

func (com *Common)Add() (status bool, err error) {
	has, err := com.ExistForName()
	if err != nil{
		return false,err
	}
	if !has {
		affected, err := Engine.Table(com.Table).Insert(com.Data)
		if err != nil {
			return false, err
		}
		if affected == 1 {
			return true, nil
		} else {
			return false,nil
		}
	} else {
		return false, errors.New("当前名称已存在")
	}
}

func (com *Common)Delete() (status bool, err error) {
	has, err := com.ExistForId()
	if err != nil {
		return false, err
	}
	if has {
		affected, err := Engine.Table(com.Table).ID(com.Id).Delete(com.Data)
		if err != nil {
			return false, err
		}
		if affected == 1 {
			return true, nil
		} else {
			return false, err
		}
	} else {
		return false, errors.New("当前数据不存在或已被删除")
	}
}