use rust::is_number;

#[test]
fn zero() {
    let input = String::from("0");
    let actual = is_number(input);

    assert_eq!(actual, true);
}

#[test]
fn float() {
    let input = String::from("1.0");
    let actual = is_number(input);

    assert_eq!(actual, true);
}

#[test]
fn e() {
    let input = String::from("1e12");
    let actual = is_number(input);

    assert_eq!(actual, true);
}

#[test]
fn positive() {
    let input = String::from("+2");
    let actual = is_number(input);

    assert_eq!(actual, true);
}

#[test]
fn negative() {
    let input = String::from("-2");
    let actual = is_number(input);

    assert_eq!(actual, true);
}

#[test]
fn positive_float() {
    let input = String::from("+2.2");
    let actual = is_number(input);

    assert_eq!(actual, true);
}

#[test]
fn negative_float() {
    let input = String::from("-2.2");
    let actual = is_number(input);

    assert_eq!(actual, true);
}

#[test]
fn positive_e() {
    let input = String::from("+2e2");
    let actual = is_number(input);

    assert_eq!(actual, true);
}

#[test]
fn negative_e() {
    let input = String::from("-2e2");
    let actual = is_number(input);

    assert_eq!(actual, true);
}

#[test]
fn decimal() {
    let input = String::from("2e-2");
    let actual = is_number(input);

    assert_eq!(actual, true);
}

#[test]
fn plus_minus() {
    let input = String::from("+-2");
    let actual = is_number(input);

    assert_eq!(actual, false);
}

#[test]
fn valid_examples() {
    for input in [
        "2",
        "0089",
        "-0.1",
        "+3.14",
        "4.",
        "-.9",
        "2e10",
        "-90E3",
        "3e+7",
        "+6e-1",
        "53.5e93",
        "-123.456e789",
    ] {
        assert!(is_number(input.into()), "expected {input:?} to be valid");
    }
}

#[test]
fn invalid_examples() {
    for input in ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"] {
        assert!(!is_number(input.into()), "expected {input:?} to be invalid");
    }
}

#[test]
fn decimal_boundaries() {
    for input in [".0", "+.9", "0.", "-0.", "1.e2", ".1E-2", "000.000"] {
        assert!(is_number(input.into()), "expected {input:?} to be valid");
    }

    for input in [".", "+.", "-.", "..", "1..2", "1.2.3", ".e2", "1e.2"] {
        assert!(!is_number(input.into()), "expected {input:?} to be invalid");
    }
}

#[test]
fn incomplete_numbers() {
    for input in ["", "+", "-", "e", "E", "1e+", "1E-", "+e1", "-e1"] {
        assert!(!is_number(input.into()), "expected {input:?} to be invalid");
    }
}

#[test]
fn exponent_boundaries() {
    for input in ["0e0", "1E+2", "1e-0", "1e001", "9e99999999999999999999"] {
        assert!(is_number(input.into()), "expected {input:?} to be valid");
    }

    for input in ["1ee2", "1e2E3", "1e++2", "1e+-2", "1e--2", "1e2.0", "1e2+"] {
        assert!(!is_number(input.into()), "expected {input:?} to be invalid");
    }
}

#[test]
fn misplaced_signs() {
    for input in ["++1", "1+2", "1-2", "1+", ".-1", "1.-2"] {
        assert!(!is_number(input.into()), "expected {input:?} to be invalid");
    }
}

#[test]
fn whitespace_and_non_ascii_characters() {
    for input in [" ", " 1", "1 ", "1 2", "1\n", "\t1", "١", "１", "−1", "1,2"] {
        assert!(!is_number(input.into()), "expected {input:?} to be invalid");
    }
}
