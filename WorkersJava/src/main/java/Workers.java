import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;

public class Workers {
    private final List<Runnable> tasks;

    public Workers() {
        this.tasks = new ArrayList<>();
    }

    public List<Runnable> getTasks() {
        return tasks;
    }

    public void post(String name, Lock mutex, Condition cond, long duration) {
        mutex.lock();
        tasks.add(new WorkerThread(name, mutex, cond, duration, false));
        mutex.unlock();
    }

    public void postTimeout(String name, Lock mutex, Condition cond, long duration, long delay) {
        mutex.lock();
        tasks.add(new WorkerThread(name, mutex, cond, duration, delay, false));
        mutex.unlock();
    }

    public void postWait(String name, Lock mutex, Condition cond, long duration) {
        mutex.lock();
        tasks.add(new WorkerThread(name, mutex, cond, duration, true));
        mutex.unlock();
    }
}
