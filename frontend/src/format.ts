// Rounds away floating-point noise for display, up to 12 significant digits
export function formatResult(value: number): string {
  return String(Number.parseFloat(value.toPrecision(12)));
}
