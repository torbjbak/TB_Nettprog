import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;
import java.util.function.Function;

public class Workers {
    private final List<Worker> workers;
    private final List<Task> tasks;
    private final Lock lock = new ReentrantLock();
    private final Condition cond = lock.newCondition();

    public Workers() {
        this.workers = new ArrayList<>();
        this.tasks = new ArrayList<>();
    }

    public List<Worker> getWorkers() {
        return workers;
    }

    public List<Task> getTasks() {
        return tasks;
    }

    public void addWorker() {
        workers.add(new Worker());
    }

    public void addTask(Function<Condition, Void> f) {
        tasks.add(new Task(f));
    }
}
