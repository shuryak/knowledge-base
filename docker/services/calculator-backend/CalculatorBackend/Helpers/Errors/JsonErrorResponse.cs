using CalculatorBackend.Helpers.Json;

namespace CalculatorBackend.Helpers.Errors;

public class JsonErrorResponse<T>
{
    private readonly HttpContext _context;
    private readonly T _error;
    private readonly int _statusCode;

    public JsonErrorResponse(HttpContext context, T error, int statusCode)
    {
        _context = context;
        _error = error;
        _statusCode = statusCode;
    }

    public Task WriteAsync()
    {
        _context.Response.ContentType = "application/json";
        _context.Response.StatusCode = _statusCode;

        return _context.Response.WriteAsync(_error.ToSnakeCase());
    }
}
