import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { Calculator } from "./Calculator";
import { calculate } from "../api";

vi.mock("../api", () => ({ calculate: vi.fn() }));
const mockCalculate = vi.mocked(calculate);

beforeEach(() => {
  mockCalculate.mockReset();
});

describe("Calculator", () => {
  it("sends the operands and shows the result", async () => {
    mockCalculate.mockResolvedValue(5);
    const user = userEvent.setup();
    render(<Calculator />);

    await user.type(screen.getByLabelText("Operand A"), "2");
    await user.type(screen.getByLabelText("Operand B"), "3");
    await user.click(screen.getByRole("button", { name: "Calculate" }));

    expect(await screen.findByText("5")).toBeInTheDocument();
    expect(mockCalculate).toHaveBeenCalledWith({
      operation: "add",
      a: 2,
      b: 3,
    });
  });

  it("validates input before calling the API", async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await user.click(screen.getByRole("button", { name: "Calculate" }));

    expect(screen.getByRole("alert")).toHaveTextContent("A is required");
    expect(mockCalculate).not.toHaveBeenCalled();
  });

  it("rejects non-numeric input", async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await user.type(screen.getByLabelText("Operand A"), "abc");
    await user.click(screen.getByRole("button", { name: "Calculate" }));

    expect(screen.getByRole("alert")).toHaveTextContent(
      "A must be a valid number",
    );
    expect(mockCalculate).not.toHaveBeenCalled();
  });

  it("hides operand B for square root", async () => {
    mockCalculate.mockResolvedValue(4);
    const user = userEvent.setup();
    render(<Calculator />);

    await user.selectOptions(screen.getByLabelText("Operation"), "sqrt");
    expect(screen.queryByLabelText("Operand B")).not.toBeInTheDocument();

    await user.type(screen.getByLabelText("Operand A"), "16");
    await user.click(screen.getByRole("button", { name: "Calculate" }));

    expect(await screen.findByText("4")).toBeInTheDocument();
    expect(mockCalculate).toHaveBeenCalledWith({
      operation: "sqrt",
      a: 16,
      b: undefined,
    });
  });

  it("shows errors returned by the API", async () => {
    mockCalculate.mockRejectedValue(new Error("division by zero"));
    const user = userEvent.setup();
    render(<Calculator />);

    await user.selectOptions(screen.getByLabelText("Operation"), "divide");
    await user.type(screen.getByLabelText("Operand A"), "1");
    await user.type(screen.getByLabelText("Operand B"), "0");
    await user.click(screen.getByRole("button", { name: "Calculate" }));

    expect(await screen.findByRole("alert")).toHaveTextContent(
      "division by zero",
    );
  });

  it("validates operand B for binary operations", async () => {
    const user = userEvent.setup();
    render(<Calculator />);

    await user.type(screen.getByLabelText("Operand A"), "5");
    await user.click(screen.getByRole("button", { name: "Calculate" }));

    expect(screen.getByRole("alert")).toHaveTextContent("B is required");
    expect(mockCalculate).not.toHaveBeenCalled();
  });
});
