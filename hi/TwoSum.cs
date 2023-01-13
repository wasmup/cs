public static class TwoSum
{
    public static (int, int) twoSum(this int[] a, int target)
    {
        var m = new Dictionary<int, int>();
        for (int i = 0; i < a.Length; i++)
        {
            var v = a[i];
            if (m.ContainsKey(v)) return (m[v], i);
            m[target - v] = i;
        }
        return (-1, -1);
    }

}