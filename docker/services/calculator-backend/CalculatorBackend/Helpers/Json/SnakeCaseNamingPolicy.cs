using System.Text.Json;

namespace CalculatorBackend.Helpers.Json;

public class SnakeCaseNamingPolicy : JsonNamingPolicy
{
    public override string ConvertName(string name) => name.ToSnakeCase();
}
