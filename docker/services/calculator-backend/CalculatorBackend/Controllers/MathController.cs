using AutoMapper;
using CalculatorBackend.Abstractions;
using CalculatorBackend.Models;
using Microsoft.AspNetCore.Mvc;

namespace CalculatorBackend.Controllers;

[ApiController]
[Route("api/[controller].[action]")]
public class MathController : ControllerBase
{
    private readonly IRepository<HistoryEntry> _historyEntryRepository;
    private readonly IMapper _mapper;
    
    public MathController(IMapper mapper, IRepository<HistoryEntry> historyEntryRepository)
    {
        _mapper = mapper;
        _historyEntryRepository = historyEntryRepository;
    }
    
    [HttpPost]
    public ActionResult<OperationResultDto> Calculate(BinaryOperationDto dto)
    {
        IBinaryOperation operation;

        switch (dto.OperationSymbol)
        {
            case "+":
                operation = _mapper.Map<SumOperation>(dto);
                break;
            case "*":
                operation = _mapper.Map<MultiplyOperation>(dto);
                break;
            default:
                return BadRequest("Invalid operation symbol");
        }

        double operationResult = operation.GetResult();
        
        _historyEntryRepository.Create(new HistoryEntry
        {
            CreatedAt = DateTime.UtcNow,
            FirstOperand = dto.FirstOperand,
            SecondOperand = dto.SecondOperand,
            OperationSymbol = dto.OperationSymbol,
            Result = operationResult
        });

        return new OperationResultDto(operationResult);
    }
}
