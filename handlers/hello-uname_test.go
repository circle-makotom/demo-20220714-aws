package handlers

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestGenResponse(t *testing.T) {
	ts := time.Now()

	if actual, err := json.Marshal(new(HelloUnameHandler).BuildResponse(ts, "dummy uname")); err != nil {
		t.Error(err)
	} else {
		assert.Equal(t, string(actual), fmt.Sprintf(`{"message":"Hello world!","at":"%s","uname":"dummy uname","ctr":1}`, ts.Format(time.RFC3339Nano)))
	}
}
