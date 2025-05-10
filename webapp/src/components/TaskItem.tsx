import type { FC } from "react";
import type { Task } from "../lib/task/types";
import { useDeleteTask, useUpdateTask } from "../lib/task/query";

export const TaskItem: FC<{ task: Task }> = ({ task }) => {
  const { mutate: updateTask } = useUpdateTask(task.id);
  const { mutate: deleteTask } = useDeleteTask(task.id);

  const handleUpdateTask = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();

    const isChecked = e.target.checked;

    console.log("task", task);
    console.log("isChecked", isChecked);

    updateTask(
      { status: e.target.checked ? "done" : "todo" },
      {
        onError: (error) => {
          // Give feedback to the user
          console.error(error);
        },
        onSuccess: () => {
          // TODO: Give feedback to the user
          console.log("Task updated");
        },
      },
    );
  };

  const handleDelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    deleteTask(undefined, {
      onError: (error) => {
        // Give feedback to the user
        console.error(error);
      },
      onSuccess: () => {
        // TODO: Give feedback to the user
        console.log("Task deleted");
      },
    });
  };

  return (
    <li>
      <input
        type="checkbox"
        checked={task.status === "done"}
        onChange={handleUpdateTask}
      />
      <h3>{task.title}</h3>
      <button onClick={handleDelete}>Delete</button>
    </li>
  );
};
