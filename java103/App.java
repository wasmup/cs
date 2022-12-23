import java.util.Arrays;

public class App {
    public static void main(String[] args) {
        String[] a = { "0", "1", "12", "121", "2", "1000" };
        Arrays.sort(a);

        System.out.println(String.join(" ", a));
        // 0 1 1000 12 121 2
    }

}