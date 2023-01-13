// System.Console.WriteLine(new int[] { 1, 2, 3, 7 }.twoSum(9));
// var bk = new Book { Title = "bk1" };
// System.Console.WriteLine($"{bk.Title}");

var r = new Random();

int correct = 0;
int wrong = 0;
var w = new List<(Int64, Int64)>();
while (true)
{

    var a = r.NextInt64(1, 10);
    var b = r.NextInt64(1, 10);
    var c = a * b;
    // System.Console.WriteLine($"{a} * {b} = {c}   ");

    System.Console.WriteLine($"{a} * {b} =");

    // var s = Console.ReadLine();
    var task = Task.Factory.StartNew(Console.ReadLine);
    var completedTask = await Task.WhenAny(task, Task.Delay(TimeSpan.FromSeconds(20)));
    var result = object.ReferenceEquals(task, completedTask) ? task.Result : "0";

    var ans = Int64.Parse(result!);
    if (ans == 0) break;

    if (ans != c)
    {
        wrong++;
        System.Console.WriteLine($"You Loose: {a} * {b} = {c}  correct={correct} wrong={wrong} percent={100.0 * correct / (correct + wrong)}");
        w.Add((a, b));
    }
    else
    {
        correct++;
        System.Console.WriteLine($"You Won! {a} * {b} = {c}  correct={correct} wrong={wrong} percent={100.0 * correct / (correct + wrong)}");
    }
}

System.Console.WriteLine("Practice List:");
foreach (var item in w)
{
    var a = item.Item1;
    var b = item.Item2;
    var c = a * b;
    System.Console.WriteLine($"{a} * {b} = {c}");
}