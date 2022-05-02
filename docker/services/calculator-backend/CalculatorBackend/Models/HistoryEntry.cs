namespace CalculatorBackend.Models;

public record struct HistoryEntry(
    int Id,
    DateTime CreatedAt,
    double FirstOperand,
    double SecondOperand,
    string OperationSymbol,
    double Result
    );
