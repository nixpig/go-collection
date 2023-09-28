fn main() {}

fn insertion_sort(arr: &mut Vec<usize>) -> &Vec<usize> {
    'outer: for i in 1..(arr.len()) {
        let current_item = arr[i];

        'inner: for j in (0..=(i - 1)).rev() {
            println!("inner j {:?}", j);
            if current_item >= arr[j] {
                arr[j + 1] = current_item;
                continue 'outer;
            } else {
                arr[j + 1] = arr[j];
                arr[j] = current_item;
                continue 'inner;
            }
        }
    }

    arr
}

#[test]
fn test_insertion_sort() {
    let mut unsorted: Vec<usize> = vec![23, 13, 7, 69, 42];
    println!("{:?}", unsorted);
    let sorted = insertion_sort(&mut unsorted);

    let expect: Vec<usize> = vec![7, 13, 23, 42, 69];

    assert_eq!(*sorted, expect);
}
