// Copyright 2017 Ken Miura
package ex03

import (
	"reflect"
	"testing"
)

func TestBool(t *testing.T) {
	test(true, t)
	test(false, t)
}

func TestFloat(t *testing.T) {
	testFloat32(t)
	testFloat64(t)
}

func TestComplex(t *testing.T) {
	testComplex64(t)
	testComplex128(t)
}

func test(b bool, t *testing.T) {
	// Encode it
	data, err := Marshal(b)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var decodedB bool
	if err := Unmarshal(data, &decodedB); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", decodedB)

	// Check equality.
	if decodedB != b {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(b)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func testFloat32(t *testing.T) {
	var f float32 = 3.75
	// Encode it
	data, err := Marshal(f)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var decodedF float32
	if err := Unmarshal(data, &decodedF); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", decodedF)

	// Check equality.
	//if !reflect.DeepEqual(decodedF, f) {
	if decodedF != f {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func testFloat64(t *testing.T) {
	var f float64 = 12345.67890
	// Encode it
	data, err := Marshal(f)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var decodedF float64
	if err := Unmarshal(data, &decodedF); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", decodedF)

	// Check equality.
	if decodedF != f {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func testComplex64(t *testing.T) {
	var c complex64 = 1.25 + 0.75i
	// Encode it
	data, err := Marshal(c)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var decodedC complex64
	if err := Unmarshal(data, &decodedC); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", decodedC)

	// Check equality.
	if decodedC != c {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func testComplex128(t *testing.T) {
	var c complex128 = 12345.67890 + 9876.543210i
	// Encode it
	data, err := Marshal(c)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var decodedC complex128
	if err := Unmarshal(data, &decodedC); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", decodedC)

	// Check equality.
	if decodedC != c {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}
