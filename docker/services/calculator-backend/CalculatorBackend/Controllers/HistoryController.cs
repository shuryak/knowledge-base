using CalculatorBackend.Abstractions;
using CalculatorBackend.Models;
using Microsoft.AspNetCore.Mvc;

namespace CalculatorBackend.Controllers;

[ApiController]
[Route("api/[controller].[action]")]
public class HistoryController : ControllerBase
{
    private readonly IRepository<HistoryEntry> _historyEntryRepository;

    public HistoryController(IRepository<HistoryEntry> historyEntryRepository)
    {
        _historyEntryRepository = historyEntryRepository;
    }

    [HttpGet]
    public ActionResult<List<HistoryEntry>> Get()
    {
        return _historyEntryRepository.GetList();
    }

    [HttpDelete]
    public void Clear()
    {
        _historyEntryRepository.RemoveAll();
    }
}
