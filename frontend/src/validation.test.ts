import { describe, it, expect } from "vitest";
import { parseOperand } from "./validation";

describe("parseOperand", () => {
  it.each([
    ["42", 42],
    ["-3", -3],
    ["3.5", 3.5],
    ["  7  ", 7],
    ["1e3", 1000],
  ])("parses %j as %d", (raw, expected) => {
    expect(parseOperand(raw, "A")).toEqual({ ok: true, value: expected });
  });

  it.each(["", "   "])("rejects empty input %j", (raw) => {
    expect(parseOperand(raw, "A")).toEqual({
      ok: false,
      error: "A is required",
    });
  });

  it.each(["abc", "1,5", "Infinity", "--2"])(
    "rejects invalid input %j",
    (raw) => {
      expect(parseOperand(raw, "B")).toEqual({
        ok: false,
        error: "B must be a valid number",
      });
    },
  );
});
