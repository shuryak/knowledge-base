using CalculatorBackend;
using CalculatorBackend.Abstractions;
using CalculatorBackend.Helpers.Errors;
using CalculatorBackend.Helpers.Json;
using CalculatorBackend.Models;
using CalculatorBackend.Repositories;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

WebApplicationBuilder builder = WebApplication.CreateBuilder(args);
ConfigurationManager configuration = builder.Configuration;

builder.Services.AddTransient<IRepository<HistoryEntry>, HistoryEntryRepository>();

builder.Services.AddAutoMapper(typeof(DtoMappingProfile));

builder.Services.AddCors(options =>
{
    options.AddPolicy("AllowOrigin", corsPolicyBuilder => 
        corsPolicyBuilder.AllowAnyOrigin().AllowAnyMethod().AllowAnyHeader());
});

builder.Services.AddControllers()
    .AddJsonOptions(x =>
    {
        x.JsonSerializerOptions.PropertyNamingPolicy = new SnakeCaseNamingPolicy();
    });
builder.Services
    .Configure<ApiBehaviorOptions>(x =>
    {
        x.InvalidModelStateResponseFactory = ctx => new ValidationProblemDetailsResult();
    });

string connectionString = $"Server={Environment.GetEnvironmentVariable("DB_HOST")};" +
                          $"Port={Environment.GetEnvironmentVariable("DB_PORT")};" +
                          $"Database={Environment.GetEnvironmentVariable("DB_NAME")};" +
                          $"User Id={Environment.GetEnvironmentVariable("DB_USER")};" +
                          $"Password={Environment.GetEnvironmentVariable("DB_PASSWORD")}";

builder.Services.AddDbContext<ApplicationContext>(x =>
    x.UseNpgsql(connectionString)
);

WebApplication app = builder.Build();
// Configure the HTTP request pipeline.
app.UseCors("AllowOrigin");
app.UseRouting();
app.MapControllers();
app.Run();
