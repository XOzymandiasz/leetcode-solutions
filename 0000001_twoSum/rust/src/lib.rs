pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    use std::collections::HashMap;
    let mut visited: HashMap<i32, i32> = HashMap::with_capacity(nums.len());

    for (i, &n) in nums.iter().enumerate() {
        if let Some(&idx) = visited.get(&(target - n)) {
            return vec![idx, i as i32];
        }
        visited.insert(n, i as i32);
    }
    vec!{}
}