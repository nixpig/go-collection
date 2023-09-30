export const binarySearch = <T>(arr: T[], target: T): number => {
  let indexLow = 0;
  let indexHigh = arr.length - 1;

  while (indexLow <= indexHigh) {
    const indexMid = Math.floor((indexHigh + indexLow) / 2);
    const currentValue = arr[indexMid];

    if (currentValue === target) {
      return indexMid;
    }

    if (currentValue > target) {
      indexHigh = indexMid - 1;
    } else {
      indexLow = indexMid + 1;
    }
  }

  return -1;
};
