import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;

public class WorkerThread implements Runnable{
    private final String task;
    private final Lock mutex;
    private final Condition cond;
    private final long duration;
    private final long delay;
    private final boolean wait;

    public WorkerThread(String task, Lock mutex, Condition cond, long duration, long delay, boolean wait) {
        this.task = task;
        this.mutex = mutex;
        this.cond = cond;
        this.duration = duration;
        this.delay = delay;
        this.wait = wait;
    }

    public WorkerThread(String task, Lock mutex, Condition cond, long duration, boolean wait) {
        this.task = task;
        this.mutex = mutex;
        this.cond = cond;
        this.duration = duration;
        this.delay = 0;
        this.wait = wait;
    }

    @Override
    public void run() {
        try {
            if (this.delay > 0) {
                Thread.sleep(this.delay);
            }

            if (wait) {
                mutex.lock();
                cond.await();
                mutex.unlock();
            }

            System.out.println(Thread.currentThread().getName() +" (Start) "+ task +
                    " ("+ (double)duration/1000 +"s)");

            Thread.sleep(this.duration);
            System.out.println(Thread.currentThread().getName() +" (End)"+ task);

            if (this.delay > 0 && !wait) {
                mutex.lock();
                cond.signal();
                mutex.unlock();
            }
        } catch (InterruptedException ie) {
            ie.printStackTrace();
        }
    }
}
