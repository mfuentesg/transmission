package transmission

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumBool_UnmarshalJSON(t *testing.T) {
	type testStruct struct {
		Field1 NumBool `json:"field"`
	}

	tests := []struct {
		body     []byte
		expected bool
	}{
		{body: []byte(`{}`), expected: false},
		{body: []byte(`{"field": 0}`), expected: false},
		{body: []byte(`{"field": -1}`), expected: false},
		{body: []byte(`{"field": ""}`), expected: false},
		{body: []byte(`{"field": false}`), expected: false},
		{body: []byte(`{"field": 1}`), expected: true},
		{body: []byte(`{"field": true}`), expected: true},
	}

	for _, test := range tests {
		var ts testStruct

		err := json.Unmarshal(test.body, &ts)
		assert.NoError(t, err)
		assert.Nil(t, err)
		assert.Equal(t, NumBool(test.expected), ts.Field1)
	}
}
