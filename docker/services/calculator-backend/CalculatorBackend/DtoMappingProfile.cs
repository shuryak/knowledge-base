using AutoMapper;
using CalculatorBackend.Models;
using CalculatorBackend.Models.Entities;

namespace CalculatorBackend;

public class DtoMappingProfile : Profile
{
    public DtoMappingProfile()
    {
        CreateMap<BinaryOperationDto, SumOperation>();
        
        CreateMap<BinaryOperationDto, MultiplyOperation>();

        CreateMap<HistoryEntryEntity, HistoryEntry>();
    }
}
