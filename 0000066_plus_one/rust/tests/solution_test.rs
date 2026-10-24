use rust::plus_one;

#[test]
fn basic() {
    let nums = vec![1, 2, 3];
    let actual = plus_one(nums.clone());

    assert_eq!(vec![1, 2, 4], actual);
}

#[test]
fn should_push_one() {
    let nums = vec![9];
    let actual = plus_one(nums.clone());

    assert_eq!(vec![1, 0], actual);
}

#[test]
fn should_push_one_multiple_digits() {
    let nums = vec![9, 9, 9, 9];
    let actual = plus_one(nums.clone());

    assert_eq!(vec![1, 0, 0, 0, 0], actual);
}
