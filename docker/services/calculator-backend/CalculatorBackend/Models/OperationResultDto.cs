using System.ComponentModel.DataAnnotations;

namespace CalculatorBackend.Models;

public record struct OperationResultDto([Required] double Result);
