package main

import "testing"

func TestContainsFilterWord(t *testing.T) {
	tests := []struct {
		input  string
		words  []string
		result bool
	}{
		{"Is this about alcohol?", []string{"alcohol", "18+", "drugs"}, true},
		{"This is a test", []string{"alcohol", "18+", "drugs"}, false},
		{"Are you 18+?", []string{"alcohol", "18+", "drugs"}, true},
		{"This is a drugs related question", []string{"alcohol", "18+", "drugs"}, true},
		{"No filtered words here", []string{"alcohol", "18+", "drugs"}, false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := containsFilterWord(test.input, test.words)
			if result != test.result {
				t.Errorf("Expected containsFilterWord(%q, %v) to be %v, got %v", test.input, test.words, test.result, result)
			}
		})
	}
}

func TestEmptyInput(t *testing.T) {
	result := containsFilterWord("", []string{"alcohol", "18+", "drugs"})
	if result != false {
		t.Errorf("Expected containsFilterWord('', %v) to be false, got %v", []string{"alcohol", "18+", "drugs"}, result)
	}
}

func TestMultipleFilterWords(t *testing.T) {
	result := containsFilterWord("This is about alcohol and drugs", []string{"alcohol", "18+", "drugs"})
	if result != true {
		t.Errorf("Expected containsFilterWord('This is about alcohol and drugs', %v) to be true, got %v", []string{"alcohol", "18+", "drugs"}, result)
	}
}

func TestUppercaseFilterWords(t *testing.T) {
	result := containsFilterWord("Is this about ALCOHOL?", []string{"alcohol", "18+", "drugs"})
	if result != true {
		t.Errorf("Expected containsFilterWord('Is this about ALCOHOL?', %v) to be true, got %v", []string{"alcohol", "18+", "drugs"}, result)
	}
}

func TestLowerCaseFunc(t *testing.T) {
	result := toLowerCase("HELLO")
	if result != "hello" {
		t.Errorf("Expected containsFilterWord('Is this about ALCOHOL?', %v) to be true, got %v", []string{"alcohol", "18+", "drugs"}, result)
	}
}
