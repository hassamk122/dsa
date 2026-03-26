public class ArrayList <T>{
    private T[] array;
    private int capacity;
    private int size;

    public ArrayList(){
        capacity = 2;
        array = (T[]) new Object[capacity];
        size = 0;
    }


    public void add(T element){
        if (isFull()){
            T[] newArray = (T[]) new Object[capacity*2];
            System.arraycopy(array, 0, newArray, 0, capacity);
            capacity *= 2;
            array = newArray;
        }
            array[size++] = element;
    }

    public void remove(int index){
        if (isEmpty()){
            return;
        }

        if (index >= capacity ){
            return;
        }

        for (int i = index; i < capacity - 1; i++){
                array[i] = array[i+1];
        }
        size--;
    }

    public boolean isEmpty(){
        return size == 0;
    }

    public int getCapacity(){
        return capacity;
    }

    public int getSize(){
        return size;
    }

    private boolean isFull(){
        return size == capacity;
    }


    @Override
    public String toString() {
        StringBuilder strBuilder = new StringBuilder();
        for (int i = 0; i < size; i++){
            strBuilder.append(array[i]);
        }

        return strBuilder.toString();
    }

}
