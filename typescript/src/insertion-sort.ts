export const insertionSort = <T>(arr: T[]): T[] => {
  outer: for (let i = 1; i < arr.length; ++i) {
    const currentItem = arr[i];

    inner: for (let j = i - 1; j >= 0; --j) {
      if (currentItem >= arr[j]) {
        arr[j + 1] = currentItem;
        continue outer;
      } else {
        arr[j + 1] = arr[j];
        arr[j] = currentItem;
        continue inner;
      }
    }
  }

  return arr;
};
