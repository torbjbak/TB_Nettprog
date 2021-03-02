import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class Main {
    public static void main(String[] args) {
        int nThreads = 0;
        try {
            nThreads = Integer.parseInt(args[0]);
        } catch (NumberFormatException nfe) {
            System.out.println("The argument must be an integer.");
            System.exit(1);
        }

        Workers workers = new Workers();
        ExecutorService executor = Executors.newFixedThreadPool(nThreads);
        final Lock mutex = new ReentrantLock();
        final Condition cond = mutex.newCondition();

        workers.post("TaskA", mutex, cond, 1500L);
        workers.post("TaskB", mutex,cond, 1200L);
        workers.postTimeout("TaskC", mutex, cond, 1000L, 5000L);
        workers.postWait("TaskD", mutex, cond, 2000L);

        for (Runnable wt : workers.getTasks()) {
            executor.execute(wt);
        }

        System.out.println("Finished all threads");
        executor.shutdown();
    }
}
