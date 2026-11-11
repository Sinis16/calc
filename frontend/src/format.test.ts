import { describe, it, expect } from "vitest";
import { formatResult } from "./format";

describe("formatResult", () => {
  it("removes floating-point noise", () => {
    expect(formatResult(0.1 + 0.2)).toBe("0.3");
  });

  it.each([
    [5, "5"],
    [-2.5, "-2.5"],
    [1e21, "1e+21"],
  ])("formats %d as %s", (value, expected) => {
    expect(formatResult(value)).toBe(expected);
  });
});
