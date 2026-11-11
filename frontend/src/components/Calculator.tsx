import { useState, type FormEvent } from "react";
import { calculate } from "../api";
import { OPERATIONS, type Operation } from "../operations";
import { parseOperand } from "../validation";
import { formatResult } from "../format";
import "./Calculator.css";

export function Calculator() {
  const [a, setA] = useState("");
  const [b, setB] = useState("");
  const [operation, setOperation] = useState<Operation>("add");
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const isUnary = OPERATIONS[operation].unary;

  async function handleSubmit(e: FormEvent) {
    e.preventDefault();
    setError(null);

    const parsedA = parseOperand(a, "A");
    if (!parsedA.ok) {
      setResult(null);
      setError(parsedA.error);
      return;
    }

    let bValue: number | undefined;
    if (!isUnary) {
      const parsedB = parseOperand(b, "B");
      if (!parsedB.ok) {
        setResult(null);
        setError(parsedB.error);
        return;
      }
      bValue = parsedB.value;
    }

    setLoading(true);
    try {
      setResult(await calculate({ operation, a: parsedA.value, b: bValue }));
    } catch (err) {
      setResult(null);
      setError(err instanceof Error ? err.message : "Unexpected error");
    } finally {
      setLoading(false);
    }
  }
  return (
    <main className="calculator">
      <h1>Calculator</h1>

      <form onSubmit={handleSubmit} noValidate>
        <div className="field">
          <label htmlFor="operand-a">Operand A</label>
          <input
            id="operand-a"
            inputMode="decimal"
            autoComplete="off"
            value={a}
            onChange={(e) => setA(e.target.value)}
          />
        </div>

        <div className="field">
          <label htmlFor="operation">Operation</label>
          <select
            id="operation"
            value={operation}
            onChange={(e) => setOperation(e.target.value as Operation)}
          >
            {(Object.keys(OPERATIONS) as Operation[]).map((op) => (
              <option key={op} value={op}>
                {OPERATIONS[op].symbol} {OPERATIONS[op].label}
              </option>
            ))}
          </select>
        </div>

        {!isUnary && (
          <div className="field">
            <label htmlFor="operand-b">Operand B</label>
            <input
              id="operand-b"
              inputMode="decimal"
              autoComplete="off"
              value={b}
              onChange={(e) => setB(e.target.value)}
            />
          </div>
        )}

        <button type="submit" disabled={loading}>
          {loading ? "Calculating…" : "Calculate"}
        </button>
      </form>

      <div className="output" role="status" aria-live="polite">
        {result !== null && (
          <p className="result">
            Result: <strong>{formatResult(result)}</strong>
          </p>
        )}
      </div>
      {error && (
        <p className="error" role="alert">
          {error}
        </p>
      )}
    </main>
  );
}
