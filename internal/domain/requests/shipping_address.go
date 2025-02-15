package requests

import "github.com/go-playground/validator/v10"

type CreateShippingAddressRequest struct {
	OrderID        int    `json:"order_id" validate:"required"`
	Alamat         string `json:"alamat" validate:"required,min=5"`
	Provinsi       string `json:"provinsi" validate:"required"`
	Kota           string `json:"kota" validate:"required"`
	Courier        string `json:"courier" validate:"required"`
	ShippingMethod string `json:"shipping_method" validate:"required"`
	ShippingCost   int    `json:"shipping_cost" validate:"required"`
	Negara         string `json:"negara" validate:"required"`
}

type UpdateShippingAddressRequest struct {
	ShippingID     int    `json:"shipping_id" validate:"required"`
	OrderID        int    `json:"order_id" validate:"required"`
	Alamat         string `json:"alamat,omitempty" validate:"omitempty,min=5"`
	Provinsi       string `json:"provinsi,omitempty" validate:"omitempty"`
	Kota           string `json:"kota,omitempty" validate:"omitempty"`
	Courier        string `json:"courier" validate:"required"`
	ShippingMethod string `json:"shipping_method" validate:"required"`
	ShippingCost   int    `json:"shipping_cost" validate:"required"`
	Negara         string `json:"negara,omitempty" validate:"omitempty"`
}

func (l *CreateShippingAddressRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}

func (l *UpdateShippingAddressRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
