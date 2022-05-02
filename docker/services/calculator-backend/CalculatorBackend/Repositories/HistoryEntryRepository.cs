using AutoMapper;
using CalculatorBackend.Abstractions;
using CalculatorBackend.Models;
using CalculatorBackend.Models.Entities;
using Microsoft.EntityFrameworkCore;

namespace CalculatorBackend.Repositories;

public class HistoryEntryRepository : IRepository<HistoryEntry>
{
    private readonly ApplicationContext _dbContext;
    private readonly IMapper _mapper;
    
    public HistoryEntryRepository(ApplicationContext dbContext, IMapper mapper)
    {
        _dbContext = dbContext;
        _mapper = mapper;
    }

    public void Create(HistoryEntry entry)
    {
        HistoryEntryEntity entity = new HistoryEntryEntity
        {
            CreatedAt = entry.CreatedAt,
            FirstOperand = entry.FirstOperand,
            SecondOperand = entry.SecondOperand,
            OperationSymbol = entry.OperationSymbol,
            Result = entry.Result
        };

        _dbContext.HistoryEntries.Add(entity);

        _dbContext.SaveChanges();
    }

    public HistoryEntry? GetById(int id)
    {
        HistoryEntryEntity? entity = _dbContext.HistoryEntries.AsNoTracking().FirstOrDefault(x => x.Id == id);

        return entity == null ? null : _mapper.Map<HistoryEntry>(entity);
    }

    public List<HistoryEntry> GetList()
    {
        return _dbContext.HistoryEntries.Select(entity => _mapper.Map<HistoryEntry>(entity)).ToList();
    }

    public void RemoveAll()
    {
        foreach (HistoryEntryEntity entity in _dbContext.HistoryEntries)
        {
            _dbContext.HistoryEntries.Remove(entity);
        }

        _dbContext.SaveChanges();
    }
}
