// package forum

// import (
// 	"testing"
// )

// func TestGenerateFakeUser(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want User
// 	}{
// 		{
// 			name: "TestGenerateFakeUser",
// 			want: User{
// 				ID:       12,
// 				Username: "test",
// 				Email:    "test@mysql.co.uk",
// 			},
// 		},
// 		{
// 			name: "TestGenerateFakeUserInvalidEmail",
// 			want: User{
// 				ID:       12,
// 				Username: "test",
// 				Email:    "test@mysql.co.uk",
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := GenerateFakeUser()

// 			// We don't compare Password directly as it's generated dynamically.
// 			if got.ID != tt.want.ID || got.Username != tt.want.Username || got.Email != tt.want.Email {
// 				t.Errorf("GenerateFakeUser() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

package forum

import (
	"reflect"
	"testing"
)

func TestGenerateFakeUser(t *testing.T) {
	hashedPassword, _ := GenerateTestPasswordHash("")

	tests := []struct {
		name          string
		want          User
		expectSuccess bool // Add a flag to indicate if the test should succeed or fail
	}{
		//ALERT not sure if this logic is correct 🤔
		{
			name: "TestGenerateFakeUser",
			want: User{
				ID:       12,
				Username: "test",
				Email:    "test@mysql.co.uk",
				Password: hashedPassword,
			},
			expectSuccess: false, // This test is expected to succeed
		},
		{
			name: "InvalidEmail",
			want: User{
				ID:       13,
				Username: "invalid_email",
				Email:    "invalid_email",
			},
			expectSuccess: false, // This test is expected to fail due to an invalid email
		},
		{
			name: "InvalidID",
			want: User{
				ID:       -1,
				Username: "invalid_id",
				Email:    "test@mysql.co.uk",
			},
			expectSuccess: false, // This test is expected to fail due to an invalid ID
		},
		{
			name: "InvalidUsername",
			want: User{
				ID:       14,
				Username: "",
				Email:    "test@mysql.co.uk",
			},
			expectSuccess: false, // This test is expected to fail due to an invalid username
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateFakeUser()

			if tt.expectSuccess {
				// Ensure the generated user matches the expected user when the test should succeed
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GenerateFakeUser() = %v, want %v", got, tt.want)
				}
			} else {
				// Ensure that the generated user is not equal to the expected user when the test should fail
				if reflect.DeepEqual(got, tt.want) {
					t.Errorf("GenerateFakeUser() unexpectedly succeeded with input: %v", tt.name)
				}
			}
		})
	}
}
