import org.hibernate.annotations.GenericGenerator;

import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import java.util.concurrent.locks.Condition;
import java.util.function.Function;

public class Task {
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
    private Function<Condition, Void> func;

    public Task(Function<Condition, Void> func) {
        this.func = func;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Function<Condition, Void> getFunc() {
        return func;
    }

    public void setFunc(Function<Condition, Void> func) {
        this.func = func;
    }
}
