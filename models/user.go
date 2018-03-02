// Copyright 2018 Lothar . All rights reserved.
// https://github.com/GaWaine1223

package models

import (
	"path"

	"github.com/GaWaine1223/Lothar/keygen"

	"github.com/GaWaine1223/Lothar/freecell/common"
)

type User struct{
	Name	string
	Path	string
	Public  string
	Private string
}

func Login(name string) (*User, error) {
	if err := keygen.GenRsaKey(common.RSADefaultLenth, name); err != nil {
		return nil, err
	}
	uPath := keygen.GetUserPath(name)
	pvKeyPath := path.Join(uPath, "private.pem")
	pbKeyPath := path.Join(uPath, "public.pem")
	pv, err:= keygen.GetKeyMd5(pvKeyPath)
	if err != nil {
		return nil, err
	}
	pb, err := keygen.GetKeyMd5(pbKeyPath)
	if err != nil {
		return nil, err
	}
	return &User{
		Name: name,
		Path: uPath,
		Public: pb,
		Private: pv}, nil
}
