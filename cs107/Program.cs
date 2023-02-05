var app = WebApplication.CreateBuilder(args).Build();

app.MapGet("/", () => "Hi");

app.Run();