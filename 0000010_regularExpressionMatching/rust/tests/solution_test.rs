use rust::Solution;

#[test]
fn literals_must_match_the_entire_string() {
    let cases = [
        ("a", "a", true),
        ("ab", "ab", true),
        ("a", "b", false),
        ("aa", "a", false),
        ("a", "aa", false),
        ("ab", "ac", false),
    ];

    check_cases(&cases);
}

#[test]
fn dot_matches_exactly_one_character() {
    let cases = [
        ("a", ".", true),
        ("ab", ".", false),
        ("ab", "a.", true),
        ("abc", "a.c", true),
        ("abc", "a.d", false),
    ];

    check_cases(&cases);
}

#[test]
fn star_matches_zero_or_more_occurrences_of_the_previous_character() {
    let cases = [
        ("aaa", "a*", true),
        ("aaaba", "a*", false),
        ("aaab", "a*b", true),
        ("b", "a*b", true),
        ("ab", "ab*", true),
        ("a", "ab*", true),
        ("abbb", "ab*", true),
        ("ac", "ab*", false),
        ("aab", "c*a*b", true),
    ];

    check_cases(&cases);
}

#[test]
fn dot_star_matches_any_sequence_but_still_requires_a_full_match() {
    let cases = [
        ("ab", ".*", true),
        ("abc", ".*c", true),
        ("ab", ".*c", false),
        ("abcd", "a.*d", true),
    ];

    check_cases(&cases);
}

#[test]
fn star_can_leave_characters_for_the_rest_of_the_pattern() {
    let cases = [
        ("aaa", "a*a", true),
        ("aaa", "a*aa", true),
        ("aaa", "a*aaaa", false),
        ("aaa", "ab*a*c*a", true),
        ("mississippi", "mis*is*p*.", false),
        ("mississippi", "mis*is*ip*.", true),
    ];

    check_cases(&cases);
}

fn check_cases(cases: &[(&str, &str, bool)]) {
    for &(s, p, expected) in cases {
        assert_eq!(
            Solution::is_match(s.to_owned(), p.to_owned()),
            expected,
            "s = {s:?}, p = {p:?}"
        );
    }
}
