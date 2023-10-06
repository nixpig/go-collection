pub fn binary_search(arr: Vec<isize>, target: isize) -> isize {
    let mut index_high = arr.len() as isize - 1;
    let mut index_low: isize = 0;

    while index_low <= index_high {
        let index_mid = (((index_high + index_low) / 2) as f64).floor() as isize;
        let current_value = arr[index_mid as usize];

        if current_value == target {
            return index_mid;
        }

        if current_value > target {
            index_high = index_mid - 1;
        } else {
            index_low = index_mid + 1;
        }
    }

    -1
}

#[test]
fn test_binary_search_upper_half() {
    let v: Vec<isize> = vec![1, 7, 23, 42, 75, 116, 794, 800, 1099];

    let found = binary_search(v, 800);
    assert_eq!(found, 7);
}

#[test]
fn test_binary_search_lower_half() {
    let v: Vec<isize> = vec![1, 7, 23, 42, 75, 116, 794, 800, 1099];

    let found = binary_search(v, 23);
    assert_eq!(found, 2);
}

#[test]
fn test_binary_search_mid_point() {
    let v: Vec<isize> = vec![1, 7, 23, 42, 75, 116, 794, 800, 1099];

    let found = binary_search(v, 75);
    assert_eq!(found, 4);
}

#[test]
fn test_binary_search_out_of_bounds_upper() {
    let v: Vec<isize> = vec![1, 7, 23, 42, 75, 116, 794, 800, 1099];

    let found = binary_search(v, 1200);
    assert_eq!(found, -1);
}

#[test]
fn test_binary_search_out_of_bounds_lower() {
    let v: Vec<isize> = vec![1, 7, 23, 42, 75, 116, 794, 800, 1099];

    let found = binary_search(v, -23);
    assert_eq!(found, -1);
}
