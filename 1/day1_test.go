import {
	testing
	os
}

func TestPart1(t *testing.T) {

	b, err := os.ReadFile('test1.txt')
	
	tests := []struct{
		input []byte
		expected int
	}{
		{
			expected: 142
			input: b
		}
	}

	for test : range tests {
		assert.Equal()
	}
}