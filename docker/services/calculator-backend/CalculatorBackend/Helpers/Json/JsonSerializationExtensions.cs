using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;

namespace CalculatorBackend.Helpers.Json;

public static class JsonSerializationExtensions
{
    private static readonly SnakeCaseNamingStrategy _snakeCaseNamingStrategy = new();

    private static readonly JsonSerializerSettings _snakeCaseSettings = new()
    {
        ContractResolver = new DefaultContractResolver
        {
            NamingStrategy = _snakeCaseNamingStrategy
        }
    };

    public static string ToSnakeCase<T>(this T instance)
    {
        if (instance == null)
        {
            throw new ArgumentNullException(paramName: nameof(instance));
        }

        return JsonConvert.SerializeObject(instance, _snakeCaseSettings);
    }

    public static string ToSnakeCase(this string @string)
    {
        if (@string == null)
        {
            throw new ArgumentNullException(paramName: nameof(@string));
        }

        return _snakeCaseNamingStrategy.GetPropertyName(@string, false);
    }
}
