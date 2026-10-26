use rust::Solution;

#[test]
fn three_elements() {
    let mut result = Solution::permute_unique(vec![1, 2, 3]);
    let mut expected = vec![
        vec![1, 2, 3],
        vec![1, 3, 2],
        vec![2, 1, 3],
        vec![2, 3, 1],
        vec![3, 1, 2],
        vec![3, 2, 1],
    ];

    result.sort();
    expected.sort();

    assert_eq!(result, expected);
}

#[test]
fn single_element() {
    assert_eq!(Solution::permute_unique(vec![7]), vec![vec![7]]);
}

#[test]
fn elements_with_duplicates() {
    let mut result = Solution::permute_unique(vec![1, 1, 2]);
    let mut expected = vec![vec![1, 1, 2], vec![1, 2, 1], vec![2, 1, 1]];

    result.sort();
    expected.sort();

    assert_eq!(result, expected);
}
