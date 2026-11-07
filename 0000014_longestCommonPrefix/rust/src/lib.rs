pub struct Solution;

impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        let mut answer = String::from("");
        let mut idx = 0;
        loop {
            let Some(c) = strs[0].chars().nth(idx) else {
                break;
            };
            let mut not_match = false;
            for str in strs.iter() {
                if str.chars().nth(idx) != Some(c) {
                    not_match = true;
                }
            }
            if not_match {
                break;
            }
            answer.push(c);
            idx += 1;
        }
        answer
    }
}
