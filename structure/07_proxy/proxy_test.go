package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {

	proxyCompany := RegistryProxyCompany{p: Pioneer{10000, UK}}

	//想注册一个叫fire company 的公司
	name, no, err := proxyCompany.RegisterCompany("fire company")

	if err != nil {
		t.Log("something wrong:", err.Error())
	} else {

		t.Log("got company, name:", name, "company no:", no)

	}

	proxyCompany = RegistryProxyCompany{p: Pioneer{109999, CN}}

	//想注册一个叫fire company 的公司
	name, no, err = proxyCompany.RegisterCompany("water company")

	if err != nil {
		t.Log("something wrong:", err.Error())
	} else {

		t.Log("got company, name:", name, "company no:", no)
	}
}
