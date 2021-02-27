import org.hibernate.annotations.GenericGenerator;

import javax.persistence.*;

public class Worker {
    @Id
    @GeneratedValue(
            strategy = GenerationType.AUTO,
            generator = "native"
    )
    @GenericGenerator(
            name = "native",
            strategy = "native"
    )
    private Long id;
    @OneToOne(
            cascade = CascadeType.ALL
    )
    private Task task;
    private Boolean busy;

    public Worker() {
        this.task = null;
        this.busy = false;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Task getTask() {
        return task;
    }

    public void setTask(Task task) {
        this.task = task;
    }

    public Boolean getBusy() {
        return busy;
    }

    public void setBusy(Boolean busy) {
        this.busy = busy;
    }

    public String toString() {
        return "ID: "+ id +"\nTask ID: "+ task.getId() +"\nBusy: "+ busy;
    }
}
