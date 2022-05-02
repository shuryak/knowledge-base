using System.ComponentModel.DataAnnotations;

namespace CalculatorBackend.Models;

public record struct BinaryOperationDto(
    [Required]
    double FirstOperand,
    [Required]
    double SecondOperand,
    [Required]
    string OperationSymbol
);
