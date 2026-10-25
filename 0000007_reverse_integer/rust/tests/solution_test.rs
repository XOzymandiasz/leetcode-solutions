use rust::reverse;

#[test]
fn one_hundred() {
    let actual = reverse(100);

    assert_eq!(actual, 1);
}

#[test]
fn basic() {
    let actual = reverse(123);

    assert_eq!(actual, 321);
}

#[test]
fn negative() {
    let actual = reverse(-123);

    assert_eq!(actual, -321);
}

#[test]
fn single_number() {
    let actual = reverse(1);

    assert_eq!(actual, 1);
}

#[test]
fn high_number() {
    let actual = reverse(1534236469);

    assert_eq!(actual, 0);
}
