use rust::two_sum;

fn assert_valid(nums: &[i32], target: i32, result: &[i32]) {
    let i = result[0] as usize;
    let j = result[1] as usize;

    assert_ne!(i, j, "indices must be different: {:?}", result);

    let sum = nums[i] + nums[j];
    assert_eq!(
        sum, target,
        "nums[{}]={} + nums[{}]={} = {}, expected {}",
        i, nums[i], j, nums[j], sum, target
    );
}

#[test]
fn basic() {
    let nums = vec![2, 7, 11, 15];
    let result = two_sum(nums.clone(), 9);
    assert_valid(&nums, 9, &result);
}

#[test]
fn unordered() {
    let nums = vec![3, 2, 4];
    let result = two_sum(nums.clone(), 6);
    assert_valid(&nums, 6, &result);
}

#[test]
fn duplicates() {
    let nums = vec![3, 3];
    let result = two_sum(nums.clone(), 6);
    assert_valid(&nums, 6, &result);
}

#[test]
fn negative() {
    let nums = vec![7, 3, 4, -1];
    let result = two_sum(nums.clone(), 2);
    assert_valid(&nums, 2, &result);
}

#[test]
fn no_solution() {
    let nums = vec![1];
    let result = two_sum(nums.clone(), 2);
    assert_eq!(result, vec!{}, "[{}, {}], expected []", result[0], result[1]);
}