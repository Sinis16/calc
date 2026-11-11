import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import { calculate } from "./api";

const fetchMock = vi.fn();

function mockResponse(body: unknown, status = 200) {
  return {
    ok: status >= 200 && status < 300,
    status,
    json: async () => body,
  } as Response;
}

beforeEach(() => {
  fetchMock.mockReset();
  vi.stubGlobal("fetch", fetchMock);
});

afterEach(() => {
  vi.unstubAllGlobals();
});

describe("calculate", () => {
  it("posts the request and returns the result", async () => {
    fetchMock.mockResolvedValue(mockResponse({ result: 5 }));

    await expect(calculate({ operation: "add", a: 2, b: 3 })).resolves.toBe(5);

    expect(fetchMock).toHaveBeenCalledWith("/api/v1/calculate", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ operation: "add", a: 2, b: 3 }),
    });
  });

  it("surfaces the backend error message", async () => {
    fetchMock.mockResolvedValue(
      mockResponse({ error: "division by zero" }, 400),
    );

    await expect(
      calculate({ operation: "divide", a: 1, b: 0 }),
    ).rejects.toThrow("division by zero");
  });

  it("reports network failures", async () => {
    fetchMock.mockRejectedValue(new TypeError("Failed to fetch"));

    await expect(calculate({ operation: "add", a: 1, b: 1 })).rejects.toThrow(
      "Could not reach the server",
    );
  });

  it("falls back to the status code when the body is not JSON", async () => {
    fetchMock.mockResolvedValue({
      ok: false,
      status: 502,
      json: async () => {
        throw new SyntaxError("Unexpected token");
      },
    } as unknown as Response);

    await expect(calculate({ operation: "add", a: 1, b: 1 })).rejects.toThrow(
      "Request failed with status 502",
    );
  });

  it("rejects a success response without a result", async () => {
    fetchMock.mockResolvedValue(mockResponse({}));

    await expect(calculate({ operation: "add", a: 1, b: 1 })).rejects.toThrow(
      "Unexpected response from server",
    );
  });
});
