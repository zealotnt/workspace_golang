package handlers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
)

type BindingTest struct {
	body   string
	expect string
}

type ValidateTest struct {
	body   string
	expect string
}

var _ = Describe("GetProductsHandler", func() {
	It("returns 200", func() {
		response := Request("GET", "/products", "")

		Expect(response.Code).To(Equal(200))
	})
})

var create_product_binding_tests = []BindingTest{
	{`{"name": "XBox","price": "70000","provider": "Microsoft","rating": 3.5,"status": "sale"}`, `"json: cannot unmarshal string into Go value of type int"`},
	{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": "3.5","status": "sale"}`, `"json: cannot unmarshal string into Go value of type float32"`},
}

var _ = Describe("CreateProductHandlerBindingError", func() {
	It("returns 400", func() {
		for _, test := range create_product_binding_tests {
			response := Request("POST", "/products", test.body)
			Expect(response.Code).To(Equal(400))
			Expect(response.Body).To(ContainSubstring(test.expect))
		}
	})
})

var create_product_validate_tests = []ValidateTest{
	{`{"price": 70000,"provider": "Microsoft","rating": 3.5,"status": "sale"}`, `"Name is required"`},
	{`{"name": "XBox","provider": "Microsoft","rating": 3.5,"status": "sale"}`, `"Price is required"`},
	{`{"name": "XBox","price": 70000,"rating": 3.5,"status": "sale"}`, `"Provider is required"`},
	{`{"name": "XBox","price": 70000,"provider": "Microsoft","status": "sale"}`, `"Rating is required"`},
	{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": 3.5}`, `"Status is required"`},
	{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": 3.5,"status": "sale"}`, `"Image is required"`},
	{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": 6.0,"status": "sale"}`, `"Rating must be less than or equal to 5"`},
	{`{"name": "XBox","price": 70000,"provider": "Microsoft","rating": 3.5,"status": "on sale"}`, `"Status is invalid"`},
}

var _ = Describe("CreateProductHandlerValidateError", func() {
	It("returns 422", func() {
		for _, test := range create_product_validate_tests {
			response := Request("POST", "/products", test.body)
			Expect(response.Code).To(Equal(422))
			Expect(response.Body).To(ContainSubstring(test.expect))
		}
	})
})

var _ = Describe("CreateProductHandlerSuccess", func() {
	It("returns 201", func() {
		form := url.Values{}
		form.Add("name", "XBox")
		form.Add("price", "70000")
		form.Add("provider", "Microsoft")
		form.Add("rating", "3.5")
		form.Add("status", "sale")
		fmt.Println(form.Encode())

		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)
		// this step is very important
		fileWriter, err := bodyWriter.CreateFormFile("uploadfile", "filename.png")
		if err != nil {
			fmt.Println("error writing to buffer")
		}
		// open file handle
		fh, err := os.Open("/go/src/github.com/o0khoiclub0o/piflab-store-api-go/db/seeds/factory/golang.png")
		if err != nil {
			fmt.Println("error opening file")
		}

		//iocopy
		io.Copy(fileWriter, fh)

		contentType := bodyWriter.FormDataContentType()
		bodyWriter.Close()

		response := Request("POST", "/products", string(bodyBuf))
		// response := Request("POST", "/products", form.Encode())
		Expect(response.Code).To(Equal(422))
		Expect(response.Body).To(ContainSubstring("123asd"))
	})
})

var _ = Describe("UpdateProductHandlerInvalidRecord", func() {
	It("returns 404", func() {
		response := Request("PUT", "/products/abc", `{"name": "XBox"}`)

		Expect(response.Code).To(Equal(404))
		Expect(response.Body).To(ContainSubstring(`"record not found"`))
	})
})

var _ = Describe("UpdateProductHandlerZeroRecord", func() {
	It("returns 404", func() {
		response := Request("PUT", "/products/0", `{"name": "XBox"}`)

		Expect(response.Code).To(Equal(404))
		Expect(response.Body).To(ContainSubstring(`"record not found"`))
	})
})

var update_product_binding_tests = []BindingTest{
	{`{"rating": "3.4"}`, `"json: cannot unmarshal string into Go value of type float32"`},
	{`{"price": "123"}`, `"json: cannot unmarshal string into Go value of type int"`},
}

var _ = Describe("UpdateProductHandlerBindingError", func() {
	It("returns 400", func() {
		for _, test := range update_product_binding_tests {
			response := Request("PUT", getFirstAvailableUrl(), test.body)
			Expect(response.Code).To(Equal(400))
			Expect(response.Body).To(ContainSubstring(test.expect))
		}
	})
})

var update_product_validate_tests = []ValidateTest{
	{`{"rating": 5.1}`, `"Rating must be less than or equal to 5"`},
	{`{"status": "on sale"}`, `"Status is invalid"`},
}

var _ = Describe("UpdateProductHandlerValidateError", func() {
	It("returns 422", func() {
		for _, test := range update_product_validate_tests {
			response := Request("PUT", getFirstAvailableUrl(), test.body)
			Expect(response.Code).To(Equal(422))
			Expect(response.Body).To(ContainSubstring(test.expect))
		}
	})
})

var _ = Describe("UpdateProductHandlerValidateRatingError", func() {
	It("returns 200", func() {
		response := Request("PUT", getFirstAvailableUrl(), `{"rating": 4.0}`)
		Expect(response.Code).To(Equal(200))
	})
})
