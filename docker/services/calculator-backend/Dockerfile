FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /app
EXPOSE 80

COPY CalculatorBackend.sln ./
COPY CalculatorBackend/*.csproj ./CalculatorBackend/
RUN dotnet restore -r linux-x64

COPY CalculatorBackend/* ./CalculatorBackend/
RUN dotnet publish -c Release -o out -r linux-x64 --self-contained false --no-restore

FROM mcr.microsoft.com/dotnet/aspnet:6.0
WORKDIR /app
COPY --from=build /app/out .
ENTRYPOINT ["dotnet", "CalculatorBackend.dll"]
