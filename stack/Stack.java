import java.util.Arrays;
import java.util.EmptyStackException;

/**
 * This class implements a generic stack.
 */
public class Stack<E> {
	private E[] data;
	private int size; // number of elements in stack
	private static final int START_SIZE = 8; // size of underlying array

	/**
	 * Constructs an empty stack.
	 */
	public Stack() {
		// The data array will contain only E instances, all added in
		// the push method. This is sufficient to ensure type safety.
		@SuppressWarnings("unchecked") // for this declaration only
		E[] temp = (E[]) new Object[START_SIZE];
		data = temp;
	}
	
	/**
	 * Adds e to the top of this stack.
	 *
	 * @param e element to be added to this stack
	 */
	public void push(E e) {
		if (size == data.length) {
			data = Arrays.copyOf(data, 2*data.length);
		}
		data[size++] = e;
	}

	/**
	 * Removes and returns the top element of this stack.
	 *
	 * @return the element at the top of this stack
	 * @throws java.util.EmptyStackException if the stack is empty
	 */
	public E pop() {
		if (size == 0) {
			throw new EmptyStackException();
		}
		size--;
		E e = data[size];
		data[size] = null; // to avoid memory leak
		return e;
	}
	
	/**
	 * Returns the number of elements in this stack.
	 *
	 * @return the number of elements in this stack
	 */
	public int size() {
		return size;
	}

	/**
	 * Example demonstrating how this class can be used.
	 * Output: hello, world
	 */
	public static void main(String[] args) {
		Stack<String> s = new Stack<String>();
		s.push("world");
		s.push("hello, ");
		while (s.size() > 0) {
			System.out.print(s.pop());
		}
		System.out.println();
	}
}