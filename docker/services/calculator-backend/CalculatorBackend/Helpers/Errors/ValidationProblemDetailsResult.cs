using CalculatorBackend.Helpers.Json;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.ModelBinding;

namespace CalculatorBackend.Helpers.Errors;

public class ValidationProblemDetailsResult : IActionResult
{
    public async Task ExecuteResultAsync(ActionContext context)
    {
        KeyValuePair<string, ModelStateEntry?>[] modelStateEntries = context.ModelState
            .Where(e => e.Value != null && e.Value.Errors.Count > 0)
            .ToArray();

        List<ValidationError> errors = new();

        if (modelStateEntries.Any())
        {
            foreach ((string key, ModelStateEntry? value) in modelStateEntries)
            {
                errors.AddRange(value?.Errors
                    .Select(modelStateError => new ValidationError(
                        name: key.ToSnakeCase(),
                        description: modelStateError.ErrorMessage)) ?? Array.Empty<ValidationError>());
            }
        }

        await new JsonErrorResponse<ValidationProblemDetails>(
            context: context.HttpContext,
            error: new ValidationProblemDetails(errors),
            statusCode: ValidationProblemDetails.ValidationStatusCode).WriteAsync();
    }
}
