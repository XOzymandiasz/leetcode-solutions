use rust::Solution;

#[test]
fn matching() {
    let strs = vec![
        String::from("Flower"),
        String::from("Flow"),
        String::from("Flight"),
    ];
    let actual = Solution::longest_common_prefix(strs);

    assert_eq!(actual, "Fl");
}

#[test]
fn not_matching() {
    let strs = vec![
        String::from("Flower"),
        String::from("Watermelon"),
        String::from("Cucumber"),
    ];
    let actual = Solution::longest_common_prefix(strs);

    assert_eq!(actual, "");
}
