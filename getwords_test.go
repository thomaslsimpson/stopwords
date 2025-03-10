package stopwords

import (
	"testing"
)

func TestGetWords(t *testing.T) {
	result := GetWords("en")
  
	if len(result) < 1 {
		t.Errorf("Test failed, got no words")
	}

}



