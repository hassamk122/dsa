package main

// Give 2 strings, write a method to decide if one is permutation
// of other.

// 2 strings are permutation if they have same elements
// but in different order

// static boolean isPermutation(String a, String b) {
//     int Strlen = a.length();

//     if (Strlen != b.length()) {
//         return false;
//     }

//     char[] aToChars = convertToSortedCharArray(a);
//     char[] bToChars = convertToSortedCharArray(b);

//     return Arrays.equals(aToChars, bToChars);
// }

// private static char[] convertToSortedCharArray(String str) {
//     char[] chars = str.toCharArray();
//     Arrays.sort(chars);

//     return chars;
// }
