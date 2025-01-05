package shortener

import "testing"

func TestShortenURL(t *testing.T) {
	tests := []struct {
		inputURL    string
		expectedURL string
	}{
		{"https://www.example.com/very/long/url/to/be/shortened", "29EFYe"},
		{"https://www.google.com", "30orQW"},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.inputURL, func(t *testing.T) {
			shortened := ShortenURL(test.inputURL)
			if shortened != test.expectedURL {
				t.Errorf("Expected: %s, Got: %s", test.expectedURL, shortened)
			}
		})
	}
}
