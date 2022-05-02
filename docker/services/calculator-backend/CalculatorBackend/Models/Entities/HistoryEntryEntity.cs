using System.ComponentModel.DataAnnotations;

namespace CalculatorBackend.Models.Entities;

public class HistoryEntryEntity
{
    [Key]
    public int Id { get; set; }
    [Required]
    public DateTime CreatedAt { get; set; }
    [Required]
    public double FirstOperand { get; set; }
    [Required]
    public double SecondOperand { get; set; }
    [Required]
    public string OperationSymbol { get; set; } = null!;
    [Required]
    public double Result { get; set; }
}
