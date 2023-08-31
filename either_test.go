package gods

import "testing"

func TestEitherIsLeft(t *testing.T) {
	either_value := Left[int, string](10)

	if !either_value.IsLeft() {
		t.Error("Expected the Either to be a left value but it wasn't :(")
	}
}

func TestEitherIsRight(t *testing.T) {
	either_value := Right[int, string]("10")

	if !either_value.IsRight() {
		t.Error("Expected the Either to be a right value but it wasn't :(")
	}
}

func TestEitherLeftOnLeft(t *testing.T) {
	either_value := Left[int, string](10)

	left_value := either_value.Left()
	if left_value != 10 {
		t.Errorf("Expected the Either's left value to be 10 but it was %d :(", left_value)
	}
}

func TestEitherLeftOnRight(t *testing.T) {
	either_value := Right[int, string]("what")

	shouldPanic(t, func() {
		_ = either_value.Left()
	}, "Expected calling left on a right Either value to panic but it didn't")
}

func TestEitherRightOnRight(t *testing.T) {
	either_value := Right[int, string]("I know this one")

	right_value := either_value.Right()
	if right_value != "I know this one" {
		t.Errorf("Expected the Either's left value to be 10 but it was %s :(", right_value)
	}
}

func TestEitherRightOnLeft(t *testing.T) {
	either_value := Left[int, string](3000)

	shouldPanic(t, func() {
		_ = either_value.Right()
	}, "Expected calling right on a left Either value to panic but it didn't")
}

func TestEitherLeftOrOnLeft(t *testing.T) {
	either_value := Left[int, string](1234)

	left_value := either_value.LeftOr(4321)
	if left_value != 1234 {
		t.Errorf("Expected the LeftOr of a left either value to be 1234 but got %d", left_value)
	}
}

func TestEitherLeftOrOnRight(t *testing.T) {
	either_value := Right[int, string]("You won't see this")

	left_value := either_value.LeftOr(4321)
	if left_value != 4321 {
		t.Errorf("Expected the LeftOr of a right either value to be 4321 but got %d", left_value)
	}
}

func TestEitherRightOrOnLeft(t *testing.T) {
	either_value := Right[int, string]("good value")

	right_value := either_value.RightOr("bad value")
	if right_value != "good value" {
		t.Errorf("Expected the RightOr of a right either value to be 'good value' but got %s", right_value)
	}
}

func TestEitherRightOrOnRight(t *testing.T) {
	either_value := Left[int, string](55)

	right_value := either_value.RightOr("you will see this")
	if right_value != "you will see this" {
		t.Errorf("Expected the LeftOr of a right either value to be 4321 but got %s", right_value)
	}
}
