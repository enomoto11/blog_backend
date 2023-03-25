package request

type POSTCategoryRequestBody struct {
	Name string `validate:"required"`
}
