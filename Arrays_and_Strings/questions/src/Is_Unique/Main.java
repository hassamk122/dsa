
// Implement an algorithm to determine if a string has all unique characters.
// what if you can't use any other data structures
void main(){
    String word  = "aba";
    String word2 = "ab";
    String word3 = "";
    String word4 = "a";
    String word5 = "has";
    String word6 = "hassam";


    System.out.println(word+(isUnique(word)?" is unique":"is not unique"));
    System.out.println(word2+(isUnique(word2)?" is unique":"is not unique"));
    System.out.println(word3+(isUnique(word3)?" is unique":"is not unique"));
    System.out.println(word4+(isUnique(word4)?" is unique":"is not unique"));
    System.out.println(word5+(isUnique(word5)?" is unique":"is not unique"));
    System.out.println(word6+(isUnique(word6)?" is unique":"is not unique"));
}

static boolean isUnique(String s){
    int strLen = s.length();

    if (strLen == 0) {
        return true;
    }

    boolean[] checkChar = new boolean[255];
    char[] charArray = s.toCharArray();

    for (int i = 0; i < strLen; i++){
        if (checkChar[charArray[i]]){
            return false;
        }
        checkChar[charArray[i]] = true;
    }
    return true;
}