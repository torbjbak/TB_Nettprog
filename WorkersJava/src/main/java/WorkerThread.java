import java.util.concurrent.locks.Lock;

public class WorkerThread implements Runnable{
    private final String task;
    private final Lock mutex;
    private final long duration;
    private final long delay;

    public WorkerThread(String task, Lock mutex, long duration, long delay) {
        this.task = task;
        this.mutex = mutex;
        this.duration = duration;
        this.delay = delay;
    }

    public WorkerThread(String task, Lock mutex, long duration) {
        this.task = task;
        this.mutex = mutex;
        this.duration = duration;
        this.delay = 0;
    }

    @Override
    public void run() {
        try {
            //this.mutex.lock();
            if (this.delay > 0) {
                Thread.sleep(this.delay);
            }
            System.out.println(Thread.currentThread().getName() +" (Start) "+ task +
                    " ("+ (double)duration/1000 +"s)");
            //this.mutex.unlock();
            Thread.sleep(this.duration);
            System.out.println(Thread.currentThread().getName() +" (End)"+ task);
        } catch (InterruptedException ie) {
            ie.printStackTrace();
        }
    }
}
