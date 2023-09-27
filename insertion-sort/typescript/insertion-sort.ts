export const insertionSort = <T>(arr: T[]): T[] => {
  const N = arr.length;

  let current: T;

  let i = 1;
  let j: number;

  while (i < N) {
    current = arr[i];

    j = i - 1;

    while (j >= 0 && arr[j] > current) {
      arr[j + 1] = arr[j];
      j = j - 1;
    }

    arr[j + 1] = current;
    i = i + 1;
  }

  return arr;
};
