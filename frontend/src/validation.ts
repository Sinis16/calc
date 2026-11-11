export type ParseResult =
  | { ok: true; value: number }
  | { ok: false; error: string };

export function parseOperand(raw: string, label: string): ParseResult {
  const trimmed = raw.trim();
  if (trimmed === "") {
    return { ok: false, error: `${label} is required` };
  }
  const value = Number(trimmed);
  if (!Number.isFinite(value)) {
    return { ok: false, error: `${label} must be a valid number` };
  }
  return { ok: true, value };
}
