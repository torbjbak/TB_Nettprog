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

    public void post(String name, Lock mutex, long duration) {
        mutex.lock();
        tasks.add(new WorkerThread(name, mutex, duration));
        mutex.unlock();
    }

    public void postTimeout(String name, Lock mutex, long duration, long delay) {
        mutex.lock();
        tasks.add(new WorkerThread(name, mutex, duration, delay));
        mutex.unlock();
    }
}
