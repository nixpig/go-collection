import { insertionSort } from "./insertion-sort";
const unsorted = [5, 2, 4, 6, 1, 3];

describe("insertion sort", () => {
  test("should sort an array of numbers", () => {
    expect(insertionSort<number>(unsorted)).toEqual([1, 2, 3, 4, 5, 6]);
  });
});
