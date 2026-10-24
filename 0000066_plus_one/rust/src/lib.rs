pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
    let mut digits = digits;

    for digit in digits.iter_mut().rev() {
        if *digit < 9 {
            *digit += 1;
            return digits;
        }
        *digit = 0
    }

    digits.insert(0, 1);
    digits
}
