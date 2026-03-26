void main(){
    ArrayList<String> sentence = new ArrayList<>();

    sentence.add("My");
    sentence.add("name");
    sentence.add("is");
    sentence.add("My");
    sentence.add("name");
    sentence.add("is");

    sentence.remove(1);

    System.out.println(sentence.isEmpty());
    System.out.println(sentence.getCapacity());
    System.out.println(sentence.getSize());
    System.out.println(sentence);
}
