using CalculatorBackend.Models.Entities;
using Microsoft.EntityFrameworkCore;

namespace CalculatorBackend;

public class ApplicationContext : DbContext
{
    public ApplicationContext(DbContextOptions<ApplicationContext> options) : base(options)
    {
        Database.EnsureCreated();
    }

    public DbSet<HistoryEntryEntity> HistoryEntries { get; set; } = null!;
}
