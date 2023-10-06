import { insertionSort } from "./insertion-sort";
const unsorted = [23, 13, 7, 69, 42];

describe("insertion sort", () => {
  test("should sort an array of numbers", () => {
    expect(insertionSort<number>(unsorted)).toEqual([7, 13, 23, 42, 69]);
  });
});
