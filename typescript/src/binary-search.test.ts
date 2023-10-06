import { binarySearch } from "./binary-search";

describe("binarySearch", () => {
  test("should return index of element", () => {
    const numbers = [1, 7, 23, 42, 75, 116, 794, 800, 1099];
    const strings = ["a", "r", "t", "x"];

    expect(binarySearch<number>(numbers, 1)).toBe(0);
    expect(binarySearch<number>(numbers, 794)).toBe(6);
    expect(binarySearch<number>(numbers, 1099)).toBe(8);
    expect(binarySearch<number>(numbers, 2000)).toBe(-1);
    expect(binarySearch<number>(numbers, 0)).toBe(-1);
    expect(binarySearch<string>(strings, "r")).toBe(1);
    expect(binarySearch<string>(strings, "b")).toBe(-1);
    expect(
      binarySearch<number>([1, 7, 23, 42, 75, 116, 794, 800, 1099], 1200)
    ).toBe(-1);
  });
});
