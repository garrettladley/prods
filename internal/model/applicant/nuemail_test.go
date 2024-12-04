package applicant

import "testing"

func TestParseNUEmail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		email string
	}{
		{
			name:  "Email with number works",
			email: "ladley.g1@northeastern.edu",
		},
		{
			name:  "Email without number works",
			email: "ladley.g@northeastern.edu",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			nuEmail, err := ParseNUEmail(test.email)
			if err != nil {
				t.Fatalf("expected no error, got %s", err.Error())
			}
			if nuEmail.String() != test.email {
				t.Fatalf("expected %s, got %s", test.email, nuEmail.String())
			}
		})
	}
}
