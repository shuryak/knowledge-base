using System.Net;
using Microsoft.AspNetCore.Mvc;

namespace CalculatorBackend.Helpers.Errors;

public class ValidationProblemDetails : ProblemDetails
{
    public const int ValidationStatusCode = (int)HttpStatusCode.BadRequest;

    public ValidationProblemDetails(ICollection<ValidationError> validationErrors)
    {
        ValidationErrors = validationErrors;

        Status = ValidationStatusCode;
        Title = "Request Validation Error";
    }

    public ICollection<ValidationError> ValidationErrors { get; }

    public string RequestId => Guid.NewGuid().ToString();
}
