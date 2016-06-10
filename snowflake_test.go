package snowflake

import "testing"

func TestMarshalJSON(t *testing.T) {
	id := ID(13587)
	expected := "\"13587\""

	bytes, err := id.MarshalJSON()
	if err != nil {
		t.Error("Unexpected error during MarshalJSON")
	}

	if string(bytes) != expected {
		t.Errorf("Got %s, expected %s", string(bytes), expected)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	strID := "\"13587\""
	expected := ID(13587)

	var id ID
	err := id.UnmarshalJSON([]byte(strID))
	if err != nil {
		t.Error("Unexpected error during UnmarshalJSON")
	}

	if id != expected {
		t.Errorf("Got %d, expected %d", id, expected)
	}
}

func BenchmarkGenerate(b *testing.B) {

	node, _ := NewNode(1)

	b.ReportAllocs()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = node.Generate()
	}
}

func BenchmarkUnmarshal(b *testing.B) {

	node, _ := NewNode(1)
	id := node.Generate()
	var id2 ID

	b.ReportAllocs()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = id2.UnmarshalJSON(id.Bytes())
	}
}

func BenchmarkMarshal(b *testing.B) {

	node, _ := NewNode(1)
	id := node.Generate()

	b.ReportAllocs()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = id.MarshalJSON()
	}
}
