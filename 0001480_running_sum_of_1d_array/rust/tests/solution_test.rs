use rust::Solution;

#[test]
fn first_case() {
    let actual = Solution::running_sum(vec![1, 1, 1, 1, 1]);

    assert_eq!(actual, vec![1, 2, 3, 4, 5]);
}

#[test]
fn second_case() {
    let actual = Solution::running_sum(vec![1, 2, 3, 4]);

    assert_eq!(actual, vec![1, 3, 6, 10]);
}
