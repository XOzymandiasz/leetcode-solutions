use rust::can_place_flowers;

#[test]
fn base() {
    assert_eq!(true, can_place_flowers(vec![1,0,0,0,1], 1));
}

#[test]
fn to_many_flowers() {
    assert_eq!(false, can_place_flowers(vec![1,0,0,0,1], 2));
}

#[test]
fn one_element() {
    assert_eq!(false, can_place_flowers(vec![1], 1));
}

#[test]
fn one_element_with_no_one_to_set() {
    assert_eq!(true, can_place_flowers(vec![1], 0));
}

#[test]
fn one_element_without_flowers() {
    assert_eq!(true, can_place_flowers(vec![0], 1));
}