namespace CalculatorBackend.Abstractions;

public interface IBinaryOperation
{
    public double FirstOperand { get; init; }
    public double SecondOperand { get; init; }

    public double GetResult();
}
