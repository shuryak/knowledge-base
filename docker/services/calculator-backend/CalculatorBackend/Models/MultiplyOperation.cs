using CalculatorBackend.Abstractions;

namespace CalculatorBackend.Models;

public readonly record struct MultiplyOperation : IBinaryOperation
{
    public double FirstOperand { get; init; }
    public double SecondOperand { get; init; }
    
    public double GetResult()
    {
        return FirstOperand * SecondOperand;
    }
}
