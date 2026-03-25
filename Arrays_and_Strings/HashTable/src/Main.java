void main() {
    HashTable hashTable = new HashTable();

    List<String> names = List.of(
            "hassam",
            "wajahat",
            "kiani",
            "rusty",
            "billi");

    for (String name : names) {
        hashTable.insert(name);
    }

    hashTable.display();

}
