public class HashTable {

    private int MAX_SIZE = 10;
    private Bucket[] buckets;

    public HashTable() {
        buckets = new Bucket[MAX_SIZE];
        for (int i = 0; i < MAX_SIZE; i++) {
            buckets[i] = new Bucket();
        }
    }

    public void insert(String key){
        int idx = hash(key);
        buckets[idx].insert(key);
    }

    public boolean search(String key){
        int idx = hash(key);
        return buckets[idx].search(key);
    }

    public void delete(String key){
        int idx = hash(key);
        buckets[idx].delete(key);
    }

    public void display(){
        for (int i = 0; i < MAX_SIZE; i++) {
            System.out.printf(" %d -> ",i);
            buckets[i].displayAll();
            System.out.println();
        }
    }



    private int hash(String key) {
        int sum = 0;
        for (int i = 0; i < key.length(); i++) {
            sum += (int)key.charAt(i);
        }
        return (sum % MAX_SIZE);
    }

}
