namespace CalculatorBackend.Abstractions;

public interface IRepository<TModel> where TModel : struct
{
    public void Create(TModel entry);

    public TModel? GetById(int id);

    public List<TModel> GetList();

    public void RemoveAll();
}
