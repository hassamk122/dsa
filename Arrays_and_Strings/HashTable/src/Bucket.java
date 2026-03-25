public class Bucket {

    private Node head;

    public void insert(String key){
        if( !search(key) ){
            Node newNode = new Node(key);
            newNode.next = head;
            head = newNode;
        }else {
            System.out.println("key already exists");
        }
    }

    public boolean search(String key){
        Node current = head;
        while(current != null){
            if(current.key.equals(key)){
                return true;
            }
            current = current.next;
        }
        return false;
    }

    public void delete(String key){
        if (head.key.equals(key)){
            head = head.next;
            return;
        }

        Node prevNode = head;
        while (prevNode != null){
            if (prevNode.next.key.equals(key)){
                prevNode.next = prevNode.next.next;
            }
            prevNode = prevNode.next;
        }
    }

    public void displayAll(){
        Node current = head;
        while(current != null){
            current.display();
            if (current.next != null){
                System.out.print(", ");
            }
            current = current.next;
        }
    }

}
