package test

import (
	"test/entities"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestValidateMember(t *testing.T) {
	RegisterTestingT(t)

	g := NewGomegaWithT(t)

	t.Run("Username is valid", func(t *testing.T) {
		member := entities.Member{
			Username: "SA Test", 
			Password: "123456",
			Email: "sa@gmail.com",
		}
		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run("Username is required", func(t *testing.T) {
		member := entities.Member{
			Username: "", 	// ผิดตรงนี้
			Password: "123456",
			Email: "sa@gmail.com",
		}
		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Username is required"))
	})

	t.Run("Email is valid", func(t *testing.T) {
		member := entities.Member{
			Username: "SA Test", 
			Password: "123456",
			Email: "sa@gmail.com",
		}
		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run("Email is required", func(t *testing.T) {
		member := entities.Member{
			Username: "SA Test", 
			Password: "123456",
			Email: "",	// ผิดตรงนี้
		}
		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})

	t.Run("Email format is incorrect", func(t *testing.T) {
		member := entities.Member{
			Username: "SA Test", 
			Password: "123456",
			Email: "sa@.com",	// ผิดตรงนี้
		}
		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})
}