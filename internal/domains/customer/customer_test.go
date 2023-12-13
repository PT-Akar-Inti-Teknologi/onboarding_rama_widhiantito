package customer

import "testing"

func TestCustomer_Validate(t *testing.T) {
	validCustomer := &Customer{
		Name:        "John Doe",
		Phonenumber: "123456789012", // Invalid length, should be 12 digits
		Email:       "johndoe@example.com",
		Address:     "123 Main St",
	}

	invalidCustomer := &Customer{
		Name:        "",             // Missing required field
		Phonenumber: "1234567890",   // Incorrect length, should be 12 digits
		Email:       "invalidemail", // Invalid email format
		Address:     "456 Elm St",
	}

	tests := []struct {
		name    string
		c       *Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "ValidCustomer",
			c:       validCustomer,
			wantErr: false, // Expecting error due to invalid fields
		},
		{
			name:    "InvalidCustomer",
			c:       invalidCustomer,
			wantErr: true, // Expecting error due to invalid fields
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.c.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Customer.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
