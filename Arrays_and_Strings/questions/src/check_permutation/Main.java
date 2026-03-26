package check_permutation;
// Give 2 strings, write a method to decide if one is permutation
// of other.

// 2 strings are permutation if they have same elements
// but in different order

import java.util.Arrays;

public class Main {

    void main() {
        String a = "hello";
        String b = "hello ";


        String a1 = "abba";
        String b1 = "bbaa";


        String a2 = "";
        String b2 = "";


        System.out.println(isPermutation(a, b));
        System.out.println(isPermutation(a1, b1));
        System.out.println(isPermutation(a2, b2));
    }

    static boolean isPermutation(String a, String b) {
        int Strlen = a.length();

        if (Strlen != b.length()) {
            return false;
        }

        char[] aToChars = convertToSortedCharArray(a);
        char[] bToChars = convertToSortedCharArray(b);


        return Arrays.equals(aToChars, bToChars);
    }

    private static char[] convertToSortedCharArray(String str) {
        char[] chars = str.toCharArray();
        Arrays.sort(chars);

        return chars;
    }
}