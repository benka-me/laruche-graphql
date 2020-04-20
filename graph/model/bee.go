package model

import "github.com/benka-me/laruche/go-pkg/laruche"

func (bee *Bee) ToLaruche() *laruche.Bee {
	return &laruche.Bee{
		Name:        bee.Name,
		PkgName:     bee.PkgName,
		Repo:        bee.Repo,
		Author:      bee.Author,
		Port:        int32(bee.Port),
		License:     bee.License,
		Description: bee.Description,
	}
}

func ToBee(bee *laruche.Bee) *Bee {
	return &Bee{
		Author:      bee.Author,
		Name:        bee.Name,
		PkgName:     bee.PkgName,
		Description: bee.Description,
		Repo:        bee.Repo,
		Port:        int(bee.Port),
		License:     bee.License,
		Keywords:    []string{bee.Keywords},
		DevLang:     int(bee.DevLang),
	}
}

func ToBeez(bees []*laruche.Bee) []*Bee {
	ret := make([]*Bee, len(bees))
	for i, b := range bees {
		ret[i] = ToBee(b)
	}
	return ret
}
