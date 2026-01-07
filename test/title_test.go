package test

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_title(t *testing.T) {
	g := NewGomegaWithT(t)

	valid_book := entity.Books{
		Title: "JakePual",
		Price: 55,
		Code:  "BK299322",
	}

	t.Run("Case:1 Title", func(t *testing.T) {
		v := valid_book
		v.Title = "Jd"
		ok, err := govalidator.ValidateStruct(v)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("Name of the book must contain 3-100 characters"))
	})
}
