import type { Operation } from "./operations";

export interface CalculateRequest {
  operation: Operation;
  a: number;
  b?: number;
}

interface CalculateResponse {
  result?: number;
  error?: string;
}

export async function calculate(req: CalculateRequest): Promise<number> {
  let res: Response;
  try {
    res = await fetch("/api/v1/calculate", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(req),
    });
  } catch {
    throw new Error("Could not reach the server");
  }

  const data = (await res.json().catch(() => null)) as CalculateResponse | null;

  if (!res.ok) {
    throw new Error(data?.error ?? `Request failed with status ${res.status}`);
  }
  if (typeof data?.result !== "number") {
    throw new Error("Unexpected response from server");
  }
  return data.result;
}
