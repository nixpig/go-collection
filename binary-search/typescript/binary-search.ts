export const binarySearch = <T>(arr: T[], target: T): number => {
  let indexLow = 0;
  let indexHigh = arr.length - 1;
  let indexMid: number;

  while (indexLow <= indexHigh) {
    indexMid = Math.floor((indexHigh + indexLow) / 2);

    if (arr[indexMid] === target) {
      return indexMid;
    }

    if (arr[indexMid] > target) {
      indexHigh = indexMid - 1;
    } else {
      indexLow = indexMid + 1;
    }
  }
  return -1;
};
