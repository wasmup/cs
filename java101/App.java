public class App {
    public static void main(String[] args) {
        System.out.println(spinWords(",olleH !dlrow").equals("Hello, world!"));
    }

    static String spinWords(String sentence) {
        String[] a = sentence.split(" ");
        for (int i = a.length - 1; i >= 0; i--) {
            if (a[i].length() > 4) {
                a[i] = new StringBuilder(a[i]).reverse().toString();
            }
        }
        return String.join(" ", a);
    }
}

// Write a function that takes in a string of one or more words, and returns the
// same string, but with all five or more letter words reversed:
