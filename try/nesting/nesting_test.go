package nesting

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Duck string

func (d Duck) Hi() string {
	return "Gua"
}

func (s *NestingSuite) TestNestInterface() {
	type Itf interface {
		Hi() string
	}

	type A struct {
		Itf
		Age int
	}

	a := &A{
		Itf: Duck(""),
	}
	s.EqualValues("", a.Itf)
	s.EqualValues("Gua", a.Hi()) // Method up from Itf
	s.EqualValues("Gua", a.Itf.Hi())
}

func (s *NestingSuite) TestNestingJson() {
	blog := &Blog{
		User: User{
			Name:  "a",
			Email: "b@c.com",
		},
		Title: "test",
	}
	bs, err := json.Marshal(blog)
	s.Nil(err)
	js := `{"name":"a","email":"b@c.com","title":"test"}`
	s.EqualValues(js, string(bs))

	s.Equal(5, blog.Score())
	s.Equal(5, blog.User.Score())

	blog1 := &Blog{}
	json.Unmarshal([]byte(js), blog1)
	s.Equal("a", blog1.Name)
	s.Equal("test", blog1.Title)
}

func (s *NestingSuite) TestOmitJson() {
	blog := &Blog{
		User: User{
			Name: "a",
		},
		Title: "test",
	}
	bs, err := json.Marshal(blog)
	s.Nil(err)
	js := `{"name":"a","title":"test"}` // no email part
	s.Equal(js, string(bs))

	blog = &Blog{
		Title: "test",
	}
	bs, err = json.Marshal(blog)
	s.Nil(err)
	js = `{"name":"","title":"test"}` // no email part
	s.Equal(js, string(bs))
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

func (s User) Score() int {
	return 5
}

type Blog struct {
	User
	Title string `json:"title"`
}

func TestNestingSuite(t *testing.T) {
	suite.Run(t, &NestingSuite{})
}

type NestingSuite struct {
	suite.Suite
}
