package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)


func TestPositive(t *testing.T){
	g := gomega.NewGomegaWithT(t)

	employee := Employee{
		Name: "Woottikrai Sangkomon",
		Email: "woottikraij@gmail.com",
		EmployeeID: "J12345678",
	}

	ok,err := govalidator.ValidateStruct(employee)

	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())

}


