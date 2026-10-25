use rust::Solution;

#[test]
fn zero_and_one() {
    assert_eq!(Solution::my_sqrt(0), 0);
    assert_eq!(Solution::my_sqrt(1), 1);
}

#[test]
fn perfect_squares() {
    for (x, expected) in [(4, 2), (9, 3), (16, 4), (100, 10), (2_147_395_600, 46_340)] {
        assert_eq!(Solution::my_sqrt(x), expected, "x = {x}");
    }
}

#[test]
fn non_squares_are_rounded_down() {
    for (x, expected) in [(2, 1), (3, 1), (8, 2), (15, 3), (17, 4), (99, 9), (101, 10)] {
        assert_eq!(Solution::my_sqrt(x), expected, "x = {x}");
    }
}

#[test]
fn large_inputs_do_not_overflow() {
    for (x, expected) in [
        (2_147_395_599, 46_339),
        (2_147_395_601, 46_340),
        (i32::MAX, 46_340),
    ] {
        assert_eq!(Solution::my_sqrt(x), expected, "x = {x}");
    }
}

#[test]
fn result_satisfies_floor_square_root_bounds() {
    for x in 0..=10_000 {
        let result = i64::from(Solution::my_sqrt(x));
        let input = i64::from(x);
        assert!(result >= 0, "x = {x}, result = {result}");
        assert!(result * result <= input, "x = {x}, result = {result}");
        assert!(
            (result + 1) * (result + 1) > input,
            "x = {x}, result = {result}"
        );
    }
}
