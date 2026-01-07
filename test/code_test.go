package test

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCode(t *testing.T) {
	g := NewGomegaWithT(t)

	valid_book := entity.Books{
		Title: "JakePual",
		Price: 55,
		Code:  "BK299322",
	}

	t.Run("Case:2 Price", func(t *testing.T) {
		v := valid_book
		v.Code = "B3333333"
		ok, err := govalidator.ValidateStruct(v)
		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
	})
}
