package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)


func TestNegative(t *testing.T){
	g := gomega.NewGomegaWithT(t)

	employee := Employee{
		Name: "",
		Email: "woottikraij@gmail.com",
		EmployeeID: "J1234567",
	}

	ok,err := govalidator.ValidateStruct(employee)

	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("name cant ba blank"))

}