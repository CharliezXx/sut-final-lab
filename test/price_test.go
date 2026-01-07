package test

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPrice(t *testing.T) {
	g := NewGomegaWithT(t)

	valid_book := entity.Books{
		Title: "JakePual",
		Price: 55,
		Code:  "BK299322",
	}

	t.Run("Case:2 Price", func(t *testing.T) {
		v := valid_book
		v.Price = 9999
		ok, err := govalidator.ValidateStruct(v)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("Price must between 50 and 5000"))
	})
}
